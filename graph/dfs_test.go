package graph

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	// 2 - 1 - 3
	// |       |
	// 4   5 - 6
	// | /   \ |
	// 8   7 - 9

	g := NewGraph(9)

	for _, v := range [][2]int{
		[2]int{2, 1},
		[2]int{2, 4},
		[2]int{8, 4},
		[2]int{8, 5},
		[2]int{1, 3},
		// [2]int{1, 5},
		[2]int{3, 6},
		// [2]int{4, 5},
		[2]int{5, 6},
		[2]int{5, 9},
		// [2]int{5, 7},
		[2]int{6, 9},
		[2]int{7, 9},
	} {
		g.addTwoEdge(v[0], v[1], 1)
	}

	dfs := NewDFS(g)
	dfs.Find(2, 7)

	fmt.Println(dfs.road)

}
