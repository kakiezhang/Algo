/**
单链表
增删查，反转，求中间节点
*/
package linkedlist

import "fmt"

type singlyLinkedList struct {
	head *sNode
}

type sNode struct {
	data interface{}
	next *sNode
}

func (sn *sNode) String() string {
	return fmt.Sprintf("[%v %p]", sn.data, sn.next)
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		head: newSinglyNode(nil), // sentinel
	}
}

func newSinglyNode(v interface{}) *sNode {
	return &sNode{
		data: v,
	}
}

func (sll *singlyLinkedList) Print() {
	if sll.head == nil {
		return
	}

	fmt.Printf("[SinglyLinkedList %p]: ", sll)
	p := sll.head
	for {
		fmt.Printf("%v->", p)
		p = p.next
		if p == nil {
			break
		}
	}
	fmt.Println()
}

func (sll *singlyLinkedList) Insert(v interface{}) {
	node := newSinglyNode(v)
	sll.insertNode(node)
}

func (sll *singlyLinkedList) insertNode(node *sNode) {
	node.next = sll.head.next
	sll.head.next = node
}

func (sll *singlyLinkedList) Delete() {
}

func (sll *singlyLinkedList) Find() {
}

func (sll *singlyLinkedList) Reverse() {
}

func (sll *singlyLinkedList) findMid() *sNode {
	// 利用快慢指针
	// 快指针每次走两步，慢指针走一步
	// 快指针到结尾的时候，慢指针就在中点
	if sll.head == nil {
		return nil
	}

	slow := sll.head.next
	if slow == nil {
		return nil
	}
	quick := slow.next

	for {
		if quick == nil || quick.next == nil {
			break
		}

		slow = slow.next // quick存在，slow.next一定存在
		quick = quick.next.next
	}
	return slow
}
