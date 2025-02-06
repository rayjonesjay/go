package main

/*
GRAPH REPRESENTATION
Graphs can be stored in two main ways.
1. adjacency matrix:
	- a 2d array where:
		- matrix[i][j] = 1 if there is an edge from i to j
		- mtraix[i][j] could also store weights for weighted graphs

2. adjacency list:
	- a list where each node stores a list of its neighbours
	- pros: space-efficient for sparse graphs
	- cons: slower for edge lookup compared to matrices

What is a queue? a QUEUE is a first in first out data structure
- the first itemp added to the queue is the first one removed
- common operation:
	- enqueue -> add an item to the back of queue -
	- dequeue -> remove an item from the queue  - the first element in the list

What is a dequeue? also known as double ended queue, is a data structure that allows adding or removing items from both ends.
- common operations:
	- PushFront: Add an item to the front.
	- PushBack: Add an item to the back.
	- PopFront: Remove an item from the front.
	- PopBack: Remove an item from back.


Below is an implementation of a queue in go
Go does not have a built in queue, but you can use a slices or implement one using a struct.

*/
import (
	"fmt"
)

type Queue struct {
	items []int
}

// Enqueue adds value to end of *Queue.items
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (int, bool) {
	// check if there is something to remove
	if len(q.items) == 0 {
		return 0, false
	}

	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {

	fmt.Println("hello world")
	q := &Queue{}
	q.Enqueue(10)
	q.Enqueue(29)
	q.Enqueue(40)

	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		fmt.Println(val)
	}
}
