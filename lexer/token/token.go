package token

import "fmt"

type Token struct {
	Type   string
	Lexeme string
	Lieral interface{}
	Line   int
}

func New(tokenType, lexeme string, literal interface{}, line int) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
		Lieral: literal,
		Line:   line,
	}
}

func (t Token) ToString() {
	fmt.Printf("%q %q %q", t.Type, t.Lexeme, t.Lieral)
}
