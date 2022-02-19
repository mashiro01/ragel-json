package token

const (
	J_NUMBER ID = iota
	J_STRING
	J_BOOL
	J_NULL
	J_ARRAY
	J_OBJECT
)

//go:generate stringer -type ID -output ./token_string.go
type ID int

type Token struct {
	ID    ID
	Value []byte

	SubTokens []Token
}

type StackToken struct {
	Start, End int
	Token      *Token
}
