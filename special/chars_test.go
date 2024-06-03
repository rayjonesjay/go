package special

import (
	"strings"
	"testing"
)

func TestEscapeFEscapeV(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Mixed escape positions",
			"hello\fworld\fthere\fthere",
			"hello\n     world\n          there\n               there",
		},
		{"Escape in the middle", "hello\fworld", "hello\n     world"},
		{"Mixed escape positions with trailing", "hello\fworld\fhey\f", "hello\n     world\n          hey\n"},
		{"No target escape", "hello", "hello"},
		{"Double escapes at the end", "\fhello\f\f", "\nhello\n\n"},
		{"Triple escapes in the middle", "hello\f\f\fthere", "hello\n\n\n     there"},
		{"No target escape", "\n", "\n"},
		{"Empty string", "", ""},
		{"No target escape", `hello\\nWorld`, `hello\\nWorld`},
		{"Single indent", "g\fhello", "g\n hello"},
		{"Single character trailing escape", "hello\fg", "hello\n     g"},
		{
			"Multiple Repeated Interspersed escape positions",
			"\f\f\fHello\f\f\fThere\f\f\fThis\f\f\fI\f\f\fSuppose\f\f\fIs\f\f\fToo\f\f\fComplex\f\f\f",
			"\n\n\nHello\n\n\n     There\n\n\n          This\n\n\n              I\n\n\n               Suppose\n\n\n                      Is\n\n\n                        Too\n\n\n                           Complex\n\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := EscapeF(tt.input); got != tt.expected {
					t.Errorf("got %q, want %q", got, tt.expected)
				}
				tt.input = strings.ReplaceAll(tt.input, "\f", "\v")
				if got := EscapeV(tt.input); got != tt.expected {
					t.Errorf("got %q, want %q", got, tt.expected)
				}
			},
		)
	}
}

func TestEscapeR(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"No target escape", "Hello", "Hello"},

		{"Escape at the end", "Hello\r", "Hello"},
		{"Double escapes at the end", "Hello\r\r", "Hello"},
		{"Triple escapes at the end", "Hello\r\r\r", "Hello"},

		{"Escape in the middle", "Hello\rWorld", "World"},
		{"Double escapes in the middle", "Hello\r\rWorld", "World"},
		{"Triple escapes in the middle", "Hello\r\r\rWorld", "World"},

		{"Escape at the start", "\rHello", "Hello"},
		{"Two escapes at the start", "\r\rHello", "Hello"},
		{"Triple escapes at the start", "\r\r\rHello", "Hello"},

		{"Mixed escape positions", "\rHello\rWorld\r", "World"},
		{"Double Repeated Interspersed escape positions", "\r\rHello\r\rWorld\r\r", "World"},
		{"Triple Repeated Interspersed escape positions", "\r\r\rHello\r\r\rWorld\r\r\r", "World"},

		{"Mixed escape character literal", "r\rrHellor\rrWorldr\rr", "rWorldr"},
		{"Mixed escape character literal escaped", "\\r\r\\rHello\\r\r\\rWorld\\r\r\\r", `\rWorld\r`},
		{
			"Multiple Repeated Interspersed escape positions",
			"\r\r\rHello\r\r\rThere\r\r\rThis\r\r\rI\r\r\rSuppose\r\r\rIs\r\r\rToo\r\r\rComplex\r\r\r",
			"Complex",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := EscapeR(tt.input); got != tt.expected {
					t.Errorf("got %q, want %q", got, tt.expected)
				}
			},
		)
	}
}

func TestEscapeB(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"No target escape", "Hello", "Hello"},

		{"Escape at the end", "Hello\b", "Hello"},
		{"Double escapes at the end", "Hello\b\b", "Hello"},
		{"Triple escapes at the end", "Hello\b\b\b", "Hello"},

		{"Escape in the middle", "Hello\bWorld", "HellWorld"},
		{"Double escapes in the middle", "Hello\b\bWorld", "HelWorld"},
		{"Triple escapes in the middle", "Hello\b\b\bWorld", "HeWorld"},

		{"Escape at the start", "\bHello", "Hello"},
		{"Two escapes at the start", "\b\bHello", "Hello"},
		{"Triple escapes at the start", "\b\b\bHello", "Hello"},

		{"Mixed escape positions", "\bHello\bWorld\b", "HellWorld"},
		{"Double Repeated Interspersed escape positions", "\b\bHello\b\bWorld\b\b", "HelWorld"},
		{"Triple Repeated Interspersed escape positions", "\b\b\bHello\b\b\bWorld\b\b\b", "HeWorld"},

		{"Mixed escape character literal", "b\bbHellob\bbWorldb\bb", "bHellobWorldb"},
		{"Mixed escape character literal escaped", "\\b\b\\bHello\\b\b\\bWorld\\b\b\\b", `\\bHello\\bWorld\\b`},

		{
			"Multiple Repeated Interspersed escape positions",
			"\b\b\bHello\b\b\bThere\b\b\bThis\b\b\bI\b\b\bSuppose\b\b\bIs\b\b\bToo\b\b\bComplex\b\b\b",
			"HeTSupComplex",
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := EscapeB(tt.input); got != tt.expected {
					t.Errorf("got %q, want %q", got, tt.expected)
				}
			},
		)
	}
}
