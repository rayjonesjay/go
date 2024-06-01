package args

import "ascii/output"


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

// Takes the flag '--output=file.txt' together with text and style to be printed
func Parse(args []string) ([]DrawInfo, string) {
	length_of_arguments := len(args)

	outputFile := ""

	// check if flag was passed and is valid
	if IsValidFlag(args) {
		flagAndFile := args[0]
		OutputFile, inspectError := InspectFlagAndFile(flagAndFile)
		if inspectError == nil && OutputFile != "" {
			outputFile = outputFile
		}
		args = args[1:]
		length_of_arguments = (length_of_arguments - 1)
	}

	if length_of_arguments < 1 {
		return nil, outputFile
	} else if length_of_arguments == 1 {

		text := args[0]
		return []DrawInfo{{Text: Escape(text), Style: Standard}}, outputFile

	}

	return []DrawInfo{}, outputFile
}
