/**
深度优先遍历
通过邻接表构建逆邻接表
即s->t的关系是，s依赖于t，
t优先于s执行
*/
package topology

import (
	"fmt"

	"github.com/kakiezhang/Algo/geekTime/graph"
)

func topoByDfs(g *graph.Graph) {
	invert_g := graph.NewGraph(g.Max) // 逆邻接表

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
			invert_g.AddOneEdge(k, i, 0)
		}
	}

	visited := make([]bool, g.Max)

	for i := 1; i < g.Max; i++ {
		if !visited[i] {
			visited[i] = true
			printDfs(invert_g, i, visited)
		}
	}
}

func printDfs(invert_g *graph.Graph, i int, visited []bool) {
	invert_l := invert_g.Vtx[i]
	if invert_l == nil {
		fmt.Printf("-> %d ", i)
		return
	}

	for j := 0; j < invert_l.Size(); j++ {
		node := invert_l.Get(j)
		if node == nil {
			continue
		}

		k := node.GetData().(*graph.Vertex).GetData().(int)
		if visited[k] {
			continue
		}

		visited[k] = true
		printDfs(invert_g, k, visited)
	}

	fmt.Printf("-> %d ", i)
}
