/**
数组实现循环队列
循环队列相较于顺序队列的优势在于当 tail 在数组尾端，
数组又没有满的时候，不需要进行数据平移，在时间复杂
度上减少了一次最坏情况的 O(N) 平移操作
但是，循环队列有个缺陷，就是在构造圆盘的时候，需要
比计算的数组多出一个点，因为这个点是用来判断队空和
队满情况而设置的
*/
package main

import "fmt"

func main() {
	// t0()
	t1()
	// t7()
}

func t0() {
	qc := newQueueC(0)
	fmt.Printf("qc: %v\n", qc)

	// qc.push("a")
	// fmt.Printf("qc: %v\n", qc)

	var x string
	x = qc.pop()
	fmt.Printf("x: %s\n", x)
	fmt.Printf("qc: %v\n", qc)
}

func t1() {
	qc := newQueueC(1)
	fmt.Printf("qc: %v\n", qc)

	qc.push("a")
	fmt.Printf("qc: %v\n", qc)

	var x string
	x = qc.pop()
	fmt.Printf("x: %s\n", x)
	fmt.Printf("qc: %v\n", qc)

	qc.push("b")
	fmt.Printf("qc: %v\n", qc)
}

func t7() {
	qc := newQueueC(7)
	fmt.Printf("qc: %v\n", qc)

	for _, v := range []string{
		"a", "b", "c", "d", "e", "f"} {
		qc.push(v)
	}
	fmt.Printf("qc: %v\n", qc)

	qc.push("g")
	fmt.Printf("qc: %v\n", qc)

	// qc.push("h")
	// fmt.Printf("qc: %v\n", qc)

	var x string
	x = qc.pop()
	x = qc.pop()
	fmt.Printf("x: %s\n", x)
	fmt.Printf("qc: %v\n", qc)

	qc.push("a")
	fmt.Printf("qc: %v\n", qc)
	qc.push("b")
	fmt.Printf("qc: %v\n", qc)

	for i := 0; i < 7; i++ {
		qc.pop()
	}
	fmt.Printf("qc: %v\n", qc)

	for _, v := range []string{
		"a", "b", "c", "d", "e", "f", "g"} {
		qc.push(v)
	}
	fmt.Printf("qc: %v\n", qc)
}

func newQueueC(cnt int) *QueueC {
	if cnt <= 0 {
		panic("elements cnt should be gt 0 ")
	}

	return &QueueC{
		head: 0,
		tail: cnt,
		num:  cnt + 1,
		eles: make([]string, cnt+1),
	}
}

type QueueC struct {
	head int
	tail int
	num  int // 圆盘上的分块数量
	eles []string
}

func (qc *QueueC) String() string {
	var s string
	for i := 0; i < qc.num; i++ {
		s += fmt.Sprintf("%s, ", qc.eles[i])
	}
	return s
}

func (qc *QueueC) push(x string) {
	if (qc.tail+2)%qc.num == qc.head {
		panic("queue has reached maxnum.")
	}

	qc.tail = (qc.tail + 1) % qc.num
	qc.eles[qc.tail] = x
}

func (qc *QueueC) pop() string {
	if (qc.tail+1)%qc.num == qc.head {
		panic("queue no elements.")
	}

	ele := qc.eles[qc.head]
	qc.eles[qc.head] = ""
	qc.head = (qc.head + 1) % qc.num
	return ele
}
