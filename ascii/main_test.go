package main

import (
	"strings"
	"testing"
)

func TestSplitJoin(t *testing.T) {
	// Test whether splitting then rejoining a string produces the same string
	s := ""
	sections := strings.Split(s, "\\b")
	output := strings.Join(sections, "\\b")

	if output != s {
		t.Errorf("SplitJoin() output = %v, want %v", output, s)
	}
}
