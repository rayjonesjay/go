package special

import (
	"ascii/args"
	"strings"
)

// EscapeB applies the escape sequence \b on the given string
func EscapeB(s string) string {
	out := make([]rune, len(s))

	pos := 0
	for _, r := range s {
		if r == '\b' {
			// Move cursor back one position
			pos = max(pos-1, 0)
		} else {
			// Write the current character at pos
			if pos < len(out) && pos >= 0 {
				out[pos] = r
			}
			pos++
		}
	}

	return args.ToString(out)
}

// EscapeR applies the escape sequence \r on the given string
func EscapeR(s string) string {
	result := ""
	arr := strings.Split(s, "\r")

	// if there is no \r character in the string then the arr will only have one string
	// in that case just return the string
	if len(arr) == 1 {
		return arr[0]
	}

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

// EscapeF applies the escape sequence \f on the given string
func EscapeF(s string) string {
	return escapeVF(s, '\f')
}

// EscapeV applies the escape sequence \v on the given string
func EscapeV(s string) string {
	return escapeVF(s, '\v')
}

// escapeVF is a super function to handle either of the escape sequences \f and \v
func escapeVF(s string, escape rune) string {
	sp := strings.Split(s, string(escape))
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
