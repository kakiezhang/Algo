package graph

import (
	"fmt"
	"testing"
)

func TestMatrixOneWayWeightless(t *testing.T) {
	mg := newMatrixGraph(6)
	mg.addOneEdge(1, 2, 1)
	mg.addOneEdge(6, 2, 1)
	mg.addOneEdge(6, 1, 1)
	fmt.Println("One-way weightless: ")
	fmt.Println(mg)
}

func TestMatrixTwoWayWeightless(t *testing.T) {
	mg := newMatrixGraph(6)
	mg.addTwoEdge(1, 2, 1)
	mg.addTwoEdge(6, 2, 1)
	mg.addTwoEdge(6, 1, 1)
	fmt.Println("Two-way weightless: ")
	fmt.Println(mg)
}

func TestMatrixOneWayWeighted(t *testing.T) {
	mg := newMatrixGraph(6)
	mg.addOneEdge(1, 2, 9)
	mg.addOneEdge(6, 2, 7)
	mg.addOneEdge(6, 1, 5)
	fmt.Println("One-way Weighted: ")
	fmt.Println(mg)
}

func TestMatrixTwoWayWeighted(t *testing.T) {
	mg := newMatrixGraph(6)
	mg.addTwoEdge(1, 2, 8)
	mg.addTwoEdge(6, 2, 5)
	mg.addTwoEdge(6, 1, 2)
	fmt.Println("Two-way weighted: ")
	fmt.Println(mg)

	for _, v := range [][2]int{
		[2]int{1, 4},
		[2]int{1, 3},
		[2]int{1, 2},
		[2]int{1, 6},
	} {
		fmt.Printf("findEdge between [%d, %d]?: %d\n",
			v[0], v[1], mg.findEdge(v[0], v[1]))
	}
}
