package lexer

import (
	"educationalsp/analysis/token"
	"testing"
)

func TestBasicTokens(t *testing.T) {
	input := `
	,.:;#
	#1
	1
	x1
	R0
	R7
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COMMA, ","},
		{token.PERIOD, "."},
		{token.COLON, ":"},
		{token.SEMICOLON, ";"},
		{token.HASHTAG, "#"},
		{token.INT, "#1"},
		{token.INT, "1"},
		{token.HEX, "x1"},
		{token.REGISTER, "R0"},
		{token.REGISTER, "R7"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
func TestTokenRange(t *testing.T) {
	input := `R0`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedStart   int
		expectedEnd     int
	}{
		{token.REGISTER, "R0", 0, 1},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Start != tt.expectedStart {
			t.Errorf("tests[%d] - start wrong. expected=%d, got=%d", i, tt.expectedStart, tok.Start)
		}
		if tok.End != tt.expectedEnd {
			t.Errorf("tests[%d] - end wrong. expected=%d, got=%d", i, tt.expectedEnd, tok.End)
		}

	}
}

func TestHover(t *testing.T) {
	input := `.orig x3000
START:
	LD R0,BLACK
	STR R0,COLOR
	ADD R0,R0,#1
	ST R0,LABEL

COLOR .FILL 0`
	line := 0
	character := 0

	l := New(input)
	tokens := []token.Token{}
	var tok token.Token

	for tok.Type != token.EOF {
		tok = l.NextToken()
		t.Errorf("TOK=%+v", tok)
		if tok.Type != token.EOF {
			tokens = append(tokens, tok)
		}
	}

	for _, token := range tokens {
		if line == token.Line && token.LineStart <= character && token.LineEnd >= character {
			tok = token
		}
	}

	for _, token := range tokens {
		t.Errorf("line=%d, characteridx=%d, tokenrecieved=%+v\n", line, character, token)

	}
	t.Errorf("lexpos=%d, lexline=%d, lexlineidx=%d", l.position, l.line, l.lineIndex)
	t.Errorf("tokenRecieved=%+v\n\n\n", tok)

	for _, token := range tokens {
		t.Errorf("Type=%q:\n- Literal=%s\n- Start=%d\n- End=%d\n- Line=%d\n- Linestart=%d\n- Lineend=%d\n", token.Type, token.Literal, token.Start, token.End, token.Line, token.LineStart, token.LineEnd)
	}
	t.Errorf("tokenRecieved=%+v\n\n\n", tok)
}
