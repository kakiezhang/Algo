/**
在一个类似杨辉的三角结构中，
从顶到底求经过的最短路径和
行/列：三角的层级
状态节点：初始化时，先将每一层的路径节点
映射到states上，然后每一层直接根据上一层
状态累加计算
*/
package dynamicProgramming

import "math"

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

func (yh *Yanghui) FindMinPath() int {
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

	var min = math.MaxInt64
	for i := 0; i < yh.level; i++ {
		if min > yh.states[yh.level-1][i] {
			min = yh.states[yh.level-1][i]
		}
	}

	return min
}
