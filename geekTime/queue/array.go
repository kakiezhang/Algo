/**
数组实现顺序队列
*/
package main

import "fmt"

func main() {
	qa := newQueueA(6)
	for _, v := range []string{
		"a", "b", "c", "d", "e", "f"} {
		qa.push(v)
	}
	fmt.Println(qa)

	var x string
	x = qa.pop()
	fmt.Println(x)
	x = qa.pop()
	fmt.Println(x)
	x = qa.pop()
	fmt.Println(x)
	fmt.Println(qa)

	qa.push("a")
	qa.push("b")
	qa.push("c")
	fmt.Println(qa)
	// qa.push("d")
	// fmt.Println(qa)

	for i := 0; i < 6; i++ {
		qa.pop()
	}
	qa.push("g")
	fmt.Println(qa)
}

func newQueueA(num int) *QueueA {
	return &QueueA{
		head: -1,
		tail: -1,
		num:  num,
		eles: make([]string, num),
	}
}

type QueueA struct {
	head int
	tail int
	num  int
	eles []string
}

func (qa *QueueA) String() string {
	var s = "eles: "
	for i := 0; i < qa.num; i++ {
		s += fmt.Sprintf("%s, ", qa.eles[i])
	}
	return s
}

func (qa *QueueA) push(x string) {
	if qa.tail-qa.head == qa.num {
		panic("queue length reach the max.")
	}

	if qa.tail == qa.num-1 { // 平移
		var i int
		for i = 0; i < qa.tail-qa.head; i++ {
			qa.eles[i] = qa.eles[i+qa.head+1]
		}

		for j := qa.tail - qa.head; j < qa.num; j++ {
			qa.eles[j] = ""
		}

		qa.head = -1
		qa.tail = i - 1
	}

	qa.tail += 1
	qa.eles[qa.tail] = x
}

func (qa *QueueA) pop() string {
	if qa.tail-qa.head == 0 {
		panic("queue no elements.")
	}

	qa.head += 1
	ele := qa.eles[qa.head]
	qa.eles[qa.head] = ""
	return ele
}
