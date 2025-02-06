package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numWorkers := 3
	numTasks := 10

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	wg := sync.WaitGroup{}
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}
	defer wg.Wait()
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 1; i <= numTasks; i++ {
		result := <-results
		fmt.Printf("result %d\n", result)
	}

}

/*
results is a write only channel
tasks is a read only channle
*/
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("worker %d processing task: %d\n", id, task)
		time.Sleep(2 * time.Second)
		results <- task * task
		fmt.Printf("worker %d took %d\n", id, time.Since(start))
	}
}
