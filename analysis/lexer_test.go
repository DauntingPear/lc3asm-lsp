package analysis

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `ADD R0,R0,#1
	LABEL:
	NOT R0,R1
	BRnz JumpLabel

	`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{OPCODE, "ADD"},
		{REGISTER, "R0"},
		{COMMA, ","},
		{REGISTER, "R0"},
		{COMMA, ","},
		{HASHTAG, "#"},
		{INT, "1"},
		{OPCODE, "LABEL"},
		{COLON, ":"},
		{OPCODE, "NOT"},
		{REGISTER, "R0"},
		{COMMA, ","},
		{REGISTER, "R1"},
		{OPCODE, "BRnz"},
		{OPCODE, "JumpLabel"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Printf("%v", tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
