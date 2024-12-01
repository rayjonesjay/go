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
	// format
	// sort
	// compute
	fd , _ := os.Open("input.txt")
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
	fmt.Println(distance)
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
