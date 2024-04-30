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
	var b strings.Builder
	var caret []string
	// Draw each text in order with its respective style
	for _, info := range all {
		// The current text may be on different lines, if so, we may need to advance the caret to a new line
		lines := strings.Split(info.Text, "\n")
		for i, l := range lines {
			if i == 0 {
				// first line is to be drawn, contrary to the else statement below,
				//we must not yet advance the caret to a newline, we may have more text on this same line
				caret = Drawln(caret, l, GetMap(info.Style))
			} else {
				// Write the previous line, we are yet to write another line
				output := SPrintCaret(caret)
				if output != "" {
					b.WriteString(output)
					// advance the caret to a newline, we are done with text on the current line,
					//as we prepare to write the next line
					b.WriteRune('\n')
				}
				// Prepare to write a new line, but we must not yet advance the caret to a newline,
				//we may have more text on this same line
				caret = nil
				caret = Drawln(caret, l, GetMap(info.Style))
			}
		}
	}
	b.WriteString(SPrintCaret(caret))
	return b.String()
}

// Drawln is a helper function used by the [Draw] function to draw some line of text from the current caret position.
// This therefore, assumes that s is strictly a line of text, and, thus, does not contain any newline characters
// This also expects a map of the ASCII characters to their respective art graphics
func Drawln(caret []string, s string, m map[rune][]string) []string {
	if s == "" {
		return caret
	}

	// A caret should ideally be 8 lines, we model the 8 lines with a slice of 8 items
	if caret == nil || len(caret) < 8 {
		buffer := make([]string, 8)
		for i, cl := range caret {
			buffer[i] = cl
		}
		caret = buffer
	}

	// Map each ASCII character to its graphics, and append to the current caret position
	for _, char := range s {
		// Get the graphics of the current character
		g, ok := m[char]
		if !ok {
			// The current character does not exist in the (character -> graphics) map, its most likely
			//a non-ascii character or a non-printable ASCII character
			if char < 32 {
				fmt.Printf("Encountered Non-printable ASCII character: \"%c\"\n", char)
			} else {
				fmt.Printf("Invalid ASCII character: \"%c\"\n", char)
			}
			os.Exit(1)
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
		// Caret empty, nothing to print
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

func PrintCaret(caret []string) {
	fmt.Print(SPrintCaret(caret))
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
