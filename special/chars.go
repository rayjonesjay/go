package special_chars

// handling \b
func SlashB(s string) string {
	result := ""

	// change the string to an array of runes
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// check if the current rune is not at the last index and it's \ and the next is b
		if i+1 < len(runes) && runes[i] == '\\' && runes[i+1] == 'b' && i+1 != len(runes)-1 {
			// remove the last element concatinated in the string "result"
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
			i++ // skip the '\' and 'b' character
		} else {
			result += string(runes[i])
		}
	}

	// when \b is at the end of the string, ignore
	for i, v := range result {
		if v == '\\' && result[i+1] == 'b' {
			result = result[:len(result)-2]
		}
	}
	return result
}

// handling \r
// handling \f
// handling \v
