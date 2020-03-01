/**
数组实现顺序栈
*/
package main

import "fmt"

const MAXNUM = 2

func main() {
	sa := newStackA(MAXNUM)
	sa.push(3)
	fmt.Println(sa)
	sa.push(4)
	fmt.Println(sa)
	// sa.push(5)
	// fmt.Println(sa)

	var x int
	x = sa.pop()
	// x = sa.pop()
	// x = sa.pop()
	fmt.Println(x)
	fmt.Println(sa)
}

func newStackA(num int) *stackA {
	return &stackA{
		num:  num,
		tail: -1,
		eles: make([]int, num),
	}
}

type stackA struct {
	num  int
	tail int
	eles []int
}

func (sa *stackA) String() string {
	var s string
	for i := 0; i <= sa.tail; i++ {
		s += fmt.Sprintf("%d,", sa.eles[i])
	}
	return s
}

func (sa *stackA) push(x int) {
	if sa.tail == sa.num-1 {
		panic("stack is full")
	}

	sa.tail += 1
	sa.eles[sa.tail] = x
}

func (sa *stackA) pop() int {
	if sa.tail == -1 {
		panic("stack is empty")
	}

	x := sa.eles[sa.tail]
	sa.tail -= 1
	return x
}
