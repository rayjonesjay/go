package special_chars

// handling \b
func SlashB(s string) string {
	result := ""

	// change the string to an array of runes
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// check if the current rune is not at the last index and it's \ and the next is b and
		if i+1 < len(runes) && runes[i] == '\\' && runes[i+1] == 'b' {
			// remove the last element concatinated in the string "result"
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
			i++ // skip the 'b' character
		} else {
			result += string(runes[i])
		}
	}
	return result
}

// handling \r
// handling \f
// handling \v
