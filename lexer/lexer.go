package lexer

import (
	t "github.com/spookbench/glox/lexer/token"
)

type Lexer struct {
	Start   int
	Current int
	Line    int
	Source  string
	Tokens  []t.Token
}

func New(s string) *Lexer {
	return &Lexer{
		Source: s,
	}
}

func (l *Lexer) ScanTokens() []t.Token {
	for {
		if l.isAtEnd(l.Source) {
			break
		}
		l.Start = l.Current
		l.scanToken()

	}

	l.Tokens = append(l.Tokens, t.New(t.EOF, "", nil, l.Line))
	return l.Tokens
}

func (l *Lexer) scanToken() {
	c := l.advance()
	switch c {
	case "(":
		l.addToken(t.LEFT_PAREN)
	case ")":
		l.addToken(t.RIGHT_PAREN)
	case "{":
		l.addToken(t.LEFT_BRACE)
	case "}":
		l.addToken(t.RIGHT_BRACE)
	case ",":
		l.addToken(t.COMMA)
	case ".":
		l.addToken(t.DOT)
	case "-":
		l.addToken(t.MINUS)
	case "+":
		l.addToken(t.PLUS)
	case ";":
		l.addToken(t.SEMICOLON)
	case "*":
		l.addToken(t.STAR)
	default:
		panic("Unexpected character.")
		// Lox.error(line, "Unexpected character.");
	}
}

func (l *Lexer) advance() string {
	return string(l.Source[l.Current+1])
}

func (l *Lexer) addToken(tokenType string) {
	text := l.Source[l.Start:l.Current]
	l.Tokens = append(l.Tokens, t.New(tokenType, text, nil, l.Line))
}

func (l *Lexer) isAtEnd(s string) bool {
	return l.Current >= len(s)
}
