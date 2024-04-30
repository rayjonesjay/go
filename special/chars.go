package special

import "strings"

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

// handling \0
func SlashZero(s string) string {
	result := ""

	// change the string to an array of runes
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// check if the current rune is not at the last index and it's '\' and the next is '0'
		if i+1 < len(runes) && runes[i] == '\\' && runes[i+1] == '0' {
			i++ // skip the '\' and '0' character
		} else {
			result += string(runes[i])
		}
	}
	return result
}

// handling \r
func SlashR(s string) string {
	result := ""
	arr := strings.Split(s, "\\r")

	for i := 0; i < len(arr); i++ {
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

// handling \f and \v
func SlashFSlashV(s string) string {
	// spaces will hold spaces for indentation
	spaces := ""

	// if the last character of the string is a form feed ('\f') or vertical tab ('\v'), add a space to ensure the last line is properly handled
	if s[len(s)-1] == '\f' || s[len(s)-1] == '\v' {
		s += " "
	}

	// split the string into fields using the Split function, which splits using 2 delimiters, \f and \v
	arr := strings.FieldsFunc(s, Split)

	// iterate over the splitted array of strings
	for i := 0; i < len(arr); i++ {
		// for the second line calculate the number of spaces needed for indentation
		if i == 0 {
			for j := 1; j <= len(arr[i]); j++ {
				spaces += " "
			}	
		} else {
			// for the rest of the lines, the number of spaces will be less one to avoid counting the \n character
			for j := 1; j <= len(arr[i])-1; j++ {
				spaces += " "
			}
		}

		// prevent panic error of array out of range by ensuring i is not the last index in the array
		if i+1 < len(arr) {
			// add the necessary indentations to the subsequent items
			arr[i+1] = "\n" + spaces + arr[i+1]
		}

		// reset the spaces for the next iteration
		spaces = ""
	}

	// join the array of strings into a single string
	result := strings.Join(arr, "")

	// if the original string ended with a space, remove it from the result
	if s[len(s)-1] == ' ' {
		return result[0:len(result)-1]
	}
	return result
}

// Split is a function used by strings.FieldsFunc to enable splitting using 2 delimiters
func Split(r rune) bool {
	return r == '\f' || r == '\v'
}
