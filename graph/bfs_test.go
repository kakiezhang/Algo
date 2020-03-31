package graph

import (
	"fmt"
	"testing"
)

func TestBfs(t *testing.T) {
	// 2 - 1 - 3
	// |   |   |
	// 4 - 5 - 6
	//     |   |
	//     7 - 9

	g := NewGraph(9)
	g.addTwoEdge(2, 1, 1)
	g.addTwoEdge(2, 4, 1)
	g.addTwoEdge(1, 3, 1)
	g.addTwoEdge(1, 5, 1)
	g.addTwoEdge(3, 6, 1)
	g.addTwoEdge(4, 5, 1)
	g.addTwoEdge(5, 6, 1)
	g.addTwoEdge(5, 7, 1)
	g.addTwoEdge(6, 9, 1)
	g.addTwoEdge(7, 9, 1)
	fmt.Println(g)

	bfs := NewBfs(g)
	bfs.Find(2, 7)

	fmt.Println(bfs.road)
}
