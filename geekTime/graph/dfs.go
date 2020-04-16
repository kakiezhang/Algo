/**
深度优先算法
通过穷举来实现，即在一条路走不通的时候退回到
上一个路口，然后换一条路继续走
并不一定是最优解，寻路的方式跟顶点之间建立的边的先后顺序有关
*/
package graph

import "fmt"

type DepthFirstSearch struct {
	graph   *Graph
	road    []int  // 经过该顶点之前所到达的顶点
	visited []bool // 顶点是否被访问过
	found   bool   // 是否找到
}

func NewDFS(g *Graph) *DepthFirstSearch {
	return &DepthFirstSearch{
		graph:   g,
		road:    make([]int, g.max),
		visited: make([]bool, g.max),
		found:   false,
	}
}

func (dfs *DepthFirstSearch) Find(
	s, t int) {
	if dfs.found { // 找到屏蔽其他分支
		return
	}

	for {
		vertex := dfs.graph.arr[s].Poll() // 相邻的顶点
		if vertex == nil {
			break
		}

		v := vertex.GetData().(*Vertex).data.(int)
		if v == t {
			dfs.found = true
		}

		if dfs.visited[v] {
			continue
		}

		fmt.Printf("i=> s: %d, v: %d, t: %d, found: %t \n", s, v, t, dfs.found)

		dfs.Find(v, t) // 从相邻的顶点继续往终点找

		fmt.Printf("o=> s: %d, v: %d, t: %d, found: %t \n", s, v, t, dfs.found)
		dfs.road[v] = s
	}
}

func (dfs *DepthFirstSearch) Print(s, t int) {
	if s != dfs.road[t] {
		dfs.Print(s, dfs.road[t])
	} else {
		fmt.Printf("%d ", s)
	}

	fmt.Printf("%d ", t)
}
