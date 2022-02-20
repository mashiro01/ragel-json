
//line json.rl:1
package scanner

import (
    "main/token"
)


//line json.go:11
var _json_lexer_actions []byte = []byte{
	0, 1, 0, 1, 2, 1, 3, 1, 5, 
	1, 6, 1, 7, 1, 8, 1, 9, 
	1, 10, 1, 11, 1, 12, 1, 13, 
	1, 14, 1, 15, 1, 16, 1, 18, 
	1, 19, 1, 20, 1, 21, 1, 22, 
	1, 23, 1, 24, 1, 25, 1, 26, 
	1, 27, 1, 28, 1, 29, 1, 30, 
	1, 31, 2, 0, 1, 2, 3, 4, 
	2, 3, 17, 
}

var _json_lexer_cond_offsets []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	1, 1, 1, 1, 1, 1, 
}

var _json_lexer_cond_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 1, 
	0, 0, 0, 0, 0, 0, 
}

var _json_lexer_cond_keys []int16 = []int16{
	44, 44, 
}

var _json_lexer_cond_spaces []byte = []byte{
	0, 
}

var _json_lexer_key_offsets []byte = []byte{
	0, 0, 5, 14, 20, 26, 32, 38, 
	41, 43, 47, 49, 50, 51, 52, 53, 
	54, 55, 56, 57, 58, 59, 64, 73, 
	79, 85, 91, 97, 100, 102, 106, 108, 
	109, 110, 111, 112, 113, 114, 115, 116, 
	117, 118, 123, 128, 137, 143, 149, 155, 
	161, 162, 163, 165, 176, 180, 182, 187, 
	205, 209, 213, 215, 220, 229, 
}

var _json_lexer_trans_keys []int16 = []int16{
	34, 92, 127, 0, 31, 34, 47, 92, 
	98, 102, 110, 114, 116, 117, 48, 57, 
	65, 70, 97, 102, 48, 57, 65, 70, 
	97, 102, 48, 57, 65, 70, 97, 102, 
	48, 57, 65, 70, 97, 102, 46, 48, 
	57, 48, 57, 43, 45, 48, 57, 48, 
	57, 97, 108, 115, 101, 117, 108, 108, 
	114, 117, 10, 34, 92, 127, 0, 31, 
	34, 47, 92, 98, 102, 110, 114, 116, 
	117, 48, 57, 65, 70, 97, 102, 48, 
	57, 65, 70, 97, 102, 48, 57, 65, 
	70, 97, 102, 48, 57, 65, 70, 97, 
	102, 46, 48, 57, 48, 57, 43, 45, 
	48, 57, 48, 57, 97, 108, 115, 101, 
	117, 108, 108, 114, 117, 10, 34, 92, 
	127, 0, 31, 9, 32, 58, 11, 12, 
	34, 47, 92, 98, 102, 110, 114, 116, 
	117, 48, 57, 65, 70, 97, 102, 48, 
	57, 65, 70, 97, 102, 48, 57, 65, 
	70, 97, 102, 48, 57, 65, 70, 97, 
	102, 42, 42, 42, 47, 34, 43, 45, 
	46, 91, 102, 110, 116, 123, 48, 57, 
	69, 101, 48, 57, 48, 57, 46, 69, 
	101, 48, 57, 13, 32, 34, 43, 45, 
	46, 91, 93, 102, 110, 116, 123, 300, 
	556, 9, 12, 48, 57, 13, 32, 9, 
	12, 69, 101, 48, 57, 48, 57, 46, 
	69, 101, 48, 57, 13, 32, 34, 44, 
	47, 123, 125, 9, 12, 13, 32, 9, 
	12, 
}

var _json_lexer_single_lengths []byte = []byte{
	0, 3, 9, 0, 0, 0, 0, 1, 
	0, 2, 0, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 3, 9, 0, 
	0, 0, 0, 1, 0, 2, 0, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 3, 3, 9, 0, 0, 0, 0, 
	1, 1, 2, 9, 2, 0, 3, 14, 
	2, 2, 0, 3, 7, 2, 
}

var _json_lexer_range_lengths []byte = []byte{
	0, 1, 0, 3, 3, 3, 3, 1, 
	1, 1, 1, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 3, 
	3, 3, 3, 1, 1, 1, 1, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 1, 0, 3, 3, 3, 3, 
	0, 0, 0, 1, 1, 1, 1, 2, 
	1, 1, 1, 1, 1, 1, 
}

var _json_lexer_index_offsets []int16 = []int16{
	0, 0, 5, 15, 19, 23, 27, 31, 
	34, 36, 40, 42, 44, 46, 48, 50, 
	52, 54, 56, 58, 60, 62, 67, 77, 
	81, 85, 89, 93, 96, 98, 102, 104, 
	106, 108, 110, 112, 114, 116, 118, 120, 
	122, 124, 129, 134, 144, 148, 152, 156, 
	160, 162, 164, 167, 178, 182, 184, 189, 
	206, 210, 214, 216, 221, 230, 
}

var _json_lexer_indicies []byte = []byte{
	2, 3, 1, 1, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 4, 1, 5, 
	5, 5, 1, 6, 6, 6, 1, 7, 
	7, 7, 1, 0, 0, 0, 1, 8, 
	9, 1, 10, 1, 12, 12, 13, 11, 
	13, 11, 14, 1, 15, 1, 16, 1, 
	17, 1, 18, 1, 19, 1, 20, 1, 
	21, 1, 16, 1, 23, 22, 25, 26, 
	1, 1, 24, 24, 24, 24, 24, 24, 
	24, 24, 24, 27, 1, 28, 28, 28, 
	1, 29, 29, 29, 1, 30, 30, 30, 
	1, 24, 24, 24, 1, 31, 32, 1, 
	33, 1, 35, 35, 36, 34, 36, 34, 
	37, 1, 38, 1, 39, 1, 40, 1, 
	41, 1, 42, 1, 43, 1, 44, 1, 
	39, 1, 46, 45, 48, 49, 1, 1, 
	47, 48, 48, 50, 48, 1, 47, 47, 
	47, 47, 47, 47, 47, 47, 51, 1, 
	52, 52, 52, 1, 53, 53, 53, 1, 
	54, 54, 54, 1, 47, 47, 47, 1, 
	55, 1, 56, 55, 56, 57, 55, 0, 
	58, 58, 8, 59, 60, 61, 62, 63, 
	9, 1, 65, 65, 10, 64, 13, 64, 
	10, 65, 65, 9, 64, 66, 23, 24, 
	67, 67, 31, 68, 69, 70, 71, 72, 
	73, 74, 75, 23, 32, 1, 66, 23, 
	23, 76, 78, 78, 33, 77, 36, 77, 
	33, 78, 78, 32, 77, 79, 46, 47, 
	80, 81, 82, 83, 46, 1, 79, 46, 
	46, 84, 
}

var _json_lexer_trans_targs []byte = []byte{
	1, 0, 51, 2, 3, 4, 5, 6, 
	8, 54, 52, 51, 10, 53, 12, 13, 
	14, 51, 16, 17, 51, 19, 55, 56, 
	21, 55, 22, 23, 24, 25, 26, 28, 
	59, 57, 55, 30, 58, 32, 33, 34, 
	55, 36, 37, 55, 39, 60, 61, 41, 
	42, 43, 60, 44, 45, 46, 47, 49, 
	50, 60, 7, 51, 11, 15, 18, 51, 
	51, 9, 20, 27, 55, 55, 31, 35, 
	38, 55, 55, 55, 55, 55, 29, 40, 
	60, 48, 60, 60, 60, 
}

var _json_lexer_trans_actions []byte = []byte{
	0, 0, 45, 0, 0, 0, 0, 0, 
	0, 5, 5, 57, 0, 0, 0, 0, 
	0, 47, 0, 0, 49, 0, 29, 62, 
	0, 15, 0, 0, 0, 0, 0, 0, 
	5, 5, 27, 0, 0, 0, 0, 0, 
	17, 0, 0, 19, 0, 43, 65, 0, 
	0, 0, 39, 0, 0, 0, 0, 0, 
	0, 37, 0, 51, 0, 0, 0, 53, 
	55, 0, 0, 0, 7, 9, 0, 0, 
	0, 21, 13, 11, 23, 25, 0, 0, 
	35, 0, 31, 33, 41, 
}

var _json_lexer_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 1, 0, 0, 0, 59, 
	0, 0, 0, 0, 59, 0, 
}

var _json_lexer_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 3, 0, 0, 0, 3, 
	0, 0, 0, 0, 3, 0, 
}

var _json_lexer_eof_trans []int16 = []int16{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 12, 12, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 23, 0, 0, 0, 
	0, 0, 0, 0, 0, 35, 35, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	46, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 65, 65, 65, 0, 
	77, 78, 78, 78, 0, 85, 
}

const json_lexer_start int = 51
const json_lexer_first_final int = 51
const json_lexer_error int = 0

const json_lexer_en_j_array int = 55
const json_lexer_en_j_object int = 60
const json_lexer_en_main int = 51


//line json.rl:17


func (lex *JsonLexer) initJsonLexer() {
    
//line json.go:239
	{
	 lex.cs = json_lexer_start
	 lex.top = 0
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

//line json.rl:21
}

func (lex *JsonLexer) Lex() *token.Token {
    eof := lex.pe
    var tok token.ID
    tkn := lex.tokenPool.GenBlock()

    
//line json.go:257
	{
	var _klen int
	var _trans int
	var _widec int16
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
		case 2:
//line NONE:1
 lex.ts = ( lex.p)

//line json.go:281
		}
	}

	_widec = int16( lex.data[( lex.p)])
	_klen = int(_json_lexer_cond_lengths[ lex.cs])
	_keys = int(_json_lexer_cond_offsets[ lex.cs] * 2)
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
	COND_LOOP:
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case _widec < int16(_json_lexer_cond_keys[_mid]):
				_upper = _mid - 2
			case _widec > int16(_json_lexer_cond_keys[_mid + 1]):
				_lower = _mid + 2
			default:
				switch _json_lexer_cond_spaces[int(_json_lexer_cond_offsets[ lex.cs]) + ((_mid - _keys)>>1)] {
				case 0:
					_widec = 256 + (int16( lex.data[( lex.p)]) - 0)
					if  lex.isNotInArrayParse()  {
						_widec += 256
					}
				}
				break COND_LOOP
			}
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
			case _widec < _json_lexer_trans_keys[_mid]:
				_upper = _mid - 1
			case _widec > _json_lexer_trans_keys[_mid]:
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
			case _widec < _json_lexer_trans_keys[_mid]:
				_upper = _mid - 2
			case _widec > _json_lexer_trans_keys[_mid + 1]:
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
		case 3:
//line NONE:1
 lex.te = ( lex.p)+1

		case 4:
//line json.rl:90
 lex.act = 5;
		case 5:
//line json.rl:72
 lex.te = ( lex.p)+1
{
                lex.notInArray = false
                lex.bracketStack += 1
                lex.pushSubTokenStack(token.J_ARRAY, lex.ts)
            }
		case 6:
//line json.rl:77
 lex.te = ( lex.p)+1
{
                lex.bracketStack -= 1
                lex.popSubTokenStack(tkn, lex.te)

                if lex.bracketStack == 0 {
                     lex.top--;  lex.cs =  lex.stack[ lex.top]
goto _again

                }
            }
		case 7:
//line json.rl:85
 lex.te = ( lex.p)+1
{
                lex.popSubTokenStack(tkn, lex.te-1)
                 lex.top--;  lex.cs =  lex.stack[ lex.top]
goto _again

            }
		case 8:
//line json.rl:89
 lex.te = ( lex.p)+1
{  }
		case 9:
//line json.rl:92
 lex.te = ( lex.p)+1
{ lex.addSubToken(tkn, token.J_STRING, lex.ts, lex.te) }
		case 10:
//line json.rl:93
 lex.te = ( lex.p)+1
{ lex.addSubToken(tkn, token.J_BOOL, lex.ts, lex.te) }
		case 11:
//line json.rl:94
 lex.te = ( lex.p)+1
{ lex.addSubToken(tkn, token.J_NULL, lex.ts, lex.te) }
		case 12:
//line json.rl:95
 lex.te = ( lex.p)+1
{  lex.top--;  lex.cs =  lex.stack[ lex.top]
goto _again
 }
		case 13:
//line json.rl:90
 lex.te = ( lex.p)
( lex.p)--
{ }
		case 14:
//line json.rl:91
 lex.te = ( lex.p)
( lex.p)--
{ lex.addSubToken(tkn, token.J_NUMBER, lex.ts, lex.te) }
		case 15:
//line json.rl:91
( lex.p) = ( lex.te) - 1
{ lex.addSubToken(tkn, token.J_NUMBER, lex.ts, lex.te) }
		case 16:
//line NONE:1
	switch  lex.act {
	case 0:
	{ lex.cs = 0
goto _again
}
	case 5:
	{( lex.p) = ( lex.te) - 1
 }
	}
	
		case 17:
//line json.rl:115
 lex.act = 14;
		case 18:
//line json.rl:102
 lex.te = ( lex.p)+1
{
                lex.braceStack += 1
                lex.pushSubTokenStack(token.J_OBJECT, lex.ts)
            }
		case 19:
//line json.rl:106
 lex.te = ( lex.p)+1
{
                lex.braceStack -= 1
                lex.popSubTokenStack(tkn, lex.te)

                if lex.braceStack == 0 {
                     lex.top--;  lex.cs =  lex.stack[ lex.top]
goto _again

                }
            }
		case 20:
//line json.rl:114
 lex.te = ( lex.p)+1
{ }
		case 21:
//line json.rl:116
 lex.te = ( lex.p)+1
{
                lex.addSubToken(tkn, token.J_COMMENT, lex.ts, lex.te)
            }
		case 22:
//line json.rl:119
 lex.te = ( lex.p)+1
{
                lex.pushSubTokenStack(token.J_OBJECT_KEY_VALUE_PAIR, lex.ts)
                lex.notInArray = true
                { 
        lex.growCallStack()
     lex.stack[ lex.top] =  lex.cs;  lex.top++;  lex.cs = 55; goto _again
 }
            }
		case 23:
//line json.rl:115
 lex.te = ( lex.p)
( lex.p)--
{ }
		case 24:
//line NONE:1
	switch  lex.act {
	case 0:
	{ lex.cs = 0
goto _again
}
	case 14:
	{( lex.p) = ( lex.te) - 1
 }
	}
	
		case 25:
//line json.rl:129
 lex.te = ( lex.p)+1
{ tok = token.J_STRING; ( lex.p)++; goto _out
 }
		case 26:
//line json.rl:130
 lex.te = ( lex.p)+1
{ tok = token.J_BOOL; ( lex.p)++; goto _out
 }
		case 27:
//line json.rl:131
 lex.te = ( lex.p)+1
{ tok = token.J_NULL; ( lex.p)++; goto _out
 }
		case 28:
//line json.rl:132
 lex.te = ( lex.p)+1
{
                tok = token.J_ARRAY
                lex.bracketStack += 1
                { 
        lex.growCallStack()
     lex.stack[ lex.top] =  lex.cs;  lex.top++;  lex.cs = 55; goto _again
 }
                ( lex.p)++; goto _out

            }
		case 29:
//line json.rl:138
 lex.te = ( lex.p)+1
{
                tok = token.J_OBJECT
                lex.braceStack += 1
                { 
        lex.growCallStack()
     lex.stack[ lex.top] =  lex.cs;  lex.top++;  lex.cs = 60; goto _again
 }
                ( lex.p)++; goto _out

            }
		case 30:
//line json.rl:128
 lex.te = ( lex.p)
( lex.p)--
{ tok = token.J_NUMBER; ( lex.p)++; goto _out
 }
		case 31:
//line json.rl:128
( lex.p) = ( lex.te) - 1
{ tok = token.J_NUMBER; ( lex.p)++; goto _out
 }
//line json.go:581
		}
	}

_again:
	_acts = int(_json_lexer_to_state_actions[ lex.cs])
	_nacts = uint(_json_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _json_lexer_actions[_acts-1] {
		case 0:
//line NONE:1
 lex.ts = 0

		case 1:
//line NONE:1
 lex.act = 0

//line json.go:599
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

//line json.rl:147


    tkn.Value = lex.data[lex.ts:lex.te]
    tkn.ID = token.ID(tok)

    return tkn
}
