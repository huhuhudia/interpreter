package lexer

import (
	"fmt"
	"github.com/huhuhudia/interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T){
	input := `let five = 5;
		let ten = 10;
		let add = fn(x, y){
			x + y;




		};
		let result = add(five, ten);  
		if a != 'c' {
			return 6
		}	
		
`

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.IDENT, "a"},
		{token.NOT_EQ, "!="},
		{token.CHAR, `'c'`},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "6"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests{
		tok := l.NextToken()
		fmt.Println(tok)
		if tok.Type != tt.expectedType{
			t.Fatalf("test[%d] - tokentype wrong, expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
