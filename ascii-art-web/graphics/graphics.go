package graphics

import (
	"ascii/data"
	"ascii/xerrors"
	"fmt"
	"strings"
)

// GMap represents a (character -> graphics) map
type GMap = map[rune][]string

// Draw takes a list of ASCII text to be displayed with their respective banner style, maps all the text to their
// respective banner graphics style, and returns a string that draws the graphics
func Draw(draw data.DrawInfo) (string, xerrors.AsciiError) {
	var b strings.Builder
	var caret []string

	// Finalizes the [Draw] function by formatting the caret output
	finalize := func() (string, xerrors.AsciiError) {
		b.WriteString(SPrintCaret(caret))
		b.WriteRune('\n')
		return b.String(), nil
	}

	if draw.Text == "" {
		// empty text, no graphics to draw
		return "", nil
	} else if AllNewlines(draw.Text) {
		b.WriteString(strings.Repeat("\n", len(draw.Text)-1))
		return finalize()
	}
	gMap, ok := GetMap(draw.Style)
	if !ok {
		return "", xerrors.InvalidBanner(fmt.Sprintf("invalid banner style: %v", draw.Style))
	}

	// The current text may be on different lines, if so, we may need to advance the caret to a new line
	lines := strings.Split(draw.Text, "\n")

	var err xerrors.AsciiError
	for i, l := range lines {
		if i == 0 {
			caret, err = Drawln(caret, l, gMap)
			if err != nil {
				return "", err
			}
		} else {
			// Write the previous line, we are yet to write another line
			output := SPrintCaret(caret)
			b.WriteString(output)
			// Prepare to write the next line
			b.WriteRune('\n')
			caret = nil
			caret, err = Drawln(caret, l, gMap)
			if err != nil {
				return "", err
			}
		}
	}

	return finalize()
}

// Drawln is a helper function used by the [Draw] function to draw some line of text from the current caret position.
// This therefore, assumes that s is strictly a line of text, and, thus, does not contain any newline characters
// This also expects a map of the ASCII characters to their respective art graphics
func Drawln(caret []string, s string, m map[rune][]string) ([]string, xerrors.AsciiError) {
	// A caret should ideally be 8 lines, we model the 8 lines with a slice of 8 strings
	if caret == nil || len(caret) < 8 {
		buffer := make([]string, 8)
		copy(buffer, caret)
		caret = buffer
	}

	if s == "" {
		return caret, nil
	}

	// Map each ASCII character to its graphics, and append to the current caret position
	for _, char := range s {
		g, ok := m[char]
		if !ok {
			if char < 32 || char == 127 {
				// Ignore special ASCII characters including the delete character
				continue
			} else {
				// The current character does not exist in the (character -> graphics) map
				return nil, xerrors.InvalidAscii(fmt.Sprintf("Invalid ASCII character: \"%c\"\n", char))
			}
		}

		if len(caret) != len(g) {
			return nil, xerrors.InvalidGraphics(fmt.Sprintf("Invalid graphics read for letter: \"%c\"\n", char))
		}

		// Append the current character's graphics to its respective line in the caret
		for i, line := range g {
			caret[i] = caret[i] + line
		}
	}

	return caret, nil
}

// SPrintCaret given a caret, draws the graphics for the caret to a string and returns the string
func SPrintCaret(caret []string) string {
	if CaretEmpty(caret) {
		return ""
	}

	// Print each caret line
	var b strings.Builder
	for i, line := range caret {
		b.WriteString(line)
		if i != len(caret)-1 {
			// Don't move the caret to a newline, on the last caret line
			b.WriteRune('\n')
		}
	}
	return b.String()
}

// CaretEmpty returns true if the caret is empty, i.e., entirely composed of empty strings
func CaretEmpty(caret []string) bool {
	if caret == nil {
		return false
	}
	for _, line := range caret {
		if line != "" {
			return false
		}
	}
	return true
}

// AllNewlines returns true if the given string is composed entirely of newline characters
func AllNewlines(s string) bool {
	for _, char := range s {
		if char != '\n' {
			return false
		}
	}
	return true
}

// GetMap when given a given banner style, if not cached already, creates a map from the respective banner file,
// with the defined characters matched to their graphics for drawing. Returns the created map, or the cached map
func GetMap(style string) (gMap GMap, ok bool) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	gMap = ReadBanner(style + ".txt")
	ok = true
	return
}
