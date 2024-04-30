package analysis

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
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
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhiteSpace()

	switch l.ch {
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '#':
		tok = newToken(HASHTAG, l.ch)
	case ':':
		tok = newToken(COLON, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '.':
		tok = newToken(PERIOD, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if l.isRegister() {
			tok.Literal = l.readRegister()
			tok.Type = LookupRegisters(tok.Literal)
			return tok
		}
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupOpcodes(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) isRegister() bool {
	if l.ch == 'R' && isDigit(l.peekChar(1)) && (l.peekChar(2) == ',' || l.peekChar(2) == '\n') {
		return true
	}
	return false
}

func (l *Lexer) readRegister() string {
	position := l.position
	l.readChar()
	l.readChar()
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar(ahead int) byte {
	if l.position+ahead >= len(l.input) {
		return 0
	} else {
		return l.input[l.position+ahead]
	}
}
