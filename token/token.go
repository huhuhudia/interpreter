package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

func (self Token) String()string{
	return "Type:" + string(self.Type) + "\tLiteral: " + self.Literal;
}

const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers and literals
	IDENT = "IDENT"
	INT = "INT"


	// Operations
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"


	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"





	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	EQ = "=="
	NOT_EQ = "!="


	STRING = "STRING"
	CHAR = "CHAR"

)

var keywords = map[string]TokenType{
	"fn":FUNCTION,
	"let":LET,
	"true":TRUE,
	"false":FALSE,
	"else":ELSE,
	"return":RETURN,
	"if":IF,
}

func LookupIdent(ident string)TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}



