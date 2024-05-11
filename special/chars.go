package special

import "strings"

// SlashB handling \b
func SlashB(s string) string {
	result := ""

	// change the string to an array of runes
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// check if the current rune is not at the last index, and it's '\' and the next is 'b'
		if i+1 < len(runes) && runes[i] == '\\' && runes[i+1] == 'b' && i+1 != len(runes)-1 {
			// remove the last element concatenated in the string "result"
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

// SlashR handling \r
func SlashR(s string) string {
	result := ""
	arr := strings.Split(s, "\\r")

	for i := 0; i < len(arr); i++ {
		// if there is no \r character in the string then the arr will only have one string
		// in that case just return the string
		if len(arr) == 1 {
			return arr[0]
		}

		// check if the string after \r is longer than the string before
		// if so, update the string before to be the string after \r
		if i+1 < len(arr) && len(arr[i+1]) > len(arr[i]) {
			if len(result) > 0 {
				result = ""
			}
			result += arr[i+1]
		}

		if i+1 < len(arr) {
			// if result is empty, update it to the string at the current index
			if len(result) < 1 {
				result = arr[i]
			}

			// prepend the string after \r as you overwrite to the string before
			result = arr[i+1] + result[len(arr[i+1]):]
		}
	}
	return result
}

// SlashFSlashV handling \f and \v
func SlashFSlashV(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && (runes[i] == '\\' && runes[i+1] == 'f') || (runes[i] == '\\' && runes[i+1] == 'v') {
			runes[i] = '\\'
			runes[i+1] = 'n'
		}
	}

	// fmt.Println(string(runes))

	collect := ""
	result := ""
	spaces := ""

	for j := 0; j < len(runes); j++ {
		if j+1 < len(runes) && runes[j] == '\\' && runes[j+1] == 'n' {
			if (len(collect) > 0 && j >= 2 && runes[j-1] != 'n' && runes[j-2] != '\\') || j == 1 {
				for k := 0; k < len(collect); k++ {
					spaces += " "
				}
			} else if len(collect) < 0 {
				result += "\n"
			}

			if j+3 < len(runes) && runes[j+2] == '\\' && runes[j+3] == 'n' {
				result += collect + "\n"
			} else {
				result += collect + "\n" + spaces
			}

			if (j >= 2 && runes[j-1] != 'n' && runes[j-2] != '\\') || j == 1 {
				collect = ""
			}
			j++
		} else {
			collect += string(runes[j])
		}
	}

	return result + collect
}
