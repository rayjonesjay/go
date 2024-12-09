package main

import "fmt"

/*
Think of a graph as a city with buildings(nodes) and roads(vertices) connecting them together.
vertices - nodes
edges - path connecting the nodes
weight(if present) - represents the distances,travel time,cost to travel between locations

the goal of graph algorithms is often to find the best way to travel between locations, explore
all possible connections, or identify patterns in the map.

A graph G is represented as
G = (V,E) where:
	- V a set of vertices (nodes)
	- E a set of edges (connections between nodes)



Two types of directed graphs
Weighted - edges have limits (eg, distances)
Unweighted - all edges are treated equally
Cyclic graph - contains at least one cycle (a path where you can return to the starting node)
Acyclic - no cycles exist
Connected - there is a path between every pair of node
*/

func main() {
	fmt.Println("hello world")
}
