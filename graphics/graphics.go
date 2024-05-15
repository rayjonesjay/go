package graphics

import (
	"ascii/args"
	"fmt"
	"os"
	"strings"
)

// GMap represents a (character -> graphics) map
type GMap = map[rune][]string

// StyleMap represents a (banner style -> graphics map) map
type StyleMap = map[string]GMap

// Cache the (character -> graphics) in a map structure
var styleCache = make(StyleMap)

// Draw takes a list of ASCII text to be displayed with their respective banner style, maps all the text to their
// respective banner graphics style, and returns a string that draws the graphics
func Draw(all []args.DrawInfo) string {
	// Check if the given slice of text to be drawn are all empty strings,
	//in which case there really is nothing to be drawn
	allEmpty := true
	for _, d := range all {
		if d.Text != "" {
			// Found one non-empty text to be displayed
			allEmpty = false
			// beak early, continuing the loop won't make any difference
			break
		}
	}

	if allEmpty {
		// Nothing to draw, thus, nothing to print
		return ""
	}

	// We have a series of graphics to draw
	var b strings.Builder
	var caret []string
	// Draw each text in order with its respective style
	for _, info := range all {
		if info.Text == "" {
			// empty text, nothing to draw
			continue
		} else if AllNewlines(info.Text) {
			// The current text to print is a special case of all (\n) newline characters
			// We don't need to print these on the caret, just print the newlines directly
			b.WriteString(strings.Repeat("\n", len(info.Text)-1))
			continue
		}
		// The current text may be on different lines, if so, we may need to advance the caret to a new line
		lines := strings.Split(info.Text, "\n")
		for i, l := range lines {
			if i == 0 {
				// first line is to be drawn, contrary to the `else` statement below,
				//we must not yet advance the caret to a newline, we may have more text on this same line
				caret = Drawln(caret, l, GetMap(info.Style))
			} else {
				// Write the previous line, we are yet to write another line
				output := SPrintCaret(caret)
				b.WriteString(output)
				// advance the caret to a newline, we are done with text on the current line,
				//as we prepare to write the next line
				b.WriteRune('\n')
				// Prepare to write a new line, but we must not yet advance the caret to a newline,
				//we may have more text on this same line
				caret = nil
				caret = Drawln(caret, l, GetMap(info.Style))
			}
		}
	}
	b.WriteString(SPrintCaret(caret))
	b.WriteRune('\n')
	return b.String()
}

// Drawln is a helper function used by the [Draw] function to draw some line of text from the current caret position.
// This therefore, assumes that s is strictly a line of text, and, thus, does not contain any newline characters
// This also expects a map of the ASCII characters to their respective art graphics
func Drawln(caret []string, s string, m map[rune][]string) []string {
	// A caret should ideally be 8 lines, we model the 8 lines with a slice of 8 strings
	if caret == nil || len(caret) < 8 {
		buffer := make([]string, 8)
		copy(buffer, caret)
		caret = buffer
	}

	if s == "" {
		return caret
	}

	// Map each ASCII character to its graphics, and append to the current caret position
	for _, char := range s {
		// Get the graphics of the current character
		g, ok := m[char]
		if !ok {
			// The current character does not exist in the (character -> graphics) map, its most likely
			//a non-ascii character or a non-printable ASCII character
			if char < 32 || char == 127 {
				//fmt.Printf("Encountered Non-printable ASCII character: \"%c\"\n", char)
				// Ignore special ASCII characters including the delete character
				continue
			} else {
				fmt.Printf("Invalid ASCII character: \"%c\"\n", char)
				os.Exit(1)
			}
		}

		if len(caret) != len(g) {
			fmt.Printf("Invalid graphics read for letter: \"%c\"\n", char)
			os.Exit(1)
		}

		// append the current character's graphics to its respective line in the caret
		for i, line := range g {
			caret[i] = caret[i] + line
		}
	}

	return caret
}

// SPrintCaret given a caret, draws the graphics for the caret to a string and returns the string
func SPrintCaret(caret []string) string {
	if caret == nil {
		// Caret null, nothing to print
		return ""
	} else if CaretEmpty(caret) {
		// Caret empty, we'll print an empty string
		return ""
	}

	var b strings.Builder

	// Print each caret line
	for i, line := range caret {
		b.WriteString(line)
		if i != len(caret)-1 {
			// Don't move the caret to a newline,
			//we might have some characters to print at the current caret line
			b.WriteRune('\n')
		}
	}

	return b.String()
}

// CaretEmpty returns true if the caret is empty, i.e., entirely composed of empty strings
func CaretEmpty(caret []string) bool {
	for _, line := range caret {
		if line != "" {
			return false
		}
	}
	return true
}

// AllNewlines returns true if all the string s is composed entirely of newline characters
func AllNewlines(s string) bool {

	//check if all characters in a given string are newLine characters
	for _, char := range s {
		if char != '\n' {
			return false
		}
	}
	// s is composed entirely of newline characters
	return true
}

// GetMap when given a given banner style, if not cached already, creates a map from the respective banner file,
// with the defined characters matched to their graphics for drawing. Returns the created map, or the cached map
func GetMap(style string) map[rune][]string {
	// attempt to retrieve the (character -> graphics) map for the given style from the cache
	m, ok := styleCache[style]
	if !ok {
		// This style graphics map isn't cached, create it
		// The banner files are in the banners subdirectory
		m = ReadBanner("banners/" + style + ".txt")
		styleCache[style] = m
	}
	return m
}
