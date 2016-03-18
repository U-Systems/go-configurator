package parser

import "testing"

func TestIsWhitespace(t *testing.T) {
	if !isWhitespace(' ') {
		t.Error("it is whitespace")
	}
	if isWhitespace('a') {
		t.Error("it is not whitespace")
	}
}

func TestLetter(t *testing.T) {
	var letters []rune
	for ch := 'a'; ch <= 'z'; ch++ {
		letters = append(letters, ch)
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		letters = append(letters, ch)
	}
	for _, ch := range letters {
		if !isLetter(ch) {
			t.Errorf("%c is letter", ch)
		}
	}
}

func TestDigit(t *testing.T) {
	var digits []rune
	for ch := '0'; ch < '9'; ch++ {
		digits = append(digits, ch)
	}
	for _, ch := range digits {
		if !isDigit(ch) {
			t.Errorf("%c is digit")
		}
	}
}

func TestIsStringStart(t *testing.T) {
	if !isStringStart('\'') {
		t.Error("is string started")
	}
}

func TestStringEnd(t *testing.T) {
	if !isStringEnd('\'') {
		t.Error("is string ended")
	}
}

func TestIsSliceStart(t *testing.T) {
	if !isSliceStart('[') {
		t.Error("is slice started")
	}
}

func TestIsSliceEnd(t *testing.T) {
	if !isSliceEnd(']') {
		t.Error("is slice ends")
	}
}

func TestIsStructStart(t *testing.T) {
	if !isStructStart('{') {
		t.Error("is struct started")
	}
}

func TestIsStructEnd(t *testing.T) {
	if !isStructEnd('}') {
		t.Error("struct ends")
	}
}
