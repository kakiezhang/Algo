/**
topology by kahn
通过计算顶点的入度来获取先后顺序
如果s依赖于t，关系指向t->s
先找到入度为0的那个顶点
*/
package topology

import (
	"fmt"

	"github.com/kakiezhang/Algo/geekTime/graph"
	"github.com/kakiezhang/Algo/geekTime/linkedlist"
)

func topoByKahn(g *graph.Graph) {
	indegree := make([]int, g.Max)

	for i := 1; i < g.Max; i++ {
		if g.Vtx[i] == nil {
			continue
		}
		for j := 0; j < g.Vtx[i].Size(); j++ {
			node := g.Vtx[i].Get(j)
			if node == nil {
				continue
			}
			k := node.GetData().(*graph.Vertex).GetData().(int)
			indegree[k] += 1
		}
	}

	fmt.Printf("indegree: %v\n", indegree)

	queue := linkedlist.NewDoublyLinkedList()
	for i := 1; i < g.Max; i++ {
		if indegree[i] == 0 {
			queue.Append(i)
		}
	}

	for queue.Size() > 0 {
		node := queue.Poll()
		if node == nil {
			break
		}

		v := node.GetData().(int)
		fmt.Printf("-> %d ", v)

		if g.Vtx[v] == nil {
			continue
		}

		for j := 0; j < g.Vtx[v].Size(); j++ {
			node := g.Vtx[v].Get(j)
			if node == nil {
				continue
			}
			k := node.GetData().(*graph.Vertex).GetData().(int)
			indegree[k]--
			if indegree[k] == 0 {
				queue.Append(k)
			}
		}
	}
}
