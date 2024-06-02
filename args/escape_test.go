package args

import (
	"testing"
)

// Ease running multiple tests in TestEscape
type testEscape struct {
	input    string
	expected string
}

func TestEscape(t *testing.T) {
	tests := []testEscape{
		{``, ""},
		{`\n`, "\n"},
		{`\n\n`, "\n\n"},
		{`\\n\\n`, "\\n\\n"},

		{"Hello\\nWorld", "Hello\nWorld"},
		{`Hello\nWorld`, "Hello\nWorld"},
		{`Hello\\nWorld`, `Hello\nWorld`},

		{`Hello\aWorld`, "Hello\aWorld"},
		{`\aHello\aWorld\a`, "\aHello\aWorld\a"},

		{`Hello\bWorld`, "Hello\bWorld"},
		{`\bHello\bWorld\b`, "\bHello\bWorld\b"},

		{`\tHello\tWorld\t`, "\tHello\tWorld\t"},
		{`\\tHello\\tWorld\t`, "\\tHello\\tWorld\t"},

		{`\vHello\vWorld\v`, "\vHello\vWorld\v"},

		{`\\f\\f\\f`, "\\f\\f\\f"},

		{`\r`, "\r"},
		{`End\r`, "End\r"},
		{`\rStart`, "\rStart"},
		{`\r\rStartRepeat`, "\r\rStartRepeat"},
	}

	for _, mt := range tests {
		output := Escape(mt.input)
		if output != mt.expected {
			t.Errorf("\ninput:\n\"%s\"\nexpected:\n%#v\ngot:\n%#v", mt.input, mt.expected, output)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name   string
		input  []rune
		expect string
	}{
		{
			name:   "Empty Slice",
			input:  []rune{},
			expect: "",
		},
		{
			name:   "Single Null Character",
			input:  []rune{0},
			expect: "",
		},
		{
			name:   "Multiple Null Characters",
			input:  []rune{0, 0, 0},
			expect: "",
		},
		{
			name:   "Leading Null Characters",
			input:  []rune{0, 0, 'a', 'b', 'c'},
			expect: "abc",
		},
		{
			name:   "Trailing Null Characters",
			input:  []rune{'a', 'b', 'c', 0, 0},
			expect: "abc",
		},
		{
			name:   "Interspersed Null Characters",
			input:  []rune{'a', 0, 'b', 0, 'c', 0},
			expect: "abc",
		},
		{
			name:   "No Null Characters",
			input:  []rune{'a', 'b', 'c'},
			expect: "abc",
		},
		{
			name:   "Special Characters",
			input:  []rune{'😀', '🚀', '🌟'},
			expect: "😀🚀🌟",
		},
		{
			name:   "Mixed Characters",
			input:  []rune{'a', '1', 0, '😀', '🚀', 0, 'b', '2'},
			expect: "a1😀🚀b2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToString(tt.input)
			if result != tt.expect {
				t.Errorf("expected %q, got %q", tt.expect, result)
			}
		})
	}
}

func TestHexStringToDecimal(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
		ok     bool
	}{
		{
			name:   "Empty String",
			input:  "",
			expect: 0,
			ok:     false,
		},
		{
			name:   "Invalid Characters",
			input:  "GHIJK",
			expect: 0,
			ok:     false,
		},
		{
			name:   "Valid Hexadecimal String",
			input:  "1A3F",
			expect: 6719,
			ok:     false,
		},
		{
			name:   "Leading Zeros",
			input:  "0001A3F",
			expect: 6719,
			ok:     false,
		},
		{
			name:   "Maximum Value for Integer",
			input:  "7FFFFFFF",
			expect: 2147483647,
			ok:     false,
		},
		{
			name:   "Negative Hexadecimal String",
			input:  "-1A3F",
			expect: -6719,
			ok:     false,
		},
		{
			name:   "Lowercase and Uppercase Letters",
			input:  "aBcDeF",
			expect: 11259375,
			ok:     false,
		},
		{
			name:   "Zero Value",
			input:  "0",
			expect: 0,
			ok:     true,
		},
		{
			name:   "Whitespace Characters",
			input:  " 1A 3 F ",
			expect: 0,
			ok:     false,
		},
		{
			name:   "Hexadecimal Prefix",
			input:  "0x1A3F",
			expect: 0,
			ok:     false,
		},
		{
			name:   "Uppercase Valid Hexadecimal String",
			input:  "1A",
			expect: 26,
			ok:     true,
		},
		{
			name:   "Lowercase Valid Hexadecimal String",
			input:  "1a",
			expect: 26,
			ok:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := HexStringToDecimal(tt.input)
			if result != tt.expect || ok != tt.ok {
				t.Errorf("expected %d, %v; got %d, %v", tt.expect, tt.ok, result, ok)
			}
		})
	}
}

func TestOctalStringToDecimal(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
		ok     bool
	}{
		{
			name:   "Empty String",
			input:  "",
			expect: 0,
			ok:     false,
		},
		{
			name:   "Invalid Characters",
			input:  "89",
			expect: 0,
			ok:     false,
		},
		{
			name:   "Valid Octal String",
			input:  "10",
			expect: 8,
			ok:     true,
		},
		{
			name:   "Leading Zeros",
			input:  "0010",
			expect: 8,
			ok:     true,
		},
		{
			name:   "Edge of Range - Lower",
			input:  "0",
			expect: 0,
			ok:     true,
		},
		{
			name:   "Edge of Range - Upper",
			input:  "177",
			expect: 127,
			ok:     true,
		},
		{
			name:   "Out of Range Value",
			input:  "200",
			expect: 128,
			ok:     false,
		},
		{
			name:   "Mixed Valid and Invalid Characters",
			input:  "1a2",
			expect: 0,
			ok:     false,
		},
		{
			name:   "Zero Value",
			input:  "0",
			expect: 0,
			ok:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := OctalStringToDecimal(tt.input)
			if result != tt.expect || ok != tt.ok {
				t.Errorf("expected %d, %v; got %d, %v", tt.expect, tt.ok, result, ok)
			}
		})
	}
}
