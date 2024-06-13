package graphics

import (
	"ascii/caret"
	color "ascii/colors"
	"ascii/data"
	"ascii/fmtx"
	"fmt"
	"os"
	"strings"
)

// GMap represents a (character -> graphics) map
type GMap = map[rune][]string

// StyleMap represents a (banner style -> graphics map) map
type StyleMap = map[string]GMap

// ColorRange defines the range of indices (Start to End exclusive) of letters in a string,
// that need to be colored by the specified Color.
type ColorRange struct {
	Color      color.Color
	Start, End int
}

// Cache the (character -> graphics) in a map structure
var styleCache = make(StyleMap)

// Draw takes a list of ASCII text to be displayed with their respective banner style, maps all the text to their
// respective banner graphics style, and returns a string that draws the graphics
func Draw(draw data.DrawInfo) string {
	var b strings.Builder
	var drawCaret []string

	// Finalizes the [Draw] function by formatting the drawCaret output
	finalize := func() string {
		b.WriteString(SPrintCaret(drawCaret))
		b.WriteRune('\n')
		return b.String()
	}

	if draw.Text == "" {
		// empty text, no graphics to draw
		return ""
	} else if AllNewlines(draw.Text) {
		b.WriteString(strings.Repeat("\n", len(draw.Text)-1))
		return finalize()
	}

	// The current text may be on different lines, if so, we may need to advance the drawCaret to a new line
	lines := strings.Split(draw.Text, "\n")
	for i, l := range lines {
		if i == 0 {
			drawCaret = Drawln(drawCaret, l, GetMap(draw.Style))
		} else {
			// Write the previous line, we are yet to write another line
			output := SPrintCaret(drawCaret)
			b.WriteString(output)
			// Prepare to write the next line
			b.WriteRune('\n')
			drawCaret = nil
			drawCaret = Drawln(drawCaret, l, GetMap(draw.Style))
		}
	}

	return finalize()
}

// Drawln is a helper function used by the [Draw] function to draw some line of text from the current caret position.
// This therefore, assumes that s is strictly a line of text, and, thus, does not contain any newline characters
// This also expects a map of the ASCII characters to their respective art graphics
func Drawln(lineCaret []string, s string, m map[rune][]string) []string {
	// A lineCaret should ideally be 8 lines, we model the 8 lines with a slice of 8 strings
	if lineCaret == nil || len(lineCaret) < 8 {
		buffer := make([]string, 8)
		copy(buffer, lineCaret)
		lineCaret = buffer
	}

	if s == "" {
		return lineCaret
	}

	return caret.Append(lineCaret, drawln(lineCaret, s, m))
}

const (
	ALIGN_LEFT    = "left"
	ALIGN_RIGHT   = "right"
	ALIGN_CENTER  = "center"
	ALIGN_JUSTIFY = "justify"
)

// draw the text in a single line, taking into account the width of the terminal and the expected text alignment
func drawln(in caret.Caret, s string, m map[rune][]string) (out caret.Caret) {
	// Split this line into words
	words := strings.Split(s, " ")

	lineLength := 0
	termWidth := 120
	align := getAlignment()

	var wordCarets []caret.Caret
	for _, w := range words {
		wordCaret := drawWord(w, m)
		lineLength += wordCaret.Size
		wordCarets = append(wordCarets, wordCaret.Caret)
	}

	spaceWidth := (len(words) - 1) * spaceSize(m)
	if spaceWidth < 0 {
		spaceWidth = 0
	}

	addLength := termWidth - lineLength - spaceWidth
	if addLength < 0 {
		fmtx.FatalErrorf("Can't fit graphics to terminal; increase terminal size\n")
	}

	out = caret.Append(in, out)
	justifySpacers := 0
	if len(words)-1 > 0 {
		justifySpacers = addLength / (len(words) - 1)
	}

	for i, wc := range wordCarets {
		if i != 0 {
			spaceGraphics := m[' ']
			if align == ALIGN_JUSTIFY && justifySpacers > 0 {
				out = caret.Append(out, caret.NSpaceCaret(justifySpacers))
			}
			out = caret.Append(out, spaceGraphics)
		}
		out = caret.Append(out, wc)
	}

	if align == ALIGN_RIGHT {
		out = caret.Append(caret.NSpaceCaret(addLength), out)
	} else if align == ALIGN_CENTER {
		leftWidth := addLength / 2
		out = caret.Append(caret.NSpaceCaret(leftWidth), out)
	}

	return
}

func getAlignment() string {
	return "justify"
	// return "right"
}

// returns the width of drawing the space character based on the given banner file's (letter -> graphics) map
func spaceSize(m map[rune][]string) int {
	graphics, ok := m[' ']
	if ok && len(graphics) > 0 {
		return len(graphics[0])
	} else {
		fmtx.FatalErrorf("Couldn't find the ASCII graphic for space character\n")
		return 0
	}
}

func drawWord(s string, m map[rune][]string) (sizedCaret caret.SizedCaret) {
	colorRangeList := ColorRangeList(
		s, []data.ColorInfo{
			{
				Color: color.Color{
					R: 255,
					G: 0,
					B: 0,
				},
				Substr: "hi",
			},
			{
				Color: color.Color{
					R: 255,
					G: 255,
					B: 0,
				},
				Substr: "to",
			},
		},
	)
	c := caret.NewCaret()
	// Map each ASCII character to its graphics, and append to the current caret position
	for i, char := range s {
		charGraphics, ok := m[char]
		if !ok {
			if char < 32 || char == 127 {
				// Ignore special ASCII characters including the delete character
				continue
			} else {
				// The current character does not exist in the (character -> graphics) map
				fmt.Printf("Invalid ASCII character: \"%c\"\n", char)
				os.Exit(1)
			}
		}

		// Increment the caret's size by the size of the current characters graphics
		sizedCaret.Size += caret.LargestLength(charGraphics)
		colorCode, resetCode := letterColor(colorRangeList, i)
		// Append the current character's graphics to its respective line in the caret
		for j, line := range charGraphics {
			c[j] = c[j] + colorCode + line + resetCode
		}
	}

	sizedCaret.Caret = c
	return
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
func GetMap(style string) map[rune][]string {
	m, ok := styleCache[style]
	if !ok {
		// This style graphics map isn't cached, create it
		m = ReadBanner(style + ".txt")
		styleCache[style] = m
	}
	return m
}

// ColorRangeList builds a list of ColorRange objects, that define the character indices in the string, text,
// that ought to be colored by a given color. Note that the color flags define substrings in text,
// that ought to be colored by any specified color
func ColorRangeList(text string, colorFlags []data.ColorInfo) (out []ColorRange) {
	for _, cf := range colorFlags {
		iterativeText := text
		for {
			if cf.Substr == "" {
				out = append(out, ColorRange{cf.Color, 0, len(text)})
				break
			}
			startIndex := strings.Index(iterativeText, cf.Substr)
			if startIndex == -1 {
				break
			}
			endIndex := startIndex + len(cf.Substr)
			out = append(out, ColorRange{cf.Color, startIndex, endIndex})

			iterativeText = strings.Replace(iterativeText, cf.Substr, "", 1)
		}
	}
	return
}

// checks if the given letter Index exists in any of the color range. If it does,
// then it returns the ANSI escape code for the given color it is a range of,
// and the ANSI color code to reset the color back to the terminal default; otherwise an empty color and reset code is
// returned
func letterColor(colorRange []ColorRange, letterIndex int) (string, string) {
	for j := len(colorRange) - 1; j >= 0; j-- {
		cr := colorRange[j]
		if letterIndex >= cr.Start && letterIndex < cr.End {
			// The letter at the current index, j, should be colored with the current color
			return TerminalColorEscape(cr.Color), color.RESET
		}
	}
	return "", ""
}

// TerminalColorEscape returns the ANSI color code escape sequence for the given RGB color.
// Notes:
// The ANSI escape sequence "\x1b[38;2;{r};{g};{b}m" is used to set the text color, where:
// {r}, {g}, and {b} are the red, green, and blue color values, respectively
// `\x1b[` is the Control Sequence Introducer (CSI).
// `38` tells the terminal that we’re setting the foreground color.
// `2` specifies that we’re using the RGB mode.
// Another escape sequence `\x1b[0m` should be used to reset the text color to default
func TerminalColorEscape(c color.Color) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
}
