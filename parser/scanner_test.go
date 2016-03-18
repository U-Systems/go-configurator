package parser_test

import (
	"github.com/u-systems/go-configurator/parser"
	"strings"
	"testing"
)

func initScanner(s string) *parser.Scanner {
	return parser.NewScanner(strings.NewReader(s))
}

func TestTokens(t *testing.T) {
	tests := []struct{
		testName string
		input string
		expectedTokens []parser.Token
		expectedValues []string
	}{
		{
			testName: "test eof",
			input: "",
			expectedTokens: []parser.Token{parser.EOF_TOKEN},
			expectedValues: []string{""},
		},
		{
			testName: "test ws",
			input: " ",
			expectedTokens: []parser.Token{parser.WS_TOKEN, parser.EOF_TOKEN},
			expectedValues: []string{" ", ""},
		},
		{
			testName: "test string",
			input: "'value'",
			expectedTokens: []parser.Token{parser.QUOTED_VALUE_TOKEN, parser.EOF_TOKEN},
			expectedValues: []string{"value", ""},
		},
		{
			testName: "test illegal for string",
			input: "'value",
			expectedTokens: []parser.Token{parser.ILLEGAL_TOKEN, parser.EOF_TOKEN},
			expectedValues: []string{"", ""},
		},
		{
			testName: "test keywords",
			input:  "has_value is_required format if has_default",
			expectedTokens: []parser.Token{parser.HAS_VALUE_TOKEN,
				parser.WS_TOKEN, parser.IS_REQUIRED_TOKEN, parser.WS_TOKEN, parser.FORMAT_TOKEN,
				parser.WS_TOKEN, parser.IF_TOKEN, parser.WS_TOKEN, parser.HAS_DEFAULT_TOKEN,
				parser.EOF_TOKEN},
			expectedValues: []string{"has_value", " ", "is_required",
				" ", "format", " ", "if", " ", "has_default", ""},
		},
		{
			testName: "test , :, {}, []",
			input: "[{field:'value'},{field:value1}]",
			expectedTokens: []parser.Token{parser.SLICE_START_TOKEN, parser.STRUCT_START_TOKEN,
				parser.IDENT_TOKEN, parser.COLON_TOKEN, parser.QUOTED_VALUE_TOKEN,
				parser.STRUCT_END_TOKEN,
				parser.COMMA_TOKEN,
				parser.STRUCT_START_TOKEN,
				parser.IDENT_TOKEN, parser.COLON_TOKEN, parser.IDENT_TOKEN,
				parser.STRUCT_END_TOKEN,
				parser.SLICE_END_TOKEN,
				parser.EOF_TOKEN,
			},
			expectedValues: []string{
				"", "", "field", "", "value", "", "", "", "field", "", "value1", "", "", ""},
		},
	}

	for _, test := range tests {
		scanner := initScanner(test.input)
		tokens := []parser.Token{}
		values := []string{}
		for {
			token, value := scanner.Scan()
			tokens = append(tokens, token)
			values = append(values, value)
			if token == parser.EOF_TOKEN {
				break
			}
		}

		// checking tokens
		if len(test.expectedTokens) != len(tokens) {
			t.Errorf("[%s] expected tokens does not match with actual: (%q, %q)",
				test.testName,  test.expectedTokens, tokens)
		} else {
			for i, v := range test.expectedTokens {
				if tokens[i] != v {
					t.Errorf("[%s] value for tokens does not match: (%v, %v)",
						test.testName, v, tokens[i])
				}
			}
		}
		// checking values
		if len(values) != len(test.expectedValues) {
			t.Errorf("[%s] expected values does not match with actual: (%v, %v)",
				test.testName, test.expectedValues, values)
		} else {
			for i, v := range test.expectedValues {
				if values[i] != v {
					t.Errorf("[%s] expected value does not match actual: (%v, %v)",
						test.testName, v, values[i])
				}
			}
		}
	}

}


