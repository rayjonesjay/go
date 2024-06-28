package main 

import (
    "fmt"
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

}