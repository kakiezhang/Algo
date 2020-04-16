/**
一个nxn矩阵的最短路径
*/
package backTrack

import (
	"fmt"
	"testing"
)

func TestMatrixMin(t *testing.T) {
	var arr [][]int
	arr = [][]int{
		[]int{1, 1, 3},
		[]int{2, 2, 4},
		[]int{2, 3, 1},
	}

	// arr = [][]int{
	// 	[]int{1, 1},
	// 	[]int{2, 2},
	// }

	m := newMatrix(arr, len(arr))

	m.findMin(0, 0, 0)

	fmt.Println(m.min)
}
