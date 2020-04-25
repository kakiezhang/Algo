package topology

import (
	"fmt"
	"testing"

	"github.com/kakiezhang/Algo/geekTime/graph"
)

func TestTopoByDfs(t *testing.T) {
	var g *graph.Graph

	// 1 2 3 5 4
	g = graph.NewGraph(5)
	g.AddOneEdge(1, 2, 0)
	g.AddOneEdge(1, 3, 0)
	// g.AddOneEdge(2, 6, 0)
	g.AddOneEdge(2, 5, 0)
	g.AddOneEdge(5, 4, 0)
	// g.AddOneEdge(4, 5, 0)
	topoByDfs(g)

	fmt.Println()

	// 6 3 1 2 3 4 5
	g = graph.NewGraph(6)
	g.AddOneEdge(6, 3, 0)
	g.AddOneEdge(1, 2, 0)
	g.AddOneEdge(3, 4, 0)
	g.AddOneEdge(3, 5, 0)
	topoByDfs(g)
}
