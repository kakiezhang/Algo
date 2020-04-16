/**
一个nxn矩阵的最短路径
*/
package backTrack

import (
	"math"
)

type Matrix struct {
	arr [][]int
	cnt int // 矩阵的行/列
	min int
}

func newMatrix(arr [][]int, cnt int) *Matrix {
	return &Matrix{
		arr: arr,
		cnt: cnt,
		min: math.MaxInt64,
	}
}

func (m *Matrix) findMin(i, j, sum int) {
	if i >= m.cnt || j >= m.cnt {
		return
	}

	// fmt.Printf("i: %d, j: %d, sum: %d\n", i, j, sum)

	if i == m.cnt-1 && j == m.cnt-1 {
		if m.min > sum+m.arr[i][j] {
			m.min = sum + m.arr[i][j]
		}
		return
	}

	// 都加上当前路径的值再继续
	m.findMin(i, j+1, sum+m.arr[i][j]) // 选择右边的继续
	m.findMin(i+1, j, sum+m.arr[i][j]) // 选择下边的继续
}
