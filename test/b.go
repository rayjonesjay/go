package main

import (
	"fmt"
	"sync"
	"time"
)

/*
a semaphore in programming is  a way to control access to a limited resource.
Go doesn't have built-in semaphores, but you can implement one using a buffered
channel

analogy: imagine a parking lot where it has only 5 spaces.
		 each char can only park for a total of 3 seconds.
		 if 3 seconds elapses the car is removed from parking lot.
		 another one is added. so that there is efficient parking
		 and to reduce congestion.

		 i dont know where the cars will go, maybe disappear

		 the limited resource here is the parking lot with only 5 spaces
		 the 5 spaces is represented by the buffered channel, which only allows
		 5 items inside it for a period of time X. X here is 3 seconds
*/

func main() {
	fmt.Println("b file")
	wg := sync.WaitGroup{}

	// this is a buffered channle that has 5 spots
	parkingLot := make(chan struct{}, 5)

	// so today there are 10 cars trying to park and you are supposed to manage
	// the parking lot no car should stay for more than 3 seconds in parking lot
	// i know its unrealistic but stay with me now :).
	defer close(parkingLot)
	howManyParked := 0
	howManyTimeout := 0
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(carNumber int) {
			defer wg.Done()
			select {
			case parkingLot <- struct{}{}:
				fmt.Printf("car %d parked\n", carNumber)
				time.Sleep(3 * time.Second) //park only for 3 seconds
				<-parkingLot                // leave the spot after 3 seconds
				howManyParked++
				fmt.Printf("car %d left..\n", carNumber)
			case <-time.After(3 * time.Second):
				fmt.Printf("car %d left because of timeout\n", carNumber)
				howManyTimeout++
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("%d parked\n%d timeout\n", howManyParked, howManyTimeout)
}
