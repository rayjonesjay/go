package main

import (
	"bufio"
	"strconv"
	"os"
	"strings"
	"slices"
	"fmt"
)

func main(){
	// read
	// format
	// sort
	// compute
	
	// read
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("go run main.go <input.file.txt>")
		return
	}
	fd , _ := os.Open(arg[0])
	scanner := bufio.NewScanner(fd)
	leftArr := []int{}
	rightArr := []int{}
	for scanner.Scan(){
		line := strings.TrimSpace(scanner.Text())
		arr := strings.Fields(line)
		left , right := arr[0],arr[1]
		l , _ := strconv.Atoi(left)
		r , _ := strconv.Atoi(right)
		
		leftArr = append(leftArr,l)
		rightArr = append(rightArr,r)
	}

	// sort
	slices.Sort(leftArr)
	slices.Sort(rightArr)
	
	// compute
	distance := 0

	// left and right array are same lenght
	for i := 0; i < len(leftArr); i++{
		distance = distance + absolute(leftArr[i]-rightArr[i])	
	}
	fmt.Println("distance",distance)


	// question2
	var sscore int64 = 0
	frequencyRight := make(map[int]int64)
	
	for _, n := range rightArr {
		frequencyRight[n]++
	}
	for _, n := range leftArr {
		sscore = sscore + (frequencyRight[n]*int64(n))
	}
	fmt.Println("similarity score",sscore)	
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
