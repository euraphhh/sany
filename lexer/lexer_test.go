package lexer

import (
	"sany/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `var five = 5;
const ten = 10;

func add(x, y) {
  var result = x + y;
  return result;
}

var result = add(five, ten);
!-/*5;
5 < 10 > 5;

if 5 < 10 {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;

class User {
  var name
}
agent Bot {
  model = 1
}
task analyze
tool search
import
export func foo
error null
and or not
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.CONST, "const"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNC, "func"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.IDENT, "result"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		
		// Sany specific tokens
		{token.CLASS, "class"},
		{token.IDENT, "User"},
		{token.LBRACE, "{"},
		{token.VAR, "var"},
		{token.IDENT, "name"},
		{token.RBRACE, "}"},
		{token.AGENT, "agent"},
		{token.IDENT, "Bot"},
		{token.LBRACE, "{"},
		{token.IDENT, "model"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.RBRACE, "}"},
		{token.TASK, "task"},
		{token.IDENT, "analyze"},
		{token.TOOL, "tool"},
		{token.IDENT, "search"},
		{token.IMPORT, "import"},
		{token.EXPORT, "export"},
		{token.FUNC, "func"},
		{token.IDENT, "foo"},
		{token.ERROR, "error"},
		{token.NULL, "null"},
		{token.AND, "and"},
		{token.OR, "or"},
		{token.NOT, "not"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
