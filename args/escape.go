package args

import (
	"strconv"
	"strings"
)

// Escape given a command line argument, or rather any such string, performs interpretation of the
// following backslash escapes:
//
//	\\     backslash
//
//	\a     alert (BEL)
//
//	\b     backspace
//
//	\f     form feed
//
//	\n     new line
//
//	\r     carriage return
//
//	\t     horizontal tab
//
//	\v     vertical tab
//
//	\0NNN  byte with octal value NNN (1 to 3 digits)
//
//	\xHH   byte with hexadecimal value HH (1 to 2 digits)
//
// Note: Any octal or hexadecimal values of ASCII characters that cannot be printed will be ignored
func Escape(arg string) string {
	ss := []rune(arg)
	l := len(ss)

	//  '\a' (bell)
	//  '\b' (backspace)
	//  '\t' (horizontal tab)
	//  '\n' (new line)
	//  '\v' (vertical tab)
	//  '\f' (form feed)
	//  '\r' (carriage ret)
	repMap := map[rune]rune{'\\': '\\',
		'a': '\a',
		'b': '\b',
		't': '\t',
		'n': '\n',
		'v': '\v',
		'f': '\f',
		'r': '\r',
	}

	for i := 0; i < l; i++ {
		c := ss[i]
		j := i + 1
		if c == '\\' && j < l {
			nc := ss[j]
			rep, ok := repMap[nc]
			if ok {
				ss[i] = rep
				ss[j] = -1
			} else if !replaceEscape(i, ss, 'x', 2, HexStringToDecimal) {
				replaceEscape(i, ss, '0', 3, OctalStringToDecimal)
			}
		}

	}

	return ToString(ss)
}

// Given that the ith character of the rune slice ss is the '\' character, and that the next character
// is the escape character, uses the function f to interpret the next n characters after the escape character
// as an integer representing a rune; this then replaces the whole escape sequence with this rune.
// For example, this can thus be used to replace all hexadecimal escaped ASCII characters such as `\x61`
// with its actual rune `a`
func replaceEscape(i int, ss []rune, escape rune, n int, f func(string) (decimal int, ok bool)) (match bool) {
	j := i + 1
	k := j + 1
	// \x61
	// 0123
	// i123
	// ij12
	// ijk1
	l := len(ss)
	if ss[j] == escape && l > j+n {
		hex := ss[k : k+n]
		decimal, ok := f(string(hex))
		if ok {
			match = true
			ss[i] = rune(decimal)
			for g := j; g < k+n; g++ {
				ss[g] = -1
				i++
			}
		}
	}
	return
}

// HexStringToDecimal given a string representing a hexadecimal number, returns the decimal
// representation of the input hexadecimal
func HexStringToDecimal(hex string) (decimal int, ok bool) {
	return BaseStringToDecimal(16, hex)
}

// OctalStringToDecimal given a string representing an octal number, returns the decimal
// representation of the input octal number
func OctalStringToDecimal(oct string) (decimal int, ok bool) {
	return BaseStringToDecimal(8, oct)
}

// BaseStringToDecimal given a string representing a number to a given base, returns the decimal
// representation of the input number
func BaseStringToDecimal(base int, s string) (decimal int, ok bool) {
	conv, err := strconv.ParseInt(s, base, 32)
	if err == nil {
		decimal = int(conv)
		if decimal >= 0 && decimal <= 127 {
			ok = true
		}
	}
	return
}

// ToString given a slice of runes, returns a string representation of the runes, ignoring the '\0' (null character)
func ToString(input []rune) string {
	var b strings.Builder
	for _, c := range input {
		if c > 0 {
			b.WriteRune(c)
		}
	}
	return b.String()
}
