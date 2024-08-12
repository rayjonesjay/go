package Evaluate

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func EvaluateExpression(expression string, from, to int) (output string) {

	// this means the user wants to convert from base X to base 10
	if to == 10 {
		expression = (ToBase10(from, expression))
		return expression
	}

	expression = ToBase10(from, expression)
	return ToBaseN(to, expression)
}

// 10 -> base 3
func ToBaseN(to int, input string) string {
	// using a technique called repeated division to convert to any base
	integer, _ := strconv.Atoi(input)
	if integer == 0 {
		return "0"
	}

	result := ""
	for integer > 0 {
		remainder := integer % to
		if remainder >= 10 {
			result = string('A'+remainder-10) + result
		} else {
			result = Itoa(remainder) + result
		}
		integer = integer / to
	}

	return result
}

func ToBase10(from int, input string) string {
	var digitChecker int = 0
	for _, num := range input {
		if num >= '0' && num <= '9' {
			digitChecker, _ = strconv.Atoi(string(num))
			if digitChecker >= from {
				return fmt.Sprintf("%d not base %d number", digitChecker, from)
			}
		} else if num >= 'A' && num <= 'Z' {
			digitChecker = int(num-'A') + 10
			if digitChecker >= from {
				return fmt.Sprintf("%d not base %d number", digitChecker, from)
			}
		} else if num >= 'a' && num <= 'b' {
			digitChecker = int(num-'a') + 10
			if digitChecker >= from {
				return fmt.Sprintf("%d not base %d number", digitChecker, from)
			}
		} else {
			return fmt.Sprintf("%d not base %d number", num, from)
		}
	}

	if digitChecker >= from {
		return fmt.Sprintf("%d not base %d number", digitChecker, from)
	}

	fractionPart := "0"
	integerPart := input
	isFraction := false
	// right now i wont handle negatives :)

	// check if input has decimal part, first find index of the dot which separates integer and fraction part
	index := strings.Index(input, ".")

	// if the index returned is not -1, meaning the sub string or in our case the dot position was found
	if index != -1 {
		// integer part is on left hand side
		integerPart = input[:index]

		// fraction part is on right hand side
		fractionPart = input[index+1:]

		isFraction = true
	}

	var placeValue float64 = 0
	var result float64 = 0.0
	base := float64(from)

	for i := len(integerPart) - 1; i >= 0; i-- {
		result += float64(int(integerPart[i]-'0') * int((math.Pow(base, placeValue))))
		placeValue++
	}

	// now the fraction part
	placeValue = -1
	fractionResult := 0.0
	for i := 0; i < len(fractionPart)-1; i++ {
		fractionResult += float64(int(fractionPart[i]-'0') * int(math.Pow(base, placeValue)))
		placeValue--
	}

	if isFraction {
		return Itoa(int(result)) + "." + Itoa(int(fractionResult))
	}

	return Itoa(int(result))
}

// Itoa converts an int to its string representation
func Itoa(number int) string {

	// this flag will guide us to know whether the number is negative or positive
	isNegative := false

	// numbers less than 0 are considered negative numbers
	if number < 0 {

		// multiply number with -1 to to make it positive
		number = number * -1

		// set flag to true, indicating our input is negative
		isNegative = true
	}

	// purpose: store the digits as runes for easier conversion to string
	result := []rune{}
	if number == 0 {
		return "0"
	}

	for number > 0 {
		// extract the last digit
		lastDigit := number % 10

		// convert the lastDigit of number and convert it to rune , finally append it to result
		result = append(result, rune(lastDigit+'0'))

		// remove the last digit of number by shifting the number by one place value to the right and discarding the least significant digit
		number = number / 10
	}

	// this is where the flag will help us, if the number was negative append a minus sign. '-'
	if isNegative {
		result = append(result, '-')
	}

	// reverse this whole string since we started appending from the last digit(right) towards the first digit(left)
	for i, j := 0, len(result)-1; i < len(result)/2; i, j = i+1, j-1 {

		// using two pointer technique to shift the digits in place, and its effecient since the loop goes up to len(result)/2 (halfway)
		result[i], result[j] = result[j], result[i]
	}

	// convert the result: a slice of runes to a string
	return string(result)
}
