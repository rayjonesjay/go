package main

import (
	"fmt"
)

type G struct {
	adjList map[int][]int
}

func (g *G) AddEdge(from,to int){
	g.adjList[from] = append(g.adjList[from],to)
	g.adjList[to] = append(g.adjList[to], from)	
}

func (g *G) BFS(start int) {
	seen := make(map[int]bool)
	q := []int{start}
	
	// while the length of q is not 0
	for len(q) > 0 {
			node := q[0] // get the current node infront of the queue
			q = q[1:] // remove it from the queue

			if !seen[node] { // check if we have explored the node
				 // if not explored
				 seen[node] = true
				 fmt.Println(node)
			
				// check if its neighbours have been expored
				for _, n := range g.adjList[node] {
					q  = append(q,n)	
				}
			}
	}	
}
func main(){
	g := &G{adjList: make(map[int][]int)}
	g.AddEdge(1,2)
	g.AddEdge(1,3)
	g.AddEdge(2,4)
	g.AddEdge(3,5)
	g.AddEdge(1,2)
	g.AddEdge(3,5)
	g.BFS(1)
}

