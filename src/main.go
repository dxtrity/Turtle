package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Token represents a token in the input stream.
type Token struct {
	Type  string // Type of the token (e.g., "NUMBER", "PLUS", "IDENT", "ASSIGN", "EOF")
	Value string // Value of the token (e.g., "42", "+", "x", "=")
}

// Lexer scans the input string and produces tokens.
type Lexer struct {
	scanner *bufio.Scanner
	tokens  []Token
}

// NewLexer creates a new lexer with the given input file.
func NewLexer(filename string) (*Lexer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lexer := &Lexer{
		scanner: scanner,
		tokens:  make([]Token, 0),
	}
	lexer.tokenizeInput()
	return lexer, nil
}

// tokenizeInput scans the input file and tokenizes the input.
func (l *Lexer) tokenizeInput() {
	for l.scanner.Scan() {
		line := l.scanner.Text()
		tokens := l.tokenizeLine(line)
		l.tokens = append(l.tokens, tokens...)
	}
}

// tokenizeLine tokenizes a single line of input.
func (l *Lexer) tokenizeLine(line string) []Token {
	tokens := make([]Token, 0)
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		tokenText := scanner.Text()
		tokenType := l.getTokenType(tokenText)
		token := Token{Type: tokenType, Value: tokenText}
		tokens = append(tokens, token)
	}

	return tokens
}

// getTokenType determines the type of the token.
func (l *Lexer) getTokenType(tokenText string) string {
	switch tokenText {
	case "+":
		return "PLUS"
	case "-":
		return "MINUS"
	case "*":
		return "MULTIPLY"
	case "/":
		return "DIVIDE"
	case "=":
		return "ASSIGN"
	case "(":
		return "LPAREN"
	case ")":
		return "RPAREN"
	default:
		if _, err := strconv.Atoi(tokenText); err == nil {
			return "NUMBER"
		} else if unicode.IsLetter(rune(tokenText[0])) {
			return "IDENT"
		}
	}
	return "UNKNOWN"
}

// NextToken returns the next token in the input stream.
func (l *Lexer) NextToken() Token {
	if len(l.tokens) == 0 {
		return Token{Type: "EOF", Value: ""}
	}
	token := l.tokens[0]
	l.tokens = l.tokens[1:]
	return token
}

// Parser represents a recursive descent parser.
type Parser struct {
	lexer     *Lexer
	curToken  Token
	variables map[string]int // Map of variable name to variable value
}

// NewParser creates a new parser with the given lexer.
func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{
		lexer:     lexer,
		variables: make(map[string]int),
	}
	parser.consumeToken() // Initialize current token
	return parser
}

// consumeToken advances to the next token in the input stream.
func (p *Parser) consumeToken() {
	p.curToken = p.lexer.NextToken()
}

// parseStatement parses a statement (either variable assignment or expression).
func (p *Parser) parseStatement() {
	switch p.curToken.Type {
	case "IDENT":
		// Variable assignment
		varName := p.curToken.Value
		p.consumeToken() // Consume variable name
		if p.curToken.Type == "ASSIGN" {
			p.consumeToken() // Consume ASSIGN token
			value := p.parseExpression()
			p.variables[varName] = value
		} else {
			// If no ASSIGN token, treat it as an expression
			// Print the result
			fmt.Println(p.evaluateExpression(varName))
		}
	default:
		// Expression statement
		fmt.Println(p.parseExpression())
	}
}

// parseExpression parses an expression.
func (p *Parser) parseExpression() int {
	// Parse the first term
	left := p.parseTerm()

	for p.curToken.Type == "PLUS" || p.curToken.Type == "MINUS" || p.curToken.Type == "MULTIPLY" || p.curToken.Type == "DIVIDE" {
		// Store the operator
		operator := p.curToken.Type
		p.consumeToken()

		// Parse the next term
		right := p.parseTerm()

		// Perform the operation based on the operator
		switch operator {
		case "PLUS":
			left += right
		case "MINUS":
			left -= right
		case "MULTIPLY":
			left *= right
		case "DIVIDE":
			if right == 0 {
				panic("Division by zero")
			}
			left /= right
		}

		// Check for a right parenthesis to handle parentheses
		if p.curToken.Type == "RPAREN" {
			return left
		}
	}

	return left
}

// parseTerm parses a term (number, variable reference, or parentheses expression).
func (p *Parser) parseTerm() int {
	switch p.curToken.Type {
	case "NUMBER":
		// Parse the number
		number, err := strconv.Atoi(p.curToken.Value)
		if err != nil {
			panic("Failed to parse number")
		}

		// Consume the NUMBER token
		p.consumeToken()

		return number
	case "IDENT":
		// Variable reference
		varName := p.curToken.Value
		p.consumeToken() // Consume variable name

		// Check if variable exists
		value, ok := p.variables[varName]
		if !ok {
			panic(fmt.Sprintf("Undefined variable: %s", varName))
		}

		return value
	case "LPAREN":
		// Consume the left parenthesis
		p.consumeToken()

		// Parse the expression inside the parentheses
		result := p.parseExpression()

		// Ensure a matching right parenthesis
		if p.curToken.Type != "RPAREN" {
			panic("Expected RPAREN")
		}

		// Consume the right parenthesis
		p.consumeToken()

		return result
	default:
		panic("Unexpected token")
	}
}

// evaluateExpression evaluates an expression with variable references.
func (p *Parser) evaluateExpression(varName string) int {
	value, ok := p.variables[varName]
	if !ok {
		panic(fmt.Sprintf("Undefined variable: %s", varName))
	}
	return value
}

func main() {
	// Check if a filename is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go inputfile")
		return
	}

	filename := os.Args[1]

	lexer, err := NewLexer(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	parser := NewParser(lexer)

	// Parse statements
	for {
		parser.parseStatement()
		if parser.curToken.Type == "EOF" {
			break
		}
	}
}
