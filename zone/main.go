package main 

import (
	"fmt"
)

func main() {
		// s := "()()()()(()()()((()))))((())))"
		// fmt.Println(isValidParentheses(s))	

		// nums := []int{1,1}
		// fmt.Println(findErrorNums(nums))
		fmt.Println(pivotInteger(8))
}
func pivotInteger(n int) int {
    
    for i := 1; i <= n ; i++{
        l := i * (i+1)/2
        b := n-i
        fmt.Println(b)
        r := b * (b+1)/2

        if l==r{
        	fmt.Println(l,r)
        	return i
        }
    }

    return -1
}



// func findErrorNums(nums []int) []int {
    
//     mappy := make(map[int]int)

//     for _, num := range nums {
//         mappy[num]++
//     }

//     res := make([]int,2)

//     for i := 1; i < len(nums); i++{
//         if mappy[i] == 2{
//             res[0] = i
//         }else if mappy[i] == 0{
//         	fmt.Println(mappy[i])
//             res[1]= i
//         }
//     }
//     return res
// }

// func isValidParentheses(s string) int {
// 	slice := []rune(s)
// 	stack := []rune{}

// 	if len(slice) ==  1{
// 		fmt.Println("false slice len is 1")
// 		return -1
// 	}

// 	if slice[0] == ')'{
// 		slice = slice[1:]
// 	}
// 	if slice[len(slice)-1]=='('{
// 		slice = slice[:len(slice)-1]
// 	}
// 	length:=0

// 	for _, bracket := range slice {
// 		if bracket == '('{
// 			stack = append(stack, bracket)
// 		}else {
// 			// //if its closing first check if the stack is eligible for popping
// 			// if len(stack) == 0 {
// 			// 	return -1 // there is no matching closing bracket
// 			// }
// 			if len(stack) > 0 && stack[len(stack)-1] == '('{
// 				stack = stack[:len(stack)-1]
// 				length+=2
// 			}
// 		}
// 	}
// 	// fmt.Println(length)
// 	return length

// }
// // remove an element at given index
// func popAt(slice []rune, index int) []rune {
// 	result := []rune{}
// 	result = append(slice[:index], slice[index+1:]...)
// 	return result
// }


// // insert an item at index in slice
// func InsertItemAt(slice []rune, item index, index int) []rune {
// 	result := []rune{}
// 	result = append(slice[:index], append([]rune{item}, slice[index+1:]...)...)
// 	return result
// }
