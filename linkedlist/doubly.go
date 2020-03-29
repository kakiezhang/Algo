/**
双链表
*/
package linkedlist

import "fmt"

type doublyLinkedList struct {
	head *dNode
}

type dNode struct {
	data interface{}
	prev *dNode
	next *dNode
}

func (dn *dNode) String() string {
	return fmt.Sprintf("[%p][%v %p][%p]", dn.prev, dn.data, dn, dn.next)
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{
		head: newDoublyNode(nil), // sentinel
	}
}

func newDoublyNode(v interface{}) *dNode {
	return &dNode{
		data: v,
	}
}

func (dll *doublyLinkedList) Print() {
	if dll.head == nil {
		return
	}

	fmt.Printf("[DoublyLinkedList %p]: ", dll)
	p := dll.head
	for {
		fmt.Printf("%v->", p)
		p = p.next
		if p == nil {
			break
		}
	}
	fmt.Println()
}

func (dll *doublyLinkedList) Insert(v interface{}) {
	node := newDoublyNode(v)
	dll.insertNode(node)
}

func (dll *doublyLinkedList) insertNode(node *dNode) {
	node.next = dll.head.next
	if node.next != nil {
		node.next.prev = node
	}
	node.prev = dll.head
	dll.head.next = node
}

func (dll *doublyLinkedList) Delete(v interface{}) {
	node := dll.findNode(v)
	if node == nil {
		return
	}

	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev.next = node.next
}

func (dll *doublyLinkedList) findNode(v interface{}) *dNode {
	if dll.head == nil {
		return nil
	}

	p := dll.head.next

	for {
		if p == nil {
			break
		}
		if p.data == v {
			break
		}
		p = p.next
	}
	return p
}
