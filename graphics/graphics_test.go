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
