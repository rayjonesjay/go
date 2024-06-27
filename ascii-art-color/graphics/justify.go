package graphics

// Justify calculates and returns the amount of space to be equally distributed between words,
// and if there are terminal extra spaces than gaps between words,
// the extra amount of space to be distributed one at a time to the gaps from left to right.
// Example:
// Given the string |hi there to         |, let's justify it:
//
//	Total width: 20 characters
//	Length of text: 11 characters
//	Required spaces: 9 spaces
//	Words: hi, there, to
//
// There are two gaps between the words hi there to. We need to distribute 9 spaces between these gaps:
//
//	hi there to (11 characters)
//	9 spaces need to be distributed into 2 gaps.
//	Each gap gets an equal base of 9 // 2 = 4 spaces (The `//` represents integer division),
//	and there's 9 % 2 = 1 remaining space to distribute.
//
// Therefore:
//
//	First gap (hi and there): 4 + 1 = 5 spaces
//	Second gap (there and to): 4 spaces
//
// Hence; |hi      there     to|
func Justify(numberOfWords, termExtraSpace int) (justifySpacers, extraJustifySpacers int) {
	numberOfSpaces := numberOfWords - 1
	if numberOfSpaces > 0 && termExtraSpace > 0 {
		justifySpacers = termExtraSpace / numberOfSpaces
		extraJustifySpacers = termExtraSpace % numberOfSpaces
	}
	return
}
