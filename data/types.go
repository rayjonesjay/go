package data

import "ascii/colors"

// DrawInfo holds the text to be drawn, and with which style it is to be drawn
type DrawInfo struct {
	Text  string
	Style string
}

// ColorInfo defines the substring (Substr) in a text that need to be colored by the specified Color.
// For example, given the color flag, `--color=red kit`,
// then all substrings `kit` in the given text need be colored `red`
type ColorInfo struct {
	Color  colors.Color
	Substr string
}

// Options keeps track of the command-line options passed to the program
type Options struct {
	ColorFlags []ColorInfo
	Align      string
	Output     string
}
