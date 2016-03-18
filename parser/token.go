package parser

type Token int
const (
	ILLEGAL_TOKEN Token = iota
	EOF_TOKEN // end of characters
	WS_TOKEN // white space

	QUOTED_VALUE_TOKEN // any 'value'
	IDENT_TOKEN // any value except {}[],:

	STRUCT_START_TOKEN // {
	STRUCT_END_TOKEN // }

	SLICE_START_TOKEN // [
	SLICE_END_TOKEN // ]

	COMMA_TOKEN // uses in slices [value, value] and struct {field:value, field, value}
	COLON_TOKEN // uses in struct {field:value}

	FORMAT_TOKEN // has_default '2016-01-01' format 'value format'
	IS_REQUIRED_TOKEN // is_require
	HAS_DEFAULT_TOKEN // has_default value
	IF_TOKEN // if <ident> [has_value value]
	HAS_VALUE_TOKEN //has_value part of if
)

func isWhitespace(ch rune) bool {
	return ch == ' '
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isStringStart(ch rune) bool {
	return ch == '\''
}

func isStringEnd(ch rune) bool {
	return isStringStart(ch)
}

func isSliceStart(ch rune) bool {
	return ch == '['
}

func isSliceEnd(ch rune) bool {
	return ch == ']'
}

func isStructStart(ch rune) bool {
	return ch == '{'
}

func isStructEnd(ch rune) bool {
	return ch == '}'
}
