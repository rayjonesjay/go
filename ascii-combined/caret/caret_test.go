package caret

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		a, b     Caret
		expected Caret
	}{
		// Edge Cases
		{Caret{}, Caret{}, nil},                       // Both inputs are empty slices
		{nil, nil, nil},                               // Both inputs are nil slices
		{Caret{"a"}, Caret{}, Caret{"a"}},             // First input non-empty, second input empty
		{Caret{}, Caret{"b"}, Caret{"b"}},             // First input empty, second input non-empty
		{Caret{" "}, Caret{" "}, Caret{"  "}},         // Both inputs are spaces
		{Caret{"a"}, Caret{"b"}, Caret{"ab"}},         // Simple concatenation
		{Caret{"a"}, Caret{" bc"}, Caret{"a bc"}},     // Concatenation with space in second input
		{Caret{"abc"}, Caret{"def"}, Caret{"abcdef"}}, // Typical case
		{Caret{"123"}, Caret{"456"}, Caret{"123456"}}, // Numeric strings

		// Special Characters
		{Caret{"hello"}, Caret{"\nworld"}, Caret{"hello\nworld"}},       // Newline character
		{Caret{"foo"}, Caret{"\tbar"}, Caret{"foo\tbar"}},               // Tab character
		{Caret{"alpha"}, Caret{"\u03B2beta"}, Caret{"alpha\u03B2beta"}}, // Unicode character

		// Mixed Content
		{Caret{"string"}, Caret{"123"}, Caret{"string123"}}, // Alphanumeric
		// Classic
		{Caret{"a", "b", "c"}, Caret{"d", "e", "f"}, Caret{"ad", "be", "cf"}},
		{
			Caret{"Line", "Line", "Line", "Line", "Line", "Line", "Line", "Line"},
			Caret{"1", "2", "3", "4", "5", "6", "7", "8"},
			Caret{"Line1", "Line2", "Line3", "Line4", "Line5", "Line6", "Line7", "Line8"},
		},
	}

	for _, tt := range tests {
		result := Append(tt.a, tt.b)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Append(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		c Caret
	}
	tests := []struct {
		name string
		args args
		want Caret
	}{
		{"Unicode character", args{Caret{"alpha\u03B2beta"}}, Caret{"alpha\u03B2beta"}},
		{"Tab character", args{Caret{"foo\tbar"}}, Caret{"foo\tbar"}},
		{"Nil caret", args{nil}, nil},
		{"Empty caret", args{Caret{}}, nil},
		{
			"More characters", args{Caret{"hello", "there", "this", "is", "the", "test"}},
			Caret{"hello", "there", "this", "is", "the", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := Copy(tt.args.c)
				if len(tt.args.c) > 0 {
					tt.args.c[0] = "test string"
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Copy() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestNewCaret(t *testing.T) {
	want := []string{"", "", "", "", "", "", "", ""}
	if got := NewCaret(); !reflect.DeepEqual(got, want) {
		t.Errorf("NewCaret() = %v, want %v", got, want)
	}
}

func TestNilOrEmpty(t *testing.T) {
	type args struct {
		caret []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Nil Caret", args: args{nil}, want: true},
		{name: "Empty Caret", args: args{Caret{}}, want: true},
		{name: "Caret with empty string", args: args{Caret{""}}, want: true},
		{name: "Caret with two string", args: args{Caret{"", ""}}, want: true},
		{name: "Caret with empty strings", args: args{Caret{"", "", "", "", "", "", "", ""}}, want: true},
		{
			name: "Caret with empty strings - space at end", args: args{Caret{"", "", "", "", "", "", "", " "}},
			want: false,
		},
		{
			name: "Caret with empty strings at the end", args: args{Caret{"Hi", "this", "is", "it", "", "", "", ""}},
			want: false,
		},
		{
			name: "Caret with empty strings in the middle",
			args: args{Caret{"Hi", "this", "is", "it", "", "", "at", "end"}},
			want: false,
		},
		{
			name: "Caret with empty strings at start", args: args{Caret{"", "", "", "Hello", "", "This", "is", "it"}},
			want: false,
		},
		{name: "Caret with single space", args: args{Caret{" "}}, want: false},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := NilOrEmpty(tt.args.caret); got != tt.want {
					t.Errorf("NilOrEmpty() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
