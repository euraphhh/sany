package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	FLOAT  = "FLOAT" // 1.23
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	VAR      = "VAR"
	CONST    = "CONST"
	FUNC     = "FUNC"
	CLASS    = "CLASS"
	AGENT    = "AGENT"
	TASK     = "TASK"
	TOOL     = "TOOL"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	IN       = "IN"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	RETURN   = "RETURN"
	AND      = "AND"
	OR       = "OR"
	NOT      = "NOT"
	NULL     = "NULL"
	SELF     = "SELF"
	IMPORT   = "IMPORT"
	EXPORT   = "EXPORT"
	ERROR    = "ERROR"
)

var keywords = map[string]TokenType{
	"var":      VAR,
	"const":    CONST,
	"func":     FUNC,
	"class":    CLASS,
	"agent":    AGENT,
	"task":     TASK,
	"tool":     TOOL,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"in":       IN,
	"break":    BREAK,
	"continue": CONTINUE,
	"return":   RETURN,
	"and":      AND,
	"or":       OR,
	"not":      NOT,
	"null":     NULL,
	"self":     SELF,
	"import":   IMPORT,
	"export":   EXPORT,
	"error":    ERROR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
