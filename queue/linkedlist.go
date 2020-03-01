/**
用链表实现链式队列
*/
package main

import "fmt"

func main() {
	ql := newQueueL(5)
	fmt.Printf("ql: %v\n", ql)

	for _, v := range []string{
		"a", "b", "c", "d", "e"} {
		ql.push(v)
		fmt.Printf("ql: %v\n", ql)
	}

	// ql.push("f")
	// fmt.Printf("ql: %v\n", ql)

	var x string
	x = ql.pop()
	x = ql.pop()
	x = ql.pop()
	x = ql.pop()
	x = ql.pop()
	fmt.Printf("x: %s\n", x)
	fmt.Printf("ql: %v\n", ql)
	// x = ql.pop()
}

func newQueueL(num int) *QueueL {
	head := &Node{
		prev: nil,
		next: nil,
	}

	tail := &Node{
		prev: head,
		next: nil,
	}

	head.next = tail

	// fmt.Printf("tail: %p\n", tail)
	return &QueueL{
		head: head,
		tail: tail,
		cnt:  0,
		num:  num,
	}
}

type QueueL struct {
	head *Node
	tail *Node
	cnt  int
	num  int
}

func (ql *QueueL) String() string {
	var n = ql.head
	var s = fmt.Sprintf("[head][%p][%v][%p]->", n.prev, n, n.next)
	n = n.next

	for i := 0; i < ql.cnt; i++ {
		s += fmt.Sprintf("[%p][%v][%p]->", n.prev, n, n.next)
		n = n.next
	}
	s += fmt.Sprintf("[tail][%p][%v][%p]", n.prev, n, n.next)

	return s
}

type Node struct {
	data string
	prev *Node
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%p %s", n, n.data)
}

func (ql *QueueL) push(x string) {
	ql.pushNode(&Node{
		data: x,
		prev: nil,
		next: nil,
	})
}

func (ql *QueueL) pushNode(xn *Node) {
	if ql.cnt == ql.num {
		panic("queue has reached maxnum.")
	}

	xn.next = ql.tail
	xn.prev = ql.tail.prev
	ql.tail.prev.next = xn
	ql.tail.prev = xn
	// ql.tail = xn
	ql.cnt += 1
}

func (ql *QueueL) pop() string {
	xn := ql.popNode()
	return xn.data
}

func (ql *QueueL) popNode() *Node {
	if ql.cnt == 0 {
		panic("queue no elements.")
	}

	ql.cnt -= 1
	ele := ql.head.next
	ql.head.next.next.prev = ql.head
	ql.head.next = ql.head.next.next
	return ele
}
