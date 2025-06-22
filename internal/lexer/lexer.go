package lexer

import myToken "github.com/pcristin/monkry_interpet/internal/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() myToken.Token {
	var tok myToken.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(myToken.ASSIGN, l.ch)
	case ';':
		tok = newToken(myToken.SEMICOLON, l.ch)
	case '(':
		tok = newToken(myToken.LPAREN, l.ch)
	case ')':
		tok = newToken(myToken.RPAREN, l.ch)
	case ',':
		tok = newToken(myToken.COMMA, l.ch)
	case '+':
		tok = newToken(myToken.PLUS, l.ch)
	case '{':
		tok = newToken(myToken.LBRACE, l.ch)
	case '}':
		tok = newToken(myToken.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = myToken.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = myToken.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = myToken.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(myToken.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokentype myToken.TokenType, ch byte) myToken.Token {
	return myToken.Token{Type: tokentype,
		Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
