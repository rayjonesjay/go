package graphics

import "testing"

// The caret is basically a graphics line, which in essence is actually 8 lines as below
// -------------------------------------
// aaa
// bbb
// ccc
// ddd
// eee
// fff
// ggg
// hhh
// --------------------------------------
func TestSPrintCaret(t *testing.T) {
	caret := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh"}
	expected := "aaa\nbbb\nccc\nddd\neee\nfff\nggg\nhhh"
	output := SPrintCaret(caret)

	if output != expected {
		t.Errorf("SPrintCaret(%v) returned:\n%v \nexpected:\n%v\n", caret, output, expected)
	}
}

func TestCaretEmpty(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		expect bool
	}{
		{
			name:   "Empty Slice",
			input:  []string{},
			expect: true,
		},
		{
			name:   "Single Empty String",
			input:  []string{""},
			expect: true,
		},
		{
			name:   "Single Non-Empty String",
			input:  []string{"non-empty"},
			expect: false,
		},
		{
			name:   "Multiple Empty Strings",
			input:  []string{"", "", ""},
			expect: true,
		},
		{
			name:   "Multiple Non-Empty Strings",
			input:  []string{"a", "b", "c"},
			expect: false,
		},
		{
			name:   "Mixed Empty and Non-Empty Strings",
			input:  []string{"", "non-empty", ""},
			expect: false,
		},
		{
			name:   "Whitespace Strings",
			input:  []string{" ", "\t", "\n"},
			expect: false,
		},
		{
			name:   "Whitespace and Empty Strings",
			input:  []string{"", " ", "\t", "\n", ""},
			expect: false,
		},
		{
			name:   "Long Strings",
			input:  []string{string(make([]byte, 1000)), "non-empty"},
			expect: false,
		},
		{
			name:   "Special Characters",
			input:  []string{"\n", "\t", "\r"},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CaretEmpty(tt.input)
			if result != tt.expect {
				t.Errorf("expected %v; got %v", tt.expect, result)
			}
		})
	}
}
