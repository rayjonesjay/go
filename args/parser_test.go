package args

import (
	"ascii/data"
	"reflect"
	"testing"
)

// Ease running multiple tests in TestParse
type testParse struct {
	input  []string
	expect ParserOut
}

func TestParse(t *testing.T) {
	tests := []testParse{
		// go run .
		{[]string{}, ParserOut{}},
		// go run . "Hello"
		{[]string{"Hello"},
			ParserOut{Draws: &data.DrawInfo{Text: "Hello", Style: Standard}}},
		// go run . "\n"
		{[]string{"\n"},
			ParserOut{Draws: &data.DrawInfo{Text: "\n", Style: Standard}}},
		// go run . "Hello" standard
		{[]string{"Hello", "standard"},
			ParserOut{Draws: &data.DrawInfo{Text: "Hello", Style: Standard}}},
		// go run . "Hey" shadow
		{[]string{"Hey", "shadow"},
			ParserOut{Draws: &data.DrawInfo{Text: "Hey", Style: Shadow}}},
		// go run . "Hey" thinkertoy
		{[]string{"Hey", "thinkertoy"},
			ParserOut{Draws: &data.DrawInfo{Text: "Hey", Style: Thinkertoy}}},
		//go run . --output=file.txt Hey thinkertoy
		{[]string{"--output=file.txt", "Hey", "thinkertoy"},
			ParserOut{Draws: &data.DrawInfo{Text: "Hey", Style: Thinkertoy}, OutputFile: "file.txt"}},
	}

	for i, test := range tests {
		output := Parse(test.input)
		if !reflect.DeepEqual(output, test.expect) {
			t.Errorf("Test: %d Expected %v, got %v\n", i+1, test.expect, output)
		}
	}
}
