/**
双链表
增删查
*/
package linkedlist

import "fmt"

type DoublyLinkedList struct {
	head *DNode
}

type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}

type DirectDNodeData func(interface{}) interface{}

func (dll *DoublyLinkedList) String() string {
	if dll.head == nil {
		return ""
	}

	rs := fmt.Sprintf("[%p]: ", dll)
	p := dll.head
	for {
		rs += fmt.Sprintf("%v->", p)
		p = p.next
		if p == nil {
			break
		}
	}
	return rs
}

func (dn *DNode) String() string {
	return fmt.Sprintf("[%p][%v %p][%p]", dn.prev, dn.data, dn, dn.next)
}

func (dn *DNode) GetData() interface{} {
	return dn.data
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: newDoublyNode(nil), // sentinel
	}
}

func newDoublyNode(v interface{}) *DNode {
	return &DNode{
		data: v,
	}
}

func (dll *DoublyLinkedList) Print() {
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

func (dll *DoublyLinkedList) Poll() *DNode {
	// get & pop first
	if dll.head == nil {
		return nil
	}

	first := dll.head.next
	dll.DeleteNode(first)
	return first
}

func (dll *DoublyLinkedList) Insert(v interface{}) {
	node := newDoublyNode(v)
	dll.InsertNode(node)
}

func (dll *DoublyLinkedList) InsertNode(node *DNode) {
	node.next = dll.head.next
	if node.next != nil {
		node.next.prev = node
	}
	node.prev = dll.head
	dll.head.next = node
}

func (dll *DoublyLinkedList) Delete(v interface{}) {
	node := dll.FindNode(v, nil)
	dll.DeleteNode(node)
}

func (dll *DoublyLinkedList) DeleteNode(node *DNode) {
	if node == nil {
		return
	}

	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev.next = node.next
}

func (dll *DoublyLinkedList) FindNode(
	v interface{}, f DirectDNodeData) *DNode {
	if dll.head == nil {
		return nil
	}

	p := dll.head.next

	for {
		if p == nil {
			break
		}

		pv := p.data
		if f != nil {
			pv = f(p.data)
		}

		if pv == v {
			break
		}
		p = p.next
	}
	return p
}
