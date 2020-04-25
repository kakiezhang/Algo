package topology

import (
	"fmt"
	"testing"

	"github.com/kakiezhang/Algo/geekTime/graph"
)

func TestTopoByKahn(t *testing.T) {
	var g *graph.Graph

	g = graph.NewGraph(5)

	g.AddOneEdge(1, 3, 0)
	// g.AddOneEdge(1, 2, 0)
	g.AddOneEdge(2, 5, 0)
	// g.AddOneEdge(4, 5, 0)
	g.AddOneEdge(5, 4, 0)
	fmt.Println(g)

	topoByKahn(g)
}
