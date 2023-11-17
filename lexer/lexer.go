package lexer

import t "github.com/spookbench/glox/token"

type Lexer struct {
	Start   int
	Current int
	Line    int
}

func New() *Lexer {
	return &Lexer{}
}

func (l *Lexer) ScanTokens(s string) []t.Token {
	var isAtEnd bool
	var tokens = make([]t.Token, 0)
	for {
		if isAtEnd {
			break
		}
		l.Start = l.Current
		l.scanToken()

	}

	tokens = append(tokens, t.New(t.EOF, "", nil, l.Line))
	return tokens
}

func (l *Lexer) scanToken() {

}
