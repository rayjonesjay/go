package solutions

import (
	"fmt"
	"sort"
)

/*
Write a function to find the longest common prefix string amongst an
an array of strings.
*/
func LongestCommonPrefix14(strs []string) string {
	var pref string
	// find the shortest word in the slice, by sorting
	sort.Slice(strs, func(i, j int) bool { return len(strs[i]) < len(strs[j]) }) // returns in ascending order
	fmt.Println(strs)
	shortestWord := strs[0]
	l := len(shortestWord)
	for {
		// check if each word in strs contains the prefix, if not remove the last character from the prefix
		// and try again until the shortest word has a length of 0, it is at this point we return an empty string as the prefix
		if shortestWord == "" {	
			break
		}
		if ContainsAll(strs, shortestWord) == len(strs){
				return shortestWord
		}
		l--
		shortestWord = shortestWord[:l]
		}
	return pref
}

// ContainsAll counts all strings in the list which contain pref
func ContainsAll(list []string, pref string) int {
	var count int = 0
	for _, w := range list {
	if Contains(w,pref){
			count++
		}
	}
	return count
}

// Contains reports whether a begins strictly with what is b
func Contains(a, b string) bool {
	return a[:len(b)] == b
}
