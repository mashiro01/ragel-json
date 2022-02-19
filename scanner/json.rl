package scanner

import (
    "main/token"
)

%%{
    machine json_lexer;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;

    prepush {
        /* stack growing code */
    }

    postpop {
        /* stack shrinking code */
    }
}%%

func (lex *JsonLexer) initJsonLexer() {
    %% write init;
}

func (lex *JsonLexer) Lex() *token.Token {
    eof := lex.pe
    var tok token.ID
    tkn := lex.tokenPool.GenBlock()

    %%{
        newline  = ( '\r\n' >(nl, 1) | '\r' >(nl, 0) | '\n' >(nl, 0) );
        any_line = any | newline;

        # comment
        j_comment = "/*" any_line* :>> "*/";

        # whitespace
        j_whitespace = [\t\v\f ];
        j_whitespace_newline = j_whitespace | newline;

        # number
        positive = '+' | '-';
        integer  = digit+;
        fraction = (integer? '.' integer) | (integer '.' integer?);
        exponent = (integer | fraction) [eE] positive? integer;
        j_number =
            (
                positive?
                (
                    (
                        integer  |
                        fraction |
                        exponent
                    )
                )
            );

        # string
        unicode_format = 'u' xdigit {4};
        escaped_charset = '\\' ( '"' | '\\' | '/' | 'b' | 'f' | 'n' | 't' | 'r' | unicode_format );
        unescaped_charset = any - '"' - '\\' - cntrl;
        j_string = '"' (escaped_charset | unescaped_charset)* '"';

        # bool
        j_bool = "true" | "false";

        # null
        j_null = "null";

        # array
        j_array := |*
            '['                         => {
                lex.bracketStack += 1
                lex.pushSubTokenStack(token.J_ARRAY, lex.ts)
            };
            ']'                         => {
                lex.bracketStack -= 1
                lex.popSubTokenStack(tkn, lex.te)

                if lex.bracketStack == 0 {
                    fret;
                }
            };
            ','                         => { };
            j_whitespace                => { };
            j_number                    => { lex.addSubToken(tkn, token.J_NUMBER, lex.ts, lex.te) };
            j_string                    => { lex.addSubToken(tkn, token.J_STRING, lex.ts, lex.te) };
            j_bool                      => { lex.addSubToken(tkn, token.J_BOOL, lex.ts, lex.te) };
            j_null                      => { lex.addSubToken(tkn, token.J_NULL, lex.ts, lex.te) };
        *|;

        # object
        j_object := |*
        *|;

        # value
        main := |*
            j_number => { tok = token.J_NUMBER; fbreak; };
            j_string => { tok = token.J_STRING; fbreak; };
            j_bool   => { tok = token.J_BOOL; fbreak; };
            j_null   => { tok = token.J_NULL; fbreak; };
            '['      => {
                tok = token.J_ARRAY;
                lex.bracketStack += 1;
                fcall j_array;
                fbreak;
            };
        *|;

        write exec;
    }%%

    tkn.Value = lex.data[lex.ts:lex.te]
    tkn.ID = token.ID(tok)

    return tkn
}
