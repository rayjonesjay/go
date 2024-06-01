package args

import (
	"reflect"
	"testing"
)

// Given a testing context, runs tests on the [Parse] function with the given input ensuring the tests only passes
// when the [Parse] function produces the expected output
func helperParse(t *testing.T, input []string, expect []DrawInfo) {
	output := Parse(input)
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("Expected %v, got %v\n", expect, output)
	}
}

func TestParse(t *testing.T) {
	// Default case, parse text, to be printed, as the only command line argument
	// go run . "Hello"
	input := []string{"Hello"}
	// We expect the text to be mapped to the default standard banner style
	expect := []DrawInfo{{"Hello", Standard}}
	output := Parse(input)
	// Compare the expected output array, with the actual output array
	if !reflect.DeepEqual(output, expect) {
		t.Errorf("Expected %v, got %v\n", expect, output)
	}

	// Other cases using the test helper function

	// 0. go run .
	// No text specified, hence nothing to draw
	helperParse(t, []string{}, nil)

	// 1. go run . "\n"
	helperParse(t, []string{"\n"}, []DrawInfo{{"\n", Standard}})

	// 2. go run . "Hello" standard "World" shadow
	helperParse(t,
		[]string{"Hello", "standard", "World", "shadow"},
		[]DrawInfo{{"Hello", Standard}, {"World", Shadow}},
	)

	// 3. go run . "Hello" standard "There" shadow "World" thinkertoy
	helperParse(t,
		[]string{"Hello", "standard", "There", "shadow", "World", "thinkertoy"},
		[]DrawInfo{{"Hello", Standard}, {"There", Shadow}, {"World", Thinkertoy}},
	)

	// 4. go run . "Hello" standard "There"
	// If the last text's style is omitted, we assume it is the default standard style
	helperParse(t,
		[]string{"Hello", "standard", "There"},
		[]DrawInfo{{"Hello", Standard}, {"There", Standard}},
	)

	// 5. go run . "Hello" thinkertoy "There" shadow "World"
	// If the last text's style is omitted, we assume it is the default standard style
	helperParse(t,
		[]string{"Hello", "thinkertoy", "There", "shadow", "World"},
		[]DrawInfo{{"Hello", Thinkertoy}, {"There", Shadow}, {"World", Standard}},
	)

	// //6. go run . --output=file.txt hello standard
	// helperParse(t,
	// 	[]string{"--output.txt","hello","standard"},
	// 	[]DrawInfo{{"hello", Standard}},
	// )
}
