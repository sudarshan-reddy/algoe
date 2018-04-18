package main

import (
	"fmt"

	"github.com/sudarshan-reddy/algoe/graph"
)

func main() {
	bfs := graph.NewBFS(10)
	bfs.AddEdge(1, 2)
	bfs.AddEdge(2, 5)
	bfs.AddEdge(1, 3)
	bfs.AddEdge(3, 4)
	bfs.AddEdge(3, 5)

	op := bfs.Run()
	for val := range op {
		fmt.Println(val)
	}
}
