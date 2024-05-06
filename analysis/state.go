package analysis

import (
	"educationalsp/analysis/lexer"
	"educationalsp/analysis/token"
	"educationalsp/lsp"
	"fmt"
)

type State struct {
	Documents map[string]string
	Lexer     *lexer.Lexer
}

func NewState() State {
	l := lexer.New("")
	return State{
		Documents: map[string]string{},
		Lexer:     l,
	}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// in a real LSP, this would look up the type in our type analysis code

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// in a real LSP, this would look up the type in our type analysis code

	// HACK: Inefficient, lexes over whole document each request
	document := s.Documents[uri]
	// pos := position.Character
	s.Lexer = lexer.New(document)
	tokens := []token.Token{}
	var tok token.Token
	line := position.Line
	character := position.Character

	for tok.Type != token.EOF {
		tok = s.Lexer.NextToken()
		if tok.Type != token.EOF {
			tokens = append(tokens, tok)
		}
	}

	for _, token := range tokens {
		if line == token.Line && token.LineStart <= character && token.LineEnd >= character {
			tok = token
		}
	}

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			// Contents: fmt.Sprintf("Line=%d, Char=%d\n", position.Line, position.Character),
			Contents: fmt.Sprintf("Token:\n- Type: %q\n- Literal: %s\n- Token Line: %d\n  - Start: %d\n  - End: %d\nLSP:\n- Request Line: %d\n- Request Index: %d", tok.Type, tok.Literal, tok.Line, tok.LineStart, tok.LineEnd, line, character),
		},
	}
}
