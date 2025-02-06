package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("hello world")

	start := time.Now()
	for i := 0; i < 3; i++ {
		wg.Add(1)
	}
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("go routine 1 done")
		wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("go routine 2 done")
		wg.Done()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("go routine 3 done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("total time elapsed", time.Since(start))
}
