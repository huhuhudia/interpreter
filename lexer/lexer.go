package lexer

import (
	"fmt"
	"github.com/huhuhudia/interpreter/token"
)

type Lexer struct {
	input string
	position int
	ch byte
}

func New(input string)  *Lexer{
	l := &Lexer{input:input}
	return l
}


func (l *Lexer)BackWard(){
	l.position -= 1
}


func isLetter(ch byte)  bool{
	return 'a'<= ch && ch <= 'z' || 'A' <= ch && ch<= 'Z' || ch == '_'
}

func (l *Lexer)isOutOfRange() bool{
	return l.position >= len(l.input)
}

func (l *Lexer)NextToken() token.Token {
	var tok token.Token
	// out of range
	// discard whitespace
	for !l.isOutOfRange() &&  discard(l.GetCurrentChar()){
		l.Forward()
	}
	if l.isOutOfRange(){
		tok.Literal = ""
		tok.Type = token.EOF
		return tok
	}
	// get indent
	if isLetter(l.GetCurrentChar()){
		startPos := l.position
		for isAlphabet(l.GetCurrentChar()){
			l.Forward()
		}
		tok.Literal = l.input[startPos:l.position]
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	}
	// get int
	if isDigit(l.GetCurrentChar()){
		startPos := l.position
		for  isDigit(l.GetCurrentChar()){
			l.Forward()
		}
		tok.Literal = l.input[startPos:l.position]
		tok.Type = token.INT
		return tok
	}
	// get string
	if l.GetCurrentChar() == '"'{
		startPos := l.position
		l.Forward()
		for !l.isOutOfRange() && l.GetCurrentChar() != '"'{
			l.Forward()
		}
		if l.isOutOfRange(){
			tok.Type = token.ILLEGAL
			tok.Literal = token.ILLEGAL
		}
		l.Forward()
		fmt.Println()
		val := l.input[startPos:l.position]
		tok.Literal = val
		tok.Type = token.STRING
		return tok
	}


	if l.GetCurrentChar() == '\''{
		startPos := l.position
		l.Forward()
		for !l.isOutOfRange() && l.GetCurrentChar() != '\''{
			l.Forward()
		}
		if l.isOutOfRange(){
			tok.Type = token.ILLEGAL
			tok.Literal = token.ILLEGAL
		}
		l.Forward()
		fmt.Println()
		val := l.input[startPos:l.position]
		tok.Literal = val
		tok.Type = token.CHAR
		return tok
	}


	switch l.GetCurrentChar() {
	case '=':
		if l.nextChar() == '=' {
			tok.Type = token.EQ
			tok.Literal = "=="
			l.Forward()
			l.Forward()
			return tok
		}else{
			tok = newToken(token.ASSIGN, l.GetCurrentChar())
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.GetCurrentChar())
	case '(':
		tok = newToken(token.LPAREN, l.GetCurrentChar())
	case ')':
		tok = newToken(token.RPAREN, l.GetCurrentChar())
	case ',':
		tok = newToken(token.COMMA, l.GetCurrentChar())
	case '+':
		tok = newToken(token.PLUS, l.GetCurrentChar())
	case '{':
		tok = newToken(token.LBRACE, l.GetCurrentChar())
	case '}':
		tok = newToken(token.RBRACE, l.GetCurrentChar())
	case '!':
		if l.nextChar() == '='{
			tok.Type = token.NOT_EQ
			tok.Literal = "!="
			l.Forward()
			l.Forward()
			return tok
		}else{
			tok = newToken(token.BANG, l.GetCurrentChar())
		}
	case '-':
		tok = newToken(token.MINUS, l.GetCurrentChar())
	case '*':
		tok = newToken(token.MINUS, l.GetCurrentChar())
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.Forward()
	return tok
}

func (l *Lexer) GetCurrentChar() byte{
	return l.input[l.position]
}

func (l *Lexer)Forward(){
	fmt.Println("[forward]" +"[" + string(l.GetCurrentChar()) + "]")
	l.position += 1
}

func isAlphabet(ch byte) bool{
	return ch >= 'a'  && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func discard(ch byte) bool{
	for ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'{
		return true
	}
	return false
}

func isDigit(ch byte) bool{
	return ch <= '9' && ch >= '0'
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type:tokenType, Literal:string(ch)}
}

func (l *Lexer) nextChar() byte{
	l.Forward()
	if l.isOutOfRange(){
		return ' '
	}
	res := l.GetCurrentChar()
	l.BackWard()
	return res
}