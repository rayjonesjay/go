package graphics

import (
	"fmt"
	"os"
	"strings"

	"ascii/caret"
	color "ascii/colors"
	"ascii/data"
	"ascii/fmtx"
	"ascii/terminal"
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

// DrawContext holds the necessary drawing information based on the arguments/options supplied from the command-line
type DrawContext struct {
	Options    data.Options
	GMap       GMap
	TermWidth  int
	IsTerminal bool
}

// Define constants signifying the alignment of text in the terminal
const (
	AlignRight   = "right"
	AlignCenter  = "center"
	AlignJustify = "justify"
)

// Cache the (character -> graphics) in a map structure
var styleCache = make(StyleMap)

// Draw takes a list of ASCII text to be displayed with their respective banner style, maps all the text to their
// respective banner graphics style, and returns a string that draws the graphics
func Draw(draw data.DrawInfo) string {
	var b strings.Builder
	var drawCaret caret.Caret

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

	ctx := DrawContext{
		Options: draw.Options,
		GMap:    GetMap(draw.Style),
	}
	if terminal.IsTerminal() {
		// Initialize the drawing context terminal width (TermWidth)
		termWidth, _, err := terminal.GetTerminalSize()
		if err != nil {
			fmtx.FatalErrorf("error getting terminal size\n%v\n", err)
		}
		ctx.TermWidth = termWidth
		ctx.IsTerminal = true
	} else {
		ctx.TermWidth = -1
	}

	// The current text may be on different lines, if so, we may need to advance the drawCaret to a new line
	lines := strings.Split(draw.Text, "\n")
	for i, currentLine := range lines {
		if i == 0 {
			drawCaret = Drawln(drawCaret, currentLine, ctx)
		} else {
			// Write the previous line, we are yet to write another line
			output := SPrintCaret(drawCaret)
			b.WriteString(output)
			// Prepare to write the next line
			b.WriteRune('\n')
			drawCaret = nil
			drawCaret = Drawln(drawCaret, currentLine, ctx)
		}
	}

	return finalize()
}

// Drawln is a helper function used by the [Draw] function to draw some line of text from the current caret position.
// This therefore, assumes that s is strictly a line of text, and, thus, does not contain any newline characters
// This also expects a map of the ASCII characters to their respective art graphics
func Drawln(lineCaret []string, s string, ctx DrawContext) []string {
	// A lineCaret should ideally be 8 lines, we model the 8 lines with a slice of 8 strings
	if lineCaret == nil || len(lineCaret) < 8 {
		buffer := make([]string, 8)
		copy(buffer, lineCaret)
		lineCaret = buffer
	}

	if s == "" {
		return lineCaret
	}

	return caret.Append(lineCaret, drawln(lineCaret, s, ctx))
}

// draw the text in a single line, taking into account the width of the terminal and the expected text alignment
func drawln(lineCaret caret.Caret, s string, ctx DrawContext) caret.Caret {
	// Split this line into words
	words := strings.Split(s, " ")

	lineLength := 0
	termWidth := ctx.TermWidth
	// We can only align text on the terminal
	isTerminal := ctx.IsTerminal
	align := ctx.Options.Align

	// Make the graphic representation of each word separately
	var wordCarets []caret.Caret
	for _, w := range words {
		wordCaret := drawWord(w, ctx)
		lineLength += wordCaret.Size
		wordCarets = append(wordCarets, wordCaret.Caret)
	}

	// Calculate the size of the graphics of spaces between words
	spaceWidth := (len(words) - 1) * spaceSize(ctx.GMap)
	if spaceWidth < 0 {
		spaceWidth = 0
	}

	// Calculate the terminal's available extra spaces after drawing the words and spaces
	termExtraSpace := termWidth - lineLength - spaceWidth
	if termExtraSpace < 0 && isTerminal {
		fmtx.FatalErrorf(
			"can't fit graphics to terminal; increase terminal size!\ncurrent terminal width: %d\n"+
				"expected terminal width: %d\n", termWidth, lineLength+spaceWidth,
		)
	}

	// In case we need to justify the text
	justifySpacers, extraJustifySpacers := Justify(len(words), termExtraSpace)

	// Draw the words and the spaces between them to the line caret
	for i, wc := range wordCarets {
		if i != 0 {
			spaceGraphics := ctx.GMap[' ']
			lineCaret = caret.Append(lineCaret, spaceGraphics)
			if align == AlignJustify && (justifySpacers > 0 || extraJustifySpacers > 0) && isTerminal {
				n := justifySpacers
				if extraJustifySpacers > 0 {
					extraJustifySpacers--
					n++
				}
				lineCaret = caret.Append(lineCaret, caret.NSpaceCaret(n))
			}
		}
		lineCaret = caret.Append(lineCaret, wc)
	}

	if align == AlignRight && termExtraSpace > 0 && isTerminal {
		lineCaret = caret.Append(caret.NSpaceCaret(termExtraSpace), lineCaret)
	} else if align == AlignCenter && termExtraSpace > 0 && isTerminal {
		leftWidth := termExtraSpace / 2
		lineCaret = caret.Append(caret.NSpaceCaret(leftWidth), lineCaret)
	}

	return lineCaret
}

// returns the width of drawing the space character based on the given banner file's (letter -> graphics) map
func spaceSize(m GMap) int {
	graphics, ok := m[' ']
	if ok {
		return caret.LargestLength(graphics)
	} else {
		fmtx.FatalErrorf("Couldn't find the ASCII graphic for space character\n")
		return 0
	}
}

// draw the word s, keeping track of the actual size of the caret excluding ANSI color escape codes
func drawWord(s string, ctx DrawContext) (sizedWordCaret caret.SizedCaret) {
	colorRangeList := ColorRangeList(s, ctx.Options.ColorFlags)
	wordCaret := caret.NewCaret()
	// Map each ASCII character to its graphics, and append to the current caret position
	for i, char := range s {
		charGraphics, ok := ctx.GMap[char]
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
		sizedWordCaret.Size += caret.LargestLength(charGraphics)
		// Should the current letter be colored?
		var colorCode, resetCode string
		if ctx.IsTerminal {
			// Only use ANSI color escape sequences on the terminal
			colorCode, resetCode = letterColor(colorRangeList, i)
		}
		// Append the current character's graphics to its respective line in the caret
		for j, line := range charGraphics {
			wordCaret[j] = wordCaret[j] + colorCode + line + resetCode
		}
	}

	sizedWordCaret.Caret = wordCaret
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
func GetMap(style string) GMap {
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
		if strings.Contains(cf.Substr, `\n`) {
			fmt.Fprintf(os.Stderr, "substring \"%s\" cannot contain whitespace characters\n", cf.Substr)
			os.Exit(1)
		}
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
			return TerminalColorEscape(cr.Color), color.Reset
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
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.Red, c.Green, c.Blue)
}
