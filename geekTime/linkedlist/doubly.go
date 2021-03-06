/**
双链表
增删查
*/
package linkedlist

import "fmt"

type DoublyLinkedList struct {
	head *DNode
	cnt  int
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
		cnt:  0,
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

func (dll *DoublyLinkedList) Size() int {
	return dll.cnt
}

func (dll *DoublyLinkedList) Get(idx int) *DNode {
	// get by index
	if dll.head == nil {
		return nil
	}

	var j int
	p := dll.head.next

	for {
		if j == idx {
			break
		}
		if p == nil {
			break
		}

		p = p.next
		j++
	}
	return p
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
	// insert element to the head
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
	dll.cnt++
}

func (dll *DoublyLinkedList) Append(v interface{}) {
	// insert element to the tail
	node := newDoublyNode(v)
	dll.AppendNode(node)
}

func (dll *DoublyLinkedList) AppendNode(node *DNode) {
	p := dll.head
	for {
		if p.next == nil {
			node.prev = p
			p.next = node
			break
		}
		p = p.next
	}
	dll.cnt++
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
	dll.cnt--
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
