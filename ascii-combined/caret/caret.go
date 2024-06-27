package caret

import "strings"

type Caret = []string

// SizedCaret adds a variable, Size, that keeps track of the width of the tracked Caret
type SizedCaret struct {
	Caret Caret
	Size  int
}

// Append concatenates the Caret `b` onto `a`, that is, `a` + `b`.
func Append(a, b Caret) Caret {
	if NilOrEmpty(a) {
		return Copy(b)
	} else if NilOrEmpty(b) {
		return Copy(a)
	}

	if len(a) != len(b) {
		panic("len(a) != len(b)")
	}

	out := Copy(a)
	for i := range b {
		out[i] += b[i]
	}

	return out
}

// NilOrEmpty returns true if the caret is empty, i.e., entirely composed of empty strings, or nil
func NilOrEmpty(caret []string) bool {
	if caret == nil {
		return true
	}
	for _, line := range caret {
		if line != "" {
			return false
		}
	}
	return true
}

// Copy returns a new deep copy of the given Caret
func Copy(c Caret) Caret {
	if c == nil {
		return nil
	}

	return append([]string(nil), c...)
}

// NewCaret creates a new empty caret
func NewCaret() Caret {
	return make(Caret, 8)
}

// NSpaceCaret create a new caret, whereby each line in the caret is composed of exactly n spaces
func NSpaceCaret(n int) Caret {
	spCaret := NewCaret()
	for i := range spCaret {
		spCaret[i] = strings.Repeat(" ", n)
	}
	return spCaret
}

// LargestLength returns the length of the longest line in the given caret.
func LargestLength(c Caret) (length int) {
	for _, line := range c {
		if len(line) > length {
			length = len(line)
		}
	}
	return
}
