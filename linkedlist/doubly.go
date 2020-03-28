/**
双链表
*/
package linkedlist

type doublyLinkedList struct {
	head *dNode
}

type dNode struct {
	data interface{}
	prev *dNode
	next *dNode
}

func NewDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (dll *doublyLinkedList) Insert() {
}

func (dll *doublyLinkedList) Delete() {
}

func (dll *doublyLinkedList) Find() {
}
