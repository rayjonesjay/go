package args

import "testing"

func helperNewline(t *testing.T, input, expected string) {
	output := Escape(input)
	if output != expected {
		t.Errorf("input:\n\"%s\"\n expected: \n\"%s\", got:\n\"%s\"", input, expected, output)
	}
}

func TestEscape(t *testing.T) {
	helperNewline(t, "Hello\\nWorld", "Hello\nWorld")
	helperNewline(t, `Hello\nWorld`, "Hello\nWorld")

	helperNewline(t, `Hello\\nWorld`, `Hello\nWorld`)
}
