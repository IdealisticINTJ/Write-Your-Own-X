package main

import (
	"fmt"
	"os"
	"unicode"
)

// TokenType represents the type of a token.
type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenString
	TokenNumber
	TokenBoolean
	TokenNull
	TokenObjectStart
	TokenObjectEnd
	TokenArrayStart
	TokenArrayEnd
	TokenComma
	TokenColon
	TokenBeginArray
	TokenEndArray
	TokenValueSeparator
)

// Token represents a parsed token.
type Token struct {
	Type  TokenType
	Value string
}

// JSONParser represents the JSON parser.
type JSONParser struct {
	input  string
	cursor int
}

// NewJSONParser creates a new JSON parser instance.
func NewJSONParser(input string) *JSONParser {
	return &JSONParser{input: input}
}

// NextToken parses and returns the next token in the input.
func (p *JSONParser) NextToken() Token {
	p.skipWhitespace()

	if p.cursor >= len(p.input) {
		return Token{Type: TokenEOF}
	}

	switch ch := p.input[p.cursor]; {
	case ch == '{':
		p.cursor++
		return Token{Type: TokenObjectStart, Value: "{"}
	case ch == '}':
		p.cursor++
		return Token{Type: TokenObjectEnd, Value: "}"}
	case ch == '[':
		p.cursor++
		return Token{Type: TokenBeginArray, Value: "["}
	case ch == ']':
		p.cursor++
		return Token{Type: TokenEndArray, Value: "]"}
	case ch == ',':
		p.cursor++
		return Token{Type: TokenComma, Value: ","}
	case ch == ':':
		p.cursor++
		return Token{Type: TokenColon, Value: ":"}
	case ch == 't' || ch == 'f':
		return p.parseBoolean()
	case ch == 'n':
		return p.parseNull()
	case ch == '"':
		return p.parseString()
	case unicode.IsDigit(rune(ch)) || ch == '-':
		return p.parseNumber()
	default:
		return Token{Type: TokenError, Value: fmt.Sprintf("Unexpected character: %c", ch)}
	}
}

// skipWhitespace skips whitespace characters.
func (p *JSONParser) skipWhitespace() {
	for p.cursor < len(p.input) && unicode.IsSpace(rune(p.input[p.cursor])) {
		p.cursor++
	}
}

// parseBoolean parses a boolean value.
func (p *JSONParser) parseBoolean() Token {
	start := p.cursor
	for p.cursor < len(p.input) && (unicode.IsLetter(rune(p.input[p.cursor])) || p.input[p.cursor] == '.') {
		p.cursor++
	}

	value := p.input[start:p.cursor]
	if value == "true" || value == "false" {
		return Token{Type: TokenBoolean, Value: value}
	}

	return Token{Type: TokenError, Value: fmt.Sprintf("Invalid boolean: %s", value)}
}

// parseNull parses a null value.
func (p *JSONParser) parseNull() Token {
	start := p.cursor
	for p.cursor < len(p.input) && unicode.IsLetter(rune(p.input[p.cursor])) {
		p.cursor++
	}

	value := p.input[start:p.cursor]
	if value == "null" {
		return Token{Type: TokenNull, Value: value}
	}

	return Token{Type: TokenError, Value: fmt.Sprintf("Invalid null: %s", value)}
}

// parseString parses a string value.
func (p *JSONParser) parseString() Token {
	p.cursor++ // Skip the opening quote
	start := p.cursor

	for p.cursor < len(p.input) && p.input[p.cursor] != '"' {
		p.cursor++
	}

	if p.cursor >= len(p.input) || p.input[p.cursor] != '"' {
		return Token{Type: TokenError, Value: "Unterminated string"}
	}

	value := p.input[start:p.cursor]
	p.cursor++ // Skip the closing quote
	return Token{Type: TokenString, Value: value}
}

// parseNumber parses a number value.
func (p *JSONParser) parseNumber() Token {
	start := p.cursor

	for p.cursor < len(p.input) && (unicode.IsDigit(rune(p.input[p.cursor])) || p.input[p.cursor] == '.') {
		p.cursor++
	}

	value := p.input[start:p.cursor]
	return Token{Type: TokenNumber, Value: value}
}

// parseArray parses an array value.
func (p *JSONParser) parseArray() Token {
	p.cursor++ // Skip the opening bracket
	start := p.cursor

	for p.cursor < len(p.input) && p.input[p.cursor] != ']' {
		p.cursor++
	}

	if p.cursor >= len(p.input) || p.input[p.cursor] != ']' {
		return Token{Type: TokenError, Value: "Unterminated array"}
	}

	value := p.input[start:p.cursor]
	p.cursor++ 
	return Token{Type: TokenString, Value: value}
}

func main() {
	input := `{"key": [1, 2, 3]}`
	parser := NewJSONParser(input)

	firstToken := parser.NextToken()
	lastToken := firstToken

	for {
		token := parser.NextToken()

		if token.Type == TokenEOF {
			break
		}

		if token.Type == TokenError {
			fmt.Println("Invalid JSON")
			os.Exit(1)
		}
	}

	if firstToken.Type != TokenObjectStart || lastToken.Type != TokenObjectEnd {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}

	fmt.Println("Valid JSON")
}
