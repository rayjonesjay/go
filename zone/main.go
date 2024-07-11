package main

import (
	"fmt"
	"os"
)

func main() {

    // an IIFE function
    // number := 2
    if !func(a,b int) bool {
        if a < b {
            return false
        }else{
            return true
        }
    }(1,2){
        fmt.Println("a is less than b")
    }
    args := os.Args[1:]
    a:=args[0]
    b:=args[1]
    c:=args[2]
    fmt.Println(searchreplace(a,b,c))
}