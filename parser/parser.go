package parser

import "strings"

type Parser struct {
	scanner *Scanner
	buffer struct{
		token Token
		literal string
		n int
	       }
}

// Create new parser
func NewParser(s string) *Parser {
	return &Parser{
		scanner: NewScanner(strings.NewReader(s)),
	}
}

func (parser *Parser) scan() (token Token, literal string) {
	if parser.buffer.n != 0 {
		parser.buffer.n = 0
		return parser.buffer.token, parser.buffer.literal
	}

	token, literal = parser.scanner.Scan()
	parser.buffer.token, parser.buffer.literal = token, literal
	return
}

func (parser *Parser) unScan() {
	parser.buffer.n = 1
}