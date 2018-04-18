package graph

import "fmt"

//BFS ...
type BFS struct {
	queue          chan interface{}
	edges          [][]int
	visited        map[int]bool
	traversedEdges chan interface{}
}

//NewBFS ...
func NewBFS(queueSize int) *BFS {
	return &BFS{
		queue:          make(chan interface{}, queueSize),
		traversedEdges: make(chan interface{}),
		edges:          make([][]int, 5),
		visited:        make(map[int]bool),
	}
}

//AddEdge ...
func (bfs *BFS) AddEdge(a, b int) {
	bfs.growIfNeeded(a)
	bfs.edges[a] = append(bfs.edges[a], b)
}

func (bfs *BFS) growIfNeeded(a int) {
	if len(bfs.edges) == cap(bfs.edges) {
		bfs.edges = append(bfs.edges[:cap(bfs.edges)], []int{})
	}
}

//Run ...
func (bfs *BFS) Run() <-chan interface{} {
	go func() {
		defer close(bfs.traversedEdges)
		for poppedValue := range bfs.queue {
			if val, ok := poppedValue.(int); ok {
				if !bfs.visited[val] {
					bfs.visited[val] = true
					for _, edge := range bfs.edges[val] {
						bfs.traversedEdges <- fmt.Sprintf("(%d,%d)", val, edge)
						bfs.queue <- edge
					}
				}
			}
		}
	}()
	for k, edge := range bfs.edges {
		if len(edge) > 0 {
			bfs.queue <- k
			break
		}
	}
	return bfs.traversedEdges
}

//Stop ...
func (bfs *BFS) Stop() {
	close(bfs.queue)
}
