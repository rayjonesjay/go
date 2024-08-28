package main

import (
	"fmt"
)

type point2d struct {
	x int
	y int
}

// value receiver
func (p point2d) return_y_axis() int {
	return p.y
}


func (p point2d) return_x_axis() int {
	return p.x
}

type coordinates interface {
	return_x_axis() int
	return_y_axis() int
}

func findcoordinates(a coordinates) {
	fmt.Println(a.return_x_axis(), a.return_y_axis())
}
func main() {
	x := point2d{1, 2}
	findcoordinates(x)
}
