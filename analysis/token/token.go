package token

import "strings"

type TokenType string

type Token struct {
	Type      TokenType
	Literal   string
	Start     int
	End       int
	Line      int
	LineStart int
	LineEnd   int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LABEL    = "LABEL"
	REGISTER = "REGISTER"
	INT      = "INT"
	HEX      = "HEX"

	COMMA     = ","
	PERIOD    = "."
	COLON     = ":"
	SEMICOLON = ";"
	HASHTAG   = "#"

	OPCODE    = "OPCODE"
	DIRECTIVE = "DIRECTIVE"
	TRAP      = "TRAP"

	COMMENT = "COMMENT"
)

var opcodes = map[string]TokenType{
	"ADD": OPCODE,
	"NOT": OPCODE,
	"AND": OPCODE,

	"LD":  OPCODE,
	"LDI": OPCODE,
	"LDR": OPCODE,
	"LEA": OPCODE,

	"ST":  OPCODE,
	"STR": OPCODE,
	"STI": OPCODE,

	"JMP":  OPCODE,
	"JSR":  OPCODE,
	"JSRR": OPCODE,

	"RTI": OPCODE,
	"RET": OPCODE,

	//"BR": OPCODE,

}

var directives = map[string]TokenType{
	"FILL":    DIRECTIVE,
	"BLKW":    DIRECTIVE,
	"STRINGZ": DIRECTIVE,

	"ORIG": DIRECTIVE,
	"END":  DIRECTIVE,
}

var traps = map[string]TokenType{
	"HALT":  TRAP,
	"PUTS":  TRAP,
	"GETC":  TRAP,
	"IN":    TRAP,
	"PUTSP": TRAP,
	"OUT":   TRAP,
}

// TODO: Add this function
func LookupIdent(ident string) TokenType {
	if tok, ok := opcodes[strings.ToUpper(ident)]; ok {
		return tok
	}
	if tok, ok := traps[strings.ToUpper(ident)]; ok {
		return tok
	}
	if tok, ok := directives[strings.ToUpper(ident)]; ok {
		return tok
	}
	if len(ident) == 2 {
		if ident[0] == 'r' || ident[0] == 'R' {
			if '0' <= ident[1] && ident[1] <= '9' {
				return REGISTER
			}
		}
	}
	return LABEL
}
