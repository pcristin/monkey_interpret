package lexer

import (
	"testing"

	myToken "github.com/pcristin/monkry_interpet/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    myToken.TokenType
		expectedLiteral string
	}{
		{myToken.ASSIGN, "="},
		{myToken.PLUS, "+"},
		{myToken.LPAREN, "("},
		{myToken.RPAREN, ")"},
		{myToken.LBRACE, "{"},
		{myToken.RBRACE, "}"},
		{myToken.COMMA, ","},
		{myToken.SEMICOLON, ";"},
		{myToken.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNewTokenHard(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    myToken.TokenType
		expectedLiteral string
	}{
		{myToken.LET, "let"},
		{myToken.IDENT, "five"},
		{myToken.ASSIGN, "="},
		{myToken.INT, "5"},
		{myToken.SEMICOLON, ";"},
		{myToken.LET, "let"},
		{myToken.IDENT, "ten"},
		{myToken.ASSIGN, "="},
		{myToken.INT, "10"},
		{myToken.SEMICOLON, ";"},
		{myToken.LET, "let"},
		{myToken.IDENT, "add"},
		{myToken.ASSIGN, "="},
		{myToken.FUNCTION, "fn"},
		{myToken.LPAREN, "("},
		{myToken.IDENT, "x"},
		{myToken.COMMA, ","},
		{myToken.IDENT, "y"},
		{myToken.RPAREN, ")"},
		{myToken.LBRACE, "{"},
		{myToken.IDENT, "x"},
		{myToken.PLUS, "+"},
		{myToken.IDENT, "y"},
		{myToken.SEMICOLON, ";"},
		{myToken.RBRACE, "}"},
		{myToken.SEMICOLON, ";"},
		{myToken.LET, "let"},
		{myToken.IDENT, "result"},
		{myToken.ASSIGN, "="},
		{myToken.IDENT, "add"},
		{myToken.LPAREN, "("},
		{myToken.IDENT, "five"},
		{myToken.COMMA, ","},
		{myToken.IDENT, "ten"},
		{myToken.RPAREN, ")"},
		{myToken.SEMICOLON, ";"},
		{myToken.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
