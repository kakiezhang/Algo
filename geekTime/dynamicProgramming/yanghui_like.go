/**
在一个类似杨辉的三角结构中，
从顶到底求经过的最短路径&最小和
行/列：三角的层级
状态节点：初始化时，先将每一层的路径节点
映射到states上，然后每一层直接根据上一层
状态累加计算
*/
package dynamicProgramming

import (
	"fmt"
	"math"
)

type Yanghui struct {
	arr    []int
	level  int     // 行/列/层
	states [][]int // 状态节点
}

func newYanghui(arr []int, level int) *Yanghui {
	yh := &Yanghui{
		arr:    arr,
		level:  level,
		states: make([][]int, level),
	}

	var k, m int
	for i := 0; i < yh.level; i++ { // 初始化
		yh.states[i] = make([]int, yh.level)
		for j := 0; j < yh.level; j++ {
			if j <= m {
				yh.states[i][j] = yh.arr[k]
				k += 1
			} else {
				yh.states[i][j] = -1
			}
		}
		m += 1
	}

	return yh
}

func (yh *Yanghui) PrintTriangle() {
	for i := 0; i < yh.level; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", yh.states[i][j])
		}
		fmt.Println()
	}
}

func (yh *Yanghui) FindMinPath() ([]int, []int) {
	for i := 1; i < yh.level; i++ { // 直接从第1行开始
		for j := 0; j <= i; j++ {
			var min = math.MaxInt64
			if j > 0 {
				min = yh.states[i-1][j-1]
			}
			if j < i && min > yh.states[i-1][j] {
				min = yh.states[i-1][j]
			}
			yh.states[i][j] = yh.states[i][j] + min
		}
	}

	var sum = math.MaxInt64 // 最小路径和
	var col int             // 最后一行最小路径所在的列
	for i := 0; i < yh.level; i++ {
		if sum > yh.states[yh.level-1][i] {
			sum = yh.states[yh.level-1][i]
			col = i
		}
	}

	idx := make([]int, yh.level) // 存储每一层的最短节点在states中的索引
	idx[yh.level-1] = col

	path := make([]int, yh.level) // 存储每一层的最短节点在states中的节点值
	path[0] = yh.states[0][0]

	// 最短路径，从倒数第二层开始，把上面的过程倒过来求一遍
	for i := yh.level - 2; i >= 0; i-- {
		var min = math.MaxInt64
		var j int          // 当前行最小路径所在的列
		var col = idx[i+1] // 当前行的下一层最小路径所在的列
		if col > 0 {
			j = idx[i+1] - 1
			min = yh.states[i][j]
		}
		if col <= i && min > yh.states[i][col] {
			j = col
			min = yh.states[i][j]
		}

		idx[i] = j
		path[i+1] = yh.states[i+1][col] - min
	}

	return idx, path
}
