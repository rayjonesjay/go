package special

import (
	"strings"
)

// SlashB applies the escape sequence \b on the given string
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

// SlashR applies the escape sequence \r on the given string
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

// SlashF applies the escape sequence \f on the given string
func SlashF(s string) string {
	return slashFSlashV(s, 'f')
}

// SlashV applies the escape sequence \v on the given string
func SlashV(s string) string {
	return slashFSlashV(s, 'v')
}

// slashFSlashV is a super function to handle either of the escape sequences \f and \v
func slashFSlashV(s string, escape rune) string {
	sp := strings.Split(s, `\`+string(escape))
	indents := make([]int, len(sp))

	for i, s := range sp {
		if i == 0 {
			indents[i] = len(s)
		} else {
			var b strings.Builder
			if s != "" {
				// Don't write indent spaces if the current split is an empty string
				b.WriteString(strings.Repeat(" ", indents[i-1]))
			}

			b.WriteString(s)
			sp[i] = b.String()
			indents[i] = indents[i-1] + len(s)
		}
	}

	return strings.Join(sp, "\n")
}
