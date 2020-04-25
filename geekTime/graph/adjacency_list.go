/**
邻接表
单向双向图，有权无权图（交叉4种情况）
储存顶点的数组内从第1位开始
*/
package graph

import (
	"fmt"

	"github.com/kakiezhang/Algo/geekTime/linkedlist"
)

type Graph struct {
	Vtx []*linkedlist.DoublyLinkedList // 存储顶点的数组
	Max int
}

type Vertex struct {
	data   interface{}
	weight int
}

func (g *Graph) String() string {
	var rs string
	for k, v := range g.Vtx {
		rs += fmt.Sprintf("[%d] => %v\n", k, v)
	}
	return rs
}

func (v *Vertex) String() string {
	return fmt.Sprintf("[v:%v w:%d]", v.data, v.weight)
}

func NewGraph(max int) *Graph {
	return &Graph{
		Vtx: make([]*linkedlist.DoublyLinkedList, max+1),
		Max: max + 1,
	}
}

func newVertex(v, w int) *Vertex {
	return &Vertex{
		data:   v,
		weight: w,
	}
}

func (g *Graph) addEdge(s int, v *Vertex) {
	// 添加一条从s指向v的边
	if s <= 0 {
		panic("vertex index has to be gt zero")
	}
	if s >= g.Max+1 {
		panic("vertex index has to be le max")
	}

	if g.Vtx[s] == nil {
		g.Vtx[s] = linkedlist.NewDoublyLinkedList()
	}

	g.Vtx[s].Insert(v)
}

func (g *Graph) addOneEdge(s, t, w int) {
	// 添加单向的边，s指向t
	g.addEdge(s, newVertex(t, w))
}

func (g *Graph) addTwoEdge(s, t, w int) {
	// 添加双向的边，s指向t，t指向s
	g.addEdge(s, newVertex(t, w))
	g.addEdge(t, newVertex(s, w))
}

func (g *Graph) findEdge(s int, t interface{}) bool {
	// 找s和t之间是否存在一条s指向t边
	f := func(v interface{}) interface{} {
		return v.(*Vertex).data
	}

	if g.Vtx[s].FindNode(t, f) != nil {
		return true
	} else {
		return false
	}
}
