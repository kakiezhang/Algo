/*
向单链表中插入第一个元素是不同的
单链表的尾节点没有 next
1. 从单链表的头部插入元素
2. 在单链表中找到某个对应值的元素
3. 在单链表中的某个元素后面插入元素
4. 在单链表中的某个元素前面插入元素
*/
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func (l *SinglyLinkedList) String() string {
	var node = l.head
	var rs string
	for {
		if node.next == nil {
			rs += fmt.Sprintf("[%p][%d][%p]", node, node.data, node.next)
			break
		}
		rs += fmt.Sprintf("[%p][%d][%p]->", node, node.data, node.next)
		node = node.next
	}
	return rs
}

func (l *SinglyLinkedList) findByValue(v int) (rs *Node) {
	var node = l.head
	for {
		if node.data == v {
			rs = node
			break
		}
		if node.next == nil {
			break
		}
		node = node.next
	}
	return rs
}

func (l *SinglyLinkedList) insertValueToHead(v int) {
	var node = &Node{data: v}
	l.insertNodeToHead(node)
}

func (l *SinglyLinkedList) insertNodeToHead(p *Node) {
	// fmt.Println(l.head)
	if l.head == nil {
		l.head = p
	} else {
		p.next = l.head
		l.head = p
	}
	// fmt.Println(l)
}

func (l *SinglyLinkedList) insertValueBefore(oldNode *Node, v int) {
	var newNode = &Node{data: v}
	l.insertNodeBefore(oldNode, newNode)
}

func (l *SinglyLinkedList) insertNodeBefore(oldNode *Node, newNode *Node) {
	var ptrNode = l.head
	var lastNode *Node
	for {
		// fmt.Printf("ptr[%v] old[%v]\n", ptrNode, oldNode)
		if ptrNode == oldNode {
			newNode.next = oldNode
			if lastNode == nil {
				l.head = newNode
			} else {
				lastNode.next = newNode
			}
			break
		}
		if ptrNode.next == nil {
			break
		}
		lastNode = ptrNode
		ptrNode = ptrNode.next
	}
}

func (l *SinglyLinkedList) insertValueAfter(oldNode *Node, v int) {
	var newNode = &Node{data: v}
	l.insertNodeAfter(oldNode, newNode)
}

func (l *SinglyLinkedList) insertNodeAfter(oldNode *Node, newNode *Node) {
	// fmt.Println(l.head)
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = oldNode.next
		oldNode.next = newNode
	}
	// fmt.Println(l)
}

func (l *SinglyLinkedList) findMiddleNodes() (rs []*Node) {
	var slowPtr = l.head
	var quickPtr = l.head
	for {
		if quickPtr.next == nil {
			// 奇数个结点
			rs = append(rs, slowPtr)
			break
		}
		if quickPtr.next.next == nil {
			// 偶数个结点
			rs = append(rs, slowPtr)
			rs = append(rs, slowPtr.next)
			break
		}
		slowPtr = slowPtr.next
		quickPtr = quickPtr.next.next
	}
	return rs
}

func (l *SinglyLinkedList) reverse() {
	if l.head.next == nil {
		return
	}

	var preNode = l.head
	var ptrNode = l.head.next
	for {
	}
}

func (l *SinglyLinkedList) swapTwoNodes(preNode, ptrNode *Node) {
	var tmpNode *Node

	if preNode == l.head {
	}
}

func main() {
	var l = &SinglyLinkedList{}
	for i := 10; i > 0; i-- {
		l.insertValueToHead(i)
	}
	fmt.Println(l)

	// insert after one node
	node8 := l.findByValue(8)
	fmt.Println(node8)

	// insert after one node
	l.insertValueAfter(node8, 88)
	fmt.Println(l)

	// insert before one node
	l.insertValueBefore(node8, 55)
	fmt.Println(l)
}
