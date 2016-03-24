package parser

import (
	"bufio"
	"io"
	"bytes"
	"strings"
)

const eof = rune(0)

type Scanner struct {
	reader *bufio.Reader
}

func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{
		reader: bufio.NewReader(reader),
	}
}

// Scan for token
func (scanner *Scanner) Scan() (Token, string) {
	ch := scanner.read()

	if ch == eof {
		return EOF_TOKEN, ""
	}

	switch  {
	case isWhitespace(ch):
		scanner.unRead()
		return scanner.scanWhiteSpace()
	case isStringStart(ch):
		scanner.unRead()
		return scanner.scanQuoted()
	case isSliceStart(ch):
		return SLICE_START_TOKEN, ""
	case isSliceEnd(ch):
		return SLICE_END_TOKEN, ""
	case isStructStart(ch):
		return STRUCT_START_TOKEN, ""
	case isStructEnd(ch):
		return STRUCT_END_TOKEN, ""
	case ch == ',':
		return COMMA_TOKEN, ""
	case ch == ':':
		return COLON_TOKEN, ""
	}

	scanner.unRead()
	return scanner.scanIdent()
}

func (scanner *Scanner) read() rune {
	ch, _, err := scanner.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (scanner *Scanner) unRead() {
	_ = scanner.reader.UnreadRune()
}

// scan for whitespace
func (scanner *Scanner) scanWhiteSpace() (Token, string) {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.read())

	for {
		if ch := scanner.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			scanner.unRead()
			break
		} else {
			buffer.WriteRune(ch)
		}
	}

	return WS_TOKEN, buffer.String()
}

// Scan ident
// e.g
// config:"has_value value", config:"has_default 1.5"
// config:"has_value 2006-01-01"
func (scanner *Scanner) scanIdent() (Token, string) {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.read())

	for {
		if ch := scanner.read(); ch == eof {
			break
		} else if !isLetter(ch) && ch != '_' && !isDigit(ch) && ch != '.' && ch != '-' {
			scanner.unRead()
			break
		} else {
			buffer.WriteRune(ch)
		}
	}

	switch value := strings.ToLower(buffer.String()); value {
	case "is_required":
		return IS_REQUIRED_TOKEN, value
	case "has_default":
		return HAS_DEFAULT_TOKEN, value
	case "if":
		return IF_TOKEN, value
	case "has_value":
		return HAS_VALUE_TOKEN, value
	case "format":
		return FORMAT_TOKEN, value
	}

	return IDENT_TOKEN, buffer.String()
}

// Scan for values in ''
// e.g config:"has_default 'value'"
func (scanner *Scanner) scanQuoted() (Token, string) {
	var buffer bytes.Buffer
	_ = scanner.read() // skip first
	for {
		if ch := scanner.read(); ch == eof {
			return ILLEGAL_TOKEN, ""
		} else if isStringEnd(ch) {
			break
		} else {
			buffer.WriteRune(ch)
		}
	}

	return QUOTED_VALUE_TOKEN, buffer.String()
}

