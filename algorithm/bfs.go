package main

import "fmt"

/*
Breadth first search
Traversal means visiting all the nodes of a graph. Breadth First Traversal or
Breadth First Search is a recursive algorithm for searching all the vertices of a graph or tree data structure.

BFS algorithm
A standard BFS implementation puts each vertex of the graph into one of two categories:

Visited
Not Visited
The purpose of the algorithm is to mark each vertex as visited while avoiding cycles.

The algorithm works as follows:

Start by putting any one of the graph's vertices at the back of a queue.
Take the front item of the queue and add it to the visited list.
Create a list of that vertex's adjacent nodes. Add the ones which aren't in the visited list to the back of the queue.
Keep repeating steps 2 and 3 until the queue is empty.
The graph might have two different disconnected parts so to make sure that we cover every vertex, we can also run the BFS algorithm on every node
*/

func bfs(graph map[int][]int, vertex int) []int {
	var queue []int
	visited := make(map[int]bool)
	visited[vertex] = true
	queue = append(queue, vertex)
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		for _, neighbour := range graph[next] {
			if !visited[neighbour] {
				queue = append(queue, neighbour)
				visited[neighbour] = true
			}
		}

	}

	data := []int{}
	for key, _ := range visited {
		data = append(data, key)
	}
	return data
}

func main() {
	graph := map[int][]int{
		5: {3, 7},
		3: {2, 4},
		7: {8},
		2: {},
		4: {8},
		8: {},
	}
	found := bfs(graph, 5)
	fmt.Println(found)
}
