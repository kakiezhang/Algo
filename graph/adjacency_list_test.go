package graph

import (
	"fmt"
	"testing"
)

func TestOneWayWeightless(t *testing.T) {
	g := NewGraph(3)
	g.addOneEdge(1, 2, 0)
	fmt.Println("One-Way Weightless: ")
	fmt.Println(g)
}

func TestTwoWayWeightless(t *testing.T) {
	g := NewGraph(3)
	g.addTwoEdge(1, 2, 0)
	fmt.Println("Two-Way Weightless: ")
	fmt.Println(g)
}

func TestOneWayWeighted(t *testing.T) {
	g := NewGraph(5)
	g.addOneEdge(1, 2, 3)
	g.addOneEdge(2, 3, 2)
	g.addOneEdge(2, 4, 1)
	g.addOneEdge(3, 4, 4)
	fmt.Println("One-Way Weighted: ")
	fmt.Println(g)
}

func TestTwoWayWeighted(t *testing.T) {
	g := NewGraph(5)
	g.addTwoEdge(1, 2, 3)
	g.addTwoEdge(2, 3, 2)
	g.addTwoEdge(3, 4, 4)
	g.addTwoEdge(4, 1, 1)
	fmt.Println("Two-Way Weighted: ")
	fmt.Println(g)

	for _, v := range [][2]int{
		[2]int{1, 4},
		[2]int{1, 3},
	} {
		fmt.Printf("findEdge between [%d, %d]?: %t\n",
			v[0], v[1], g.findEdge(v[0], v[1]))
	}
}
