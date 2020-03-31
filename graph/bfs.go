/**
广度优先
寻找三度关系好友
*/
package graph

import (
	"github.com/kakiezhang/Algo/linkedlist"
)

type BreathFirstSearch struct {
	graph   *Graph                       // 需要处理的图
	queue   *linkedlist.DoublyLinkedList // 挨个处理的顶点队列
	road    []int                        // 到达顶点的上一个经过的顶点
	visited []bool                       // 已经访问过的顶点
}

func NewBfs(g *Graph) *BreathFirstSearch {
	return &BreathFirstSearch{
		graph:   g,
		queue:   linkedlist.NewDoublyLinkedList(),
		road:    make([]int, g.max),
		visited: make([]bool, g.max),
	}
}

func (bfs *BreathFirstSearch) Find(
	s, t int) {
	// 找一条s到t的最短路径
	bfs.queue.Insert(s)

	for {
		vert := bfs.queue.Poll()
		if vert == nil {
			break
		}

		v := vert.GetData().(int)

		for {
			next := bfs.graph.arr[v].Poll() // 相邻的下一个顶点
			if next == nil {
				break
			}

			m := next.GetData().(*Vertex).data.(int) // 相邻的下一个顶点的值

			if !bfs.visited[m] {
				bfs.visited[m] = true
				bfs.road[m] = v
				bfs.queue.Insert(m)
			}

			if m == t {
				break
			}
		}
	}
}
