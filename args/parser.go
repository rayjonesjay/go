package args

import (
	"fmt"
	"os"
)

const (
	Shadow     = "shadow"
	Standard   = "standard"
	Thinkertoy = "thinkertoy"
)

// DrawInfo holds the text to be drawn, and with which style it is to be drawn
type DrawInfo struct {
	Text  string
	Style string
}

// Parse Given commandline arguments, excluding the program name, returns a list of all the extracted
// text to be drawn with their respective styles
func Parse(args []string) []DrawInfo {
	l := len(args)

	if l < 1 {
		// Program didn't receive any text to be printed, exit with usage instructions
		return nil
	} else if l == 1 {
		// Program received some text to be printed, use the standard banner to print the ASCII-ART
		text := args[0]
		return []DrawInfo{{Text: Escape(text), Style: Standard}}
	} else {

		// Program received a series of texts to be printed, with banner style specified for consecutive texts
		var out []DrawInfo

		for textPosition := 0; textPosition < l; textPosition += 2 {

			text := args[textPosition]

			// default style is Standard
			style := Standard

			// check if style is provided
			if textPosition+1 < l {
				// style = args[textPosition]
				switch args[textPosition+1] {

				case Standard, Shadow, Thinkertoy:
					style = args[textPosition+1]
				default:
					fmt.Fprintf(os.Stderr, "Style argument not recognized! Passed -> %s Expected -> shadow|standard|thinkertoy\n", args[textPosition+1])
					os.Exit(1)
				}
			}
			out = append(out, DrawInfo{Text: Escape(text), Style: style})
		}

		// for i := 0; i < l; i += 2 {
		// 	j := i + 1
		// 	text := args[i]
		// 	style := Standard
		// 	if j < l {
		// 		style = args[j]
		// 	}

		// 	out = append(out, DrawInfo{Text: Escape(text), Style: style})
		// }
		return out
	}
}
