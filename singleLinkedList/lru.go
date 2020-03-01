/**
目标：用单链表实现LRU(Least Recently Used)
思路：给定一个需要访问的元素值和一个固定结点数量的单链表，
遍历链表，去找是否有匹配的值
假如找到有的话，就把元素从当前位置删除，并添加到头结点；
假如没有找到有匹配的元素，那么先要判断当前的整个链表长度
是否已经达到给定的最大值
如果达到了，那么需要把尾结点删掉，然后在元素插入头节点；
如果没达到，那直接把元素插入头节点；
*/
package main

import "fmt"

type Sentinel struct {
}

type Node struct {
	data int
	next *Node
}

type LRUSingle struct {
	head    *Node // 这个头节点是个哨兵（只有 next 指针值，没有 data）
	nodeCnt int
	length  int
}

func (lru *LRUSingle) String() string {
	var rs = fmt.Sprintf("[%p %d][%p]", lru.head, lru.head.data, lru.head.next)
	if lru.head.next == nil {
		return rs
	}

	var ptrNode = lru.head.next
	rs += "->"
	for {
		rs += fmt.Sprintf("[%p %d][%p]",
			ptrNode, ptrNode.data, ptrNode.next)
		if ptrNode.next == nil {
			break
		}
		rs += "->"
		ptrNode = ptrNode.next
	}
	return rs
}

func (lru *LRUSingle) findLastByValue(v int) (rs *Node) {
	if lru.head.next == nil {
		return
	}

	var lastNode = lru.head
	var ptrNode = lru.head.next
	for {
		if ptrNode.data == v {
			rs = lastNode
			break
		}
		if ptrNode.next == nil {
			break
		}
		lastNode = ptrNode
		ptrNode = ptrNode.next
	}
	return rs
}

func (lru *LRUSingle) deleteNodeByLast(lastNode *Node) {
	if lastNode.next == nil {
		return
	}
	lastNode.next = lastNode.next.next
	lru.nodeCnt -= 1
}

func (lru *LRUSingle) deleteTailNode() {
	if lru.head.next == nil {
		return
	}

	var lastNode = lru.head
	var ptrNode = lru.head.next
	for {
		if ptrNode.next == nil {
			lru.deleteNodeByLast(lastNode)
			break
		}
		lastNode = ptrNode
		ptrNode = ptrNode.next
	}
}

func (lru *LRUSingle) insertValueToHead(v int) {
	newNode := &Node{data: v}
	lru.insertNodeToHead(newNode)
}

func (lru *LRUSingle) insertNodeToHead(newNode *Node) {
	newNode.next = lru.head.next
	originCnt := lru.nodeCnt
	lru.head.next = newNode
	lru.nodeCnt = originCnt + 1
}

func (lru *LRUSingle) set(v int) {
	lastNode := lru.findLastByValue(v)
	// fmt.Printf("lastNode: [%v]\n", lastNode)
	if lastNode != nil {
		lru.deleteNodeByLast(lastNode)
		lru.insertValueToHead(v)
	} else {
		if lru.nodeCnt == lru.length {
			lru.deleteTailNode()
			lru.insertValueToHead(v)
		} else if lru.nodeCnt < lru.length {
			lru.insertValueToHead(v)
		} else {
			panic("LRUSingle length error")
		}
	}
}

func main() {
	length := 5
	lru := &LRUSingle{
		head:   &Node{}, // 初始化哨兵
		length: length,
	}
	fmt.Printf("init lru: %s\n", lru)

	var x int

	for x = 15; x <= 18; x++ {
		fmt.Printf("x: %d\n", x)
		lru.set(x)
		fmt.Printf("lru: %s\n", lru)
	}

	x = 18
	fmt.Printf("x: %d\n", x)
	lru.set(x)
	fmt.Printf("lru: %s\n", lru)

	x = 16
	fmt.Printf("x: %d\n", x)
	lru.set(x)
	fmt.Printf("lru: %s\n", lru)

	x = 10
	fmt.Printf("x: %d\n", x)
	lru.set(x)
	fmt.Printf("lru: %s\n", lru)

	for x = 3; x <= 5; x++ {
		fmt.Printf("x: %d\n", x)
		lru.set(x)
		fmt.Printf("lru: %s\n", lru)
	}
}
