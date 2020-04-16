/**
邻接矩阵
单向双向图，有权无权图（交叉4种情况）
先把顶点都映射到矩阵上面
顶点（从1开始到max结束）的坐标是当前顶点的数值-1
*/
package graph

import "fmt"

type MatrixGraph struct {
	arr [][]int
	max int // 顶点数量
}

func (mg *MatrixGraph) String() string {
	rs := fmt.Sprintln("[")
	for i := 0; i < mg.max; i++ {
		var line string
		for j := 0; j < mg.max; j++ {
			line += fmt.Sprintf("%d, ", mg.arr[i][j])
		}
		rs += fmt.Sprintln(line)
	}
	rs += "]"
	return rs
}

func newMatrixGraph(max int) *MatrixGraph {
	mg := &MatrixGraph{
		arr: make([][]int, max),
		max: max,
	}

	for i := 0; i < max; i++ {
		mg.arr[i] = make([]int, max)
	}

	return mg
}

func (mg *MatrixGraph) addEdge(s, t, w int) {
	// 添加一条s指向t的边，w为权重，没有则为1
	if s <= 0 || t <= 0 {
		panic("vertex num has to be gt zero")
	}
	if s > mg.max || t > mg.max {
		panic("vertex num has to be lt max")
	}

	mg.arr[s-1][t-1] = w
}

func (mg *MatrixGraph) addOneEdge(s, t, w int) {
	mg.addEdge(s, t, w)
}

func (mg *MatrixGraph) addTwoEdge(s, t, w int) {
	mg.addEdge(s, t, w)
	mg.addEdge(t, s, w)
}

func (mg *MatrixGraph) findEdge(s, t int) int {
	return mg.arr[s-1][t-1]
}
