package main

import (
	"fmt"
	"leetcode/solutions"
)

func main() {
	test := []string{"flower","flow","flight"}
	test1 := []string{"dog","racecar","car"}
	fmt.Println(solutions.LongestCommonPrefix14(test1))
	fmt.Println(solutions.LongestCommonPrefix14(test))
}
