/**
邻接表
单向双向图，有权无权图（交叉4种情况）
*/
package graph

import (
	"fmt"

	"github.com/kakiezhang/Algo/linkedlist"
)

type Graph struct {
	arr []*linkedlist.DoublyLinkedList // 存储顶点的数组
	max int
}

type Vertex struct {
	data   interface{}
	weight int
}

func (g *Graph) String() string {
	var rs string
	for k, v := range g.arr {
		rs += fmt.Sprintf("[%d] => %v\n", k, v)
	}
	return rs
}

func (v *Vertex) String() string {
	return fmt.Sprintf("[v:%v w:%d]", v.data, v.weight)
}

func NewGraph(max int) *Graph {
	return &Graph{
		arr: make([]*linkedlist.DoublyLinkedList, max),
		max: max,
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
	if s >= g.max {
		panic("cannot store more vertex")
	}

	if g.arr[s] == nil {
		g.arr[s] = linkedlist.NewDoublyLinkedList()
	}

	g.arr[s].Insert(v)
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

	if g.arr[s].FindNode(t, f) != nil {
		return true
	} else {
		return false
	}
}
