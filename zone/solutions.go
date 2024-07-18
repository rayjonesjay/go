package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	print "github.com/01-edu/z01"
)

// hello e a
func Searchreplace(s string, a, b string) string {
	result := ""
	for _, char := range s {
		if char == rune(a[0]) {
			result += string(b[0])
		}else {
			result += string(char)
		}
	}
	return result
}

// removes the previous character to every occurrence of \b
func BackSlash(s string) string {
	// while string contains \b
	for strings.Contains(s, `\b`) {
		if string(s[0]) == (`\b`) {
			s = s[1:]
		}
		// get index of where the \b is
		index := strings.Index(s, `\b`)
		before := s[:index]
		after := s[index+1:]
		s = before[index-1:] + after
	}
	return s
}

func abs(number float64) float64 {
	if number < 0 {
		return -number
	}
	return number
}

func Sqrt(number float64) float64 {
	const prec = 1e-10
	currentGuess := number
	nextGuess := 0.0
	for {
		nextGuess = 0.5 * (currentGuess + number/currentGuess)
		if abs(nextGuess-currentGuess) < prec {
			break
		}
		currentGuess = nextGuess
	}
	// fmt.Println(count)
	return nextGuess
}

func FindErrorNums(nums []int) []int {
	mappy := make(map[int]int)

	for _, num := range nums {
		mappy[num]++
	}
	fmt.Println(mappy)
	res := make([]int, 2)
	if len(mappy) == 1 {
		res[0] = nums[0]
		res[1] = nums[0] + 1
		return res
	}

	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		if mappy[i] == 2 {
			// fmt.Println(res)
			res[0] = i
		} else if mappy[i] == 0 {
			fmt.Println(">>>", mappy[i])
			res[1] = i
		}
	}
	return res
}

func IsValidParentheses(s string) int {
	slice := []rune(s)
	stack := []rune{}

	if len(slice) == 1 {
		fmt.Println("false slice len is 1")
		return -1
	}

	if slice[0] == ')' {
		slice = slice[1:]
	}
	if slice[len(slice)-1] == '(' {
		slice = slice[:len(slice)-1]
	}
	length := 0

	for _, bracket := range slice {
		if bracket == '(' {
			stack = append(stack, bracket)
		} else {
			// //if its closing first check if the stack is eligible for popping
			// if len(stack) == 0 {
			// 	return -1 // there is no matching closing bracket
			// }
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				length += 2
			}
		}
	}
	// fmt.Println(length)
	return length
}

func ShiftWord(word string, shift int) string {
	slice := []rune(word)
	for i, char := range slice {
		if char >= 'a' && char <= 'z' {
			slice[i] = (char-'a'+rune(shift))%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			slice[i] = (char-'A'+rune(shift))%26 + 'A'
		} else {
			slice[i] = char
		}
	}
	fmt.Println(string(slice))
	return string(slice)
}

// remove an element at given index
func PopAt(slice []rune, index int) []rune {
	result := []rune{}
	result = append(slice[:index], slice[index+1:]...)
	return result
}

// insert an item at index in slice
func InsertItemAt(slice []rune, item rune, index int) []rune {
	result := []rune{}
	result = append(slice[:index], append([]rune{item}, slice[index+1:]...)...)
	return result
}

func RevParam() {
	result := "abcdefghijklmnopqrstuvwxyz"

	f := []rune{}
	for i, char := range result {
		// uppercase the elements whose index is odd
		if i%2 == 0 {
			f = append(f, toupper(char))
		} else if i%2 == 1 {
			f = append(f, char)
		}
	}

	// reverse the string
	for i, j := 0, len(f)-1; i < len(f)/2; i, j = i+1, j-1 {
		f[i], f[j] = f[j], f[i]
	}
	for _, char := range f {
		print.PrintRune(char)
	}
	print.PrintRune('\n')
}

func toupper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r = r - 32
	}
	return r
}

func FirstParam() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	result := args[0]
	for _, char := range result {
		print.PrintRune(char)
	}
	print.PrintRune('\n')
}

func LastParam() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	result := args[len(args)-1]
	for _, char := range result {
		print.PrintRune(char)
	}
	print.PrintRune('\n')
}



func itoa(n int) string {
	// check if n is negative or positive
	isNegative := false
	if n < 0 {
		isNegative = true
		// convert it to positve
		n = -n
	} else if n >= 0 {
		n = n * 1
	}

	// slice to store the appened runes
	result := []rune{}
	// extract last digits for processing
	for n > 0 {
		lastDigit := n % 10
		result = append(result, rune(lastDigit+48))
		n = n / 10
	}

	// reverse our result before returing
	rev := func(s []rune) []rune {
		i, j := 0, len(s)-1
		for i < len(s)/2 {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
		return s
	}
	result = rev(result)

	if isNegative {
		result = append([]rune{'-'}, result[0:]...)
	}

	return string(result)
}

func ParamCount() {
	args := os.Args[1:]
	n := len(args)
	if n < 1 {
		return
	}

	str := itoa(n)
	for _, num := range str {
		print.PrintRune(num)
	}
	print.PrintRune('\n')
}

func CountDown() {
	num := "0123456789"
	result := []rune(num)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	for _, num := range result {
		print.PrintRune(num)
	}
	print.PrintRune('\n')
}

func Atoi(s string) int {
	// when s is zero lets just return 0
	if s == "0" {
		return 0
	}

	// check if number has a sign + or -
	isNegative := false

	// so the sign comes the first before the number
	if s[0] == '-' {
		// if the sign is - then it is negative
		isNegative = true
		// remove the sign for processing
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:] // exclude the positive sign
	}

	result := 0
	for _, num := range s {
		// num is of rune type
		result = result*10 + int(num-48)
	}

	if isNegative {
		return -result
	}
	return result
}

func InsertAt(s string, index int, r rune) string {
	if index < 0 {
		index = 0
	} else if index >= len(s) {
		index = len(s)
	}

	runes := []rune(s)
	result := append(runes[:index], append([]rune{r}, runes[index:]...)...)

	return string(result)
}

func Pop(s string, index int) string {
	runes := []rune(s)
	result := append(runes[:index], runes[index+1:]...)
	return string(result)
}

func WordMatch() string {
	args := os.Args[1:]

	n := len(args)

	if n < 2 {
		return ""
	}
	firstword := args[0]
	secondword := args[1]
	i, j := 0, 0
	for i < len(firstword) && j < len(secondword) {
		if firstword[i] == secondword[j] {
			i++
		}
		j++

		if i == len(firstword) {
			return firstword
		}
	}
	return ""
}

func LastRune(s string) rune {
	runeSlice := []rune(s)
	return runeSlice[len([]rune(s))-1]
}

func ReduceInt(a []int, f func(int, int) int) int {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	return result
}

func Rot13(s string) string {
	result := []rune{}
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result = append(result, 'a'+(char-'a'+(13))%26)
		} else if char >= 'A' && char <= 'Z' {
			result = append(result, 'A'+(char-'A'+(13))%26)
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}

func AlphaMirror(s string) string {
	if s == "" {
		return ""
	}
	lowerMap := make(map[rune]rune, 26)

	j := 122
	for i := 'a'; i < 'z'; i++ {
		lowerMap[rune(i)] = rune(j)
		j--
	}

	upperMap := make(map[rune]rune, 26)
	j = 90
	for i := 'A'; i < 'Z'; i++ {
		upperMap[rune(i)] = rune(j)
		j--
	}
	result := []rune{}
	for _, char := range s {
		if IsLower(char) {
			result = append(result, lowerMap[char])
		} else if IsUpper(char) {
			result = append(result, upperMap[char])
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func Mirror(r rune) rune {
	if r >= 'a' && r < 'z' {
		return 'a' + 'z' - r
	} else if r >= 'A' && r < 'Z' {
		return 'A' + 'Z' - r
	}

	return r
}

func IsLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsUpper(r rune) bool {
	if r >= 'A' && r <= 'A' {
		return true
	}
	return false
}

func AlphaCount(s string) int {
	alphacount := 0

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			alphacount++
		}
	}
	return alphacount
}

func AlphaPosition(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	} else if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 1
	}
	return -1
}

func Chunk(slice []int, size int) {
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}
	if size == 0 {
		fmt.Println()
		return
	}

	// slice of slices
	ss := [][]int{}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		ss = append(ss, slice[i:end])
	}
	fmt.Println(ss)
}

func Compare_(a, b string) int {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}

	if a == b {
		return 0
	}

	na := len(a)
	nb := len(b)
	if na > nb && StartsWith(a, b) {
		return 1
	}

	if na > nb && EndsWith(a, b) {
		return -1
	}

	return -1
}

func StartsWith(a, b string) bool {
	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EndsWith(a, b string) bool {
	for i := len(b) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func FoldInt(f func(int, int) int, a []int, n int) int {
	for _, char := range a {
		n = f(n, char)
	}
	return n
}

func FindPrevPrime(n int) int {
	for i := n; i >= 2; i++ {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i < i*i; i++ {
		if i%2 == 0 || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}

// TrialDivisionAlgorithm
func TrialDivision(n int) bool {
	// function to check if a number is prime using trial division algorithm

	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// find the squareroot of n
	k := int(math.Sqrt(float64(n)))
	i := 5
	for i < k {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func Gcd(a, b int) int {
	// using euclidean algo
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// i will improve it later when i get a girlfriend
func RayGCD(a, b int) int {
	var i int = 2      // start from 2
	var result int = 1 // one because if i set it to zero the result will always be zero
	for a != 0 || b != 0 {
		if a%i == 0 && b%i == 0 {
			a = a / i
			b = b / i
			result = result * i
		}
		i++ // increment i
	}
	return result
}

func Multiple_gcd(List []int) int {
	if len(List) == 0 {
		return 0
	}
	res := List[0]
	for _, num := range List[1:] {
		res = Gcd(res, num)
	}
	return res
}

func Multiple_gcd2(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	res := numbers[0]
	for _, num := range numbers[1:] {
		res = Gcd(res, num)
	}
	return res
}

func RomanNumbers() (string, string) {
	args := os.Args[1:]
	if len(args) != 1 {
		return "", ""
	}
	num, err := strconv.Atoi(args[0])
	if err != nil || num < 0 || num > 4000 {
		fmt.Println("Error")
		return "", ""
	}
	return "", ""
}


func IsPowerOf2_better(n int) bool {
	return n > 0 && (n & (n-1))==0
}

func IsPowerOf2(n int) bool {
	count := 0
	if n%2 == 1 {
		return false
	}
	for n >= 0 {
		n /= 2
		count++
	}
	if math.Pow(2, float64(count)) == float64(n) {
		return false
	}
	return false
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}

// This functions behaves the same way as cariage return
func CarriageReturn(s string) string {
	for strings.Contains(s, `\r`) {
		index := strings.Index(s, `\r`)
		s = s[index+1:] + s[:index]
	}
	return s
}


func TabMult(s string) {
	num, _ := strconv.Atoi(s) // make your own atoi
	fmt.Println(num)
	for i := 1; i <= 9; i++ {
		result := i * num
		res_string := itoa(result)
		i_string := itoa(i)
		num_string := itoa(num)
		printString(i_string)
		print.PrintRune(' ')
		print.PrintRune('x')
		print.PrintRune(' ')
		printString(num_string)
		print.PrintRune(' ')
		print.PrintRune('=')
		print.PrintRune(' ')
		printString(res_string)
		print.PrintRune('\n')
	}
}

func printString(s string) {
	for _, char := range s {
		print.PrintRune(char)
	}
}
