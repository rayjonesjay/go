package main

/* using queues in graph algorithms
here's how a queue is used in breadth first search


*/

import (
	"fmt"
)

// graph represents an adjacency list graph
type Graph struct {
	adjList map[int][]int
}

func (g *Graph) AddEdge(from, to int) {
	g.adjList[from] = append(g.adjList[from], to)
	// for undirected graph where there are no directed paths you can add a link to both nodes
	g.adjList[to] = append(g.adjList[to], from)
}


/* what does BFS do?
First we need two data structures to help us
1. map[int]bool -> this is going to store the node name/number/identifier and its job is to tell us
	if that node has been visited or not. If it has not been visited then we set it true then visit its 
	neighbours if it has any.

2. []int{} -> a queue -> an array of integers to store the nodes that are to be visited.
	- initialize it with start node
*/
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	q := []int{start} // start node to indicate our start position

	// while the length of queue is greather than 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		// check if the node has been visited
		if !visited[node] {
			fmt.Println(node)
			visited[node] = true

			for _, neighbour := range g.adjList[node] {
				if !visited[neighbour] {
					q = append(q, neighbour)
				}
			}
		}
	}

}

func main() {
	fmt.Println("world")
	g := &Graph{adjList: make(map[int][]int)}
	g.AddEdge(1,2)
	g.AddEdge(1,3)
	g.AddEdge(2,4)
	g.AddEdge(2,5)
	g.AddEdge(3,6)
	g.BFS(1)
}
