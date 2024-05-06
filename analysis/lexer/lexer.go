package lexer

import (
	"educationalsp/analysis/token"
	"fmt"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	oldPosition  int
	line         int
	lineIndex    int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	l.lineIndex--
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.lineIndex++
	l.position = l.readPosition
	l.readPosition++
	if l.ch == '\n' {
		l.line++
		l.lineIndex = -1
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	tok.Start = l.position
	tok.LineStart = l.lineIndex
	l.oldPosition = l.position

	switch l.ch {
	case ',':
		tok = newToken(token.COMMA, l)
	case '.':
		tok = newToken(token.PERIOD, l)
	case ':':
		tok = newToken(token.COLON, l)
	case ';':
		tok.Type = token.COMMENT
		tok.Line = l.line
		tok.Literal = l.readLine()
		tok.End = l.position
		tok.LineEnd = tok.LineStart + l.position - l.oldPosition - 1
		return tok
	case '#':
		fmt.Printf("CASE #\n")
		if isDigit(l.peekChar()) {
			tok.Type = token.INT
			tok.Line = l.line
			tok.Literal = l.readNumber()
			tok.End = l.position
			tok.LineEnd = tok.LineStart + l.position - l.oldPosition - 1
			return tok
		} else {
			tok = newToken(token.HASHTAG, l)
		}
	case 'x':
		if isDigit(l.peekChar()) {
			tok.Type = token.HEX
			tok.Line = l.line
			tok.Literal = l.readHex()
			tok.End = l.position
			tok.LineEnd = tok.LineStart + l.position - l.oldPosition - 1
			return tok
		} else {
			tok.Line = l.line
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.End = l.position
			tok.LineEnd = tok.LineStart + l.position - l.oldPosition - 1
			return tok
		}
	case 0:
		tok.Literal = ""
		tok = newToken(token.EOF, l)
	default:
		if isLetter(l.ch) {
			tok.Line = l.line
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.End = l.position
			tok.LineEnd = tok.LineStart + l.position - l.oldPosition - 1
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Line = l.line
			tok.Literal = l.readNumber()
			tok.End = l.position
			tok.LineEnd = tok.LineStart + l.position - l.oldPosition - 1
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, l *Lexer) token.Token {
	return token.Token{
		Type:      tokenType,
		Literal:   string(l.ch),
		Start:     l.oldPosition,
		End:       l.position,
		LineEnd:   l.lineIndex,
		LineStart: l.lineIndex,
		Line:      l.line,
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isWhitespace(ch byte) bool {
	return ch == '\n' || ch == '\r' || ch == '\t' || ch == ' '
}

func isNewline(ch byte) bool {
	return ch == '\n' || ch == '\r'
}

func isHex(ch byte) bool {
	return '0' <= ch && ch <= '9' || 'A' <= ch && ch <= 'F'
}

func (l *Lexer) readRegister() string {
	position := l.position
	l.readChar()
	l.readChar()
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readHex() string {
	position := l.position

	// HACK: May cause bugs, need to test later
	if l.ch == 'x' {
		l.readChar()
	}

	for isHex(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	// HACK: May cause bugs, need to test later
	if l.ch == '#' {
		l.readChar()
	}

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readLine() string {
	position := l.position

	for !isNewline(l.ch) {
		fmt.Printf("char: %c\n", l.ch)
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == '\n' || l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		// if l.ch == '\n' {
		// 	l.line++
		// }
		l.readChar()

	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) peekChars(lookAhead int) byte {
	if l.readPosition+lookAhead >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition+lookAhead]
	}

}

func (l *Lexer) prevChar() byte {
	fmt.Printf("POS=%d", l.position)
	if l.position <= 0 {
		return 0
	} else {
		fmt.Printf(", CHAR='%c'\n", l.input[l.position-1])
		return l.input[l.position-1]
	}
}

func (l *Lexer) GetPosition() int {
	return l.position
}

func (l *Lexer) GetLine() int {
	return l.line
}
