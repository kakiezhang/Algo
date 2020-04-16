/**
链表实现链式栈
*/
package main

import "fmt"

const MAXNUM = 4

func main() {
	sl := newStackL(MAXNUM)
	sl.push(3)
	// fmt.Println(sl)
	sl.push(4)
	// fmt.Println(sl)
	sl.push(5)
	sl.push(6)
	// sl.push(7)
	// sl.push(8)
	fmt.Println(sl)

	var x int
	x = sl.pop()
	x = sl.pop()
	x = sl.pop()
	x = sl.pop()
	fmt.Println(x)
	fmt.Println(sl)
}

func newStackL(num int) *stackL {
	return &stackL{
		num: num,
		cnt: 0,
		head: &Node{
			next: nil,
		},
	}
}

type stackL struct {
	num  int
	cnt  int
	head *Node // 保持在栈顶部
}

func (sl *stackL) String() string {
	var s string
	var ptrNode = sl.head.next
	s = fmt.Sprintf("head:[%p][%d][%p]->", sl.head, sl.head.data, ptrNode)
	for i := 0; i < sl.cnt; i++ {
		s += fmt.Sprintf("[%p][%d][%p]->", ptrNode, ptrNode.data, ptrNode.next)
		ptrNode = ptrNode.next
	}
	return s
}

type Node struct {
	data int
	next *Node
}

func (sl *stackL) push(x int) {
	if sl.cnt == sl.num {
		panic("stack is full")
	}

	newNode := &Node{
		data: x,
		next: sl.head.next,
	}

	sl.head.next = newNode

	sl.cnt += 1
}

func (sl *stackL) pop() int {
	if sl.cnt == 0 {
		panic("stack is empty")
	}

	x := sl.head.next.data
	sl.head.next = sl.head.next.next

	sl.cnt -= 1

	return x
}
