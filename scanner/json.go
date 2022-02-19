
//.... json.rl:1
package scanner

import (
    "main/token"
)


//.... json.go:11
var _json_lexer_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3,
	1, 4, 1, 5, 1, 6, 1, 7,
	1, 8, 1, 9, 1, 10, 1, 11,
	1, 12, 1, 13, 1, 14, 1, 15,
	1, 16, 1, 17,
}

var _json_lexer_key_offsets []byte = []byte{
	0, 0, 5, 14, 20, 26, 32, 38,
	41, 43, 47, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 58, 63, 72, 78,
	84, 90, 96, 99, 101, 105, 107, 108,
	109, 110, 111, 112, 113, 114, 115, 116,
	126, 130, 132, 137, 153, 157, 159,
}

var _json_lexer_trans_keys []byte = []byte{
	34, 92, 127, 0, 31, 34, 47, 92,
	98, 102, 110, 114, 116, 117, 48, 57,
	65, 70, 97, 102, 48, 57, 65, 70,
	97, 102, 48, 57, 65, 70, 97, 102,
	48, 57, 65, 70, 97, 102, 46, 48,
	57, 48, 57, 43, 45, 48, 57, 48,
	57, 97, 108, 115, 101, 117, 108, 108,
	114, 117, 34, 92, 127, 0, 31, 34,
	47, 92, 98, 102, 110, 114, 116, 117,
	48, 57, 65, 70, 97, 102, 48, 57,
	65, 70, 97, 102, 48, 57, 65, 70,
	97, 102, 48, 57, 65, 70, 97, 102,
	46, 48, 57, 48, 57, 43, 45, 48,
	57, 48, 57, 97, 108, 115, 101, 117,
	108, 108, 114, 117, 34, 43, 45, 46,
	91, 102, 110, 116, 48, 57, 69, 101,
	48, 57, 48, 57, 46, 69, 101, 48,
	57, 9, 32, 34, 44, 46, 91, 93,
	102, 110, 116, 11, 12, 43, 45, 48,
	57, 69, 101, 48, 57, 48, 57, 46,
	69, 101, 48, 57,
}

var _json_lexer_single_lengths []byte = []byte{
	0, 3, 9, 0, 0, 0, 0, 1,
	0, 2, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 9, 0, 0,
	0, 0, 1, 0, 2, 0, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 8,
	2, 0, 3, 10, 2, 0, 3,
}

var _json_lexer_range_lengths []byte = []byte{
	0, 1, 0, 3, 3, 3, 3, 1,
	1, 1, 1, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 3,
	3, 3, 1, 1, 1, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1,
	1, 1, 1, 3, 1, 1, 1,
}

var _json_lexer_index_offsets []byte = []byte{
	0, 0, 5, 15, 19, 23, 27, 31,
	34, 36, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 65, 75, 79,
	83, 87, 91, 94, 96, 100, 102, 104,
	106, 108, 110, 112, 114, 116, 118, 120,
	130, 134, 136, 141, 155, 159, 161,
}

var _json_lexer_indicies []byte = []byte{
	2, 3, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 4, 1, 5,
	5, 5, 1, 6, 6, 6, 1, 7,
	7, 7, 1, 0, 0, 0, 1, 8,
	9, 1, 10, 1, 12, 12, 13, 11,
	13, 11, 14, 1, 15, 1, 16, 1,
	17, 1, 18, 1, 19, 1, 20, 1,
	21, 1, 16, 1, 23, 24, 1, 1,
	22, 22, 22, 22, 22, 22, 22, 22,
	22, 25, 1, 26, 26, 26, 1, 27,
	27, 27, 1, 28, 28, 28, 1, 22,
	22, 22, 1, 29, 30, 1, 31, 1,
	33, 33, 34, 32, 34, 32, 35, 1,
	36, 1, 37, 1, 38, 1, 39, 1,
	40, 1, 41, 1, 42, 1, 37, 1,
	0, 43, 43, 8, 44, 45, 46, 47,
	9, 1, 49, 49, 10, 48, 13, 48,
	10, 49, 49, 9, 48, 50, 50, 22,
	52, 29, 53, 54, 55, 56, 57, 50,
	51, 30, 1, 59, 59, 31, 58, 34,
	58, 31, 59, 59, 30, 58,
}

var _json_lexer_trans_targs []byte = []byte{
	1, 0, 39, 2, 3, 4, 5, 6,
	8, 42, 40, 39, 10, 41, 12, 13,
	14, 39, 16, 17, 39, 19, 20, 43,
	21, 22, 23, 24, 25, 27, 46, 44,
	43, 29, 45, 31, 32, 33, 43, 35,
	36, 43, 38, 7, 39, 11, 15, 18,
	39, 9, 43, 26, 43, 43, 43, 30,
	34, 37, 43, 28,
}

var _json_lexer_trans_actions []byte = []byte{
	0, 0, 25, 0, 0, 0, 0, 0,
	0, 5, 5, 35, 0, 0, 0, 0,
	0, 27, 0, 0, 29, 0, 0, 15,
	0, 0, 0, 0, 0, 0, 5, 5,
	23, 0, 0, 0, 0, 0, 17, 0,
	0, 19, 0, 0, 31, 0, 0, 0,
	33, 0, 13, 0, 11, 7, 9, 0,
	0, 0, 21, 0,
}

var _json_lexer_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1,
	0, 0, 0, 1, 0, 0, 0,
}

var _json_lexer_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 3,
	0, 0, 0, 3, 0, 0, 0,
}

var _json_lexer_eof_trans []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 12, 12, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 33, 33, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	49, 49, 49, 0, 59, 59, 59,
}

const json_lexer_start int = 39
const json_lexer_first_final int = 39
const json_lexer_error int = 0

const json_lexer_en_j_array int = 43
const json_lexer_en_main int = 39


//.... json.rl:21


func (lex *JsonLexer) initJsonLexer() {

//.... json.go:166
	{
	 lex.cs = json_lexer_start
	 lex.top = 0
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//.... json.rl:25
}

func (lex *JsonLexer) Lex() *token.Token {
    eof := lex.pe
    var tok token.ID
    tkn := lex.tokenPool.GenBlock()


//.... json.go:184
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	if  lex.cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_json_lexer_from_state_actions[ lex.cs])
	_nacts = uint(_json_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _json_lexer_actions[_acts - 1] {
		case 1:
//.... NONE:1
 lex.ts = ( lex.p)

//.... json.go:207
		}
	}

	_keys = int(_json_lexer_key_offsets[ lex.cs])
	_trans = int(_json_lexer_index_offsets[ lex.cs])

	_klen = int(_json_lexer_single_lengths[ lex.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case  lex.data[( lex.p)] < _json_lexer_trans_keys[_mid]:
				_upper = _mid - 1
			case  lex.data[( lex.p)] > _json_lexer_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_json_lexer_range_lengths[ lex.cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case  lex.data[( lex.p)] < _json_lexer_trans_keys[_mid]:
				_upper = _mid - 2
			case  lex.data[( lex.p)] > _json_lexer_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_json_lexer_indicies[_trans])
_eof_trans:
	 lex.cs = int(_json_lexer_trans_targs[_trans])

	if _json_lexer_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_json_lexer_trans_actions[_trans])
	_nacts = uint(_json_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _json_lexer_actions[_acts-1] {
		case 2:
//.... NONE:1
 lex.te = ( lex.p)+1

		case 3:
//.... json.rl:86
 lex.te = ( lex.p)+1
{
                lex.bracketStack += 1
                lex.pushSubTokenStack(token.J_ARRAY, lex.ts)
            }
		case 4:
//.... json.rl:90
 lex.te = ( lex.p)+1
{
                lex.bracketStack -= 1
                lex.popSubTokenStack(tkn, lex.te)

                if lex.bracketStack == 0 {
                     lex.top--;  lex.cs =  lex.stack[ lex.top]
{
        /* stack shrinking code */
     }
goto _again

                }
            }
		case 5:
//.... json.rl:98
 lex.te = ( lex.p)+1
{ }
		case 6:
//.... json.rl:99
 lex.te = ( lex.p)+1
{ }
		case 7:
//.... json.rl:101
 lex.te = ( lex.p)+1
{ lex.addSubToken(tkn, token.J_STRING, lex.ts, lex.te) }
		case 8:
//.... json.rl:102
 lex.te = ( lex.p)+1
{ lex.addSubToken(tkn, token.J_BOOL, lex.ts, lex.te) }
		case 9:
//.... json.rl:103
 lex.te = ( lex.p)+1
{ lex.addSubToken(tkn, token.J_NULL, lex.ts, lex.te) }
		case 10:
//.... json.rl:100
 lex.te = ( lex.p)
( lex.p)--
{ lex.addSubToken(tkn, token.J_NUMBER, lex.ts, lex.te) }
		case 11:
//.... json.rl:100
( lex.p) = ( lex.te) - 1
{ lex.addSubToken(tkn, token.J_NUMBER, lex.ts, lex.te) }
		case 12:
//.... json.rl:112
 lex.te = ( lex.p)+1
{ tok = token.J_STRING; ( lex.p)++; goto _out
 }
		case 13:
//.... json.rl:113
 lex.te = ( lex.p)+1
{ tok = token.J_BOOL; ( lex.p)++; goto _out
 }
		case 14:
//.... json.rl:114
 lex.te = ( lex.p)+1
{ tok = token.J_NULL; ( lex.p)++; goto _out
 }
		case 15:
//.... json.rl:115
 lex.te = ( lex.p)+1
{
                tok = token.J_ARRAY;
                lex.bracketStack += 1;
                {
        /* stack growing code */
     lex.stack[ lex.top] =  lex.cs;  lex.top++;  lex.cs = 43; goto _again
 }
                ( lex.p)++; goto _out

            }
		case 16:
//.... json.rl:111
 lex.te = ( lex.p)
( lex.p)--
{ tok = token.J_NUMBER; ( lex.p)++; goto _out
 }
		case 17:
//.... json.rl:111
( lex.p) = ( lex.te) - 1
{ tok = token.J_NUMBER; ( lex.p)++; goto _out
 }
//.... json.go:372
		}
	}

_again:
	_acts = int(_json_lexer_to_state_actions[ lex.cs])
	_nacts = uint(_json_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _json_lexer_actions[_acts-1] {
		case 0:
//.... NONE:1
 lex.ts = 0

//.... json.go:386
		}
	}

	if  lex.cs == 0 {
		goto _out
	}
	( lex.p)++
	if ( lex.p) != ( lex.pe) {
		goto _resume
	}
	_test_eof: {}
	if ( lex.p) == eof {
		if _json_lexer_eof_trans[ lex.cs] > 0 {
			_trans = int(_json_lexer_eof_trans[ lex.cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//.... json.rl:124


    tkn.Value = lex.data[lex.ts:lex.te]
    tkn.ID = token.ID(tok)

    return tkn
}
