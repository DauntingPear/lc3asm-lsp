package analysis

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

/*
LABEL:
    STUFF
ADD R0,R0,#1
NOT R0,R0
ADD R0, R0

VARLABEL .FILL 0
VARLABEL .STRINGZ "HI"
; comment
*/

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LABEL = "LABEL"
	INT   = "INT"
	HEX   = "HEX"

	HASHTAG = "#"

	PERIOD    = "."
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	OPCODE    = "OPCODE"
	REGISTER  = "REGISTER"
	DIRECTIVE = "DIRECTIVE"
)

var opcodes = map[string]TokenType{
	"ADD": OPCODE,
	"NOT": OPCODE,
}

var directives = map[string]TokenType{
	"FILL": DIRECTIVE,
}

func LookupOpcodes(opcode string) TokenType {
	if tok, ok := opcodes[opcode]; ok {
		return tok
	}
	return OPCODE
}

func LookupDirectives(directive string) TokenType {
	if tok, ok := directives[directive]; ok {
		return tok
	}
	return DIRECTIVE
}

func LookupRegisters(register string) TokenType {
	return REGISTER
}
