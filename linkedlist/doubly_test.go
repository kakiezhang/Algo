package linkedlist

import (
	"fmt"
	"testing"
)

func TestDoublyInsert(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.Print()
	for _, v := range []int{18, 3, 7, 91, 74} {
		fmt.Printf("insert x: %d\n", v)
		dll.Insert(v)
		dll.Print()
	}
}

func TestDoublyFindNode(t *testing.T) {
	var dll *DoublyLinkedList
	dll = NewDoublyLinkedList()

	fmt.Printf("find 3? node: %v\n", dll.FindNode(3))

	for _, v := range []int{18, 3, 7, 91, 74} {
		dll.Insert(v)
	}
	dll.Print()

	for _, v := range []int{3, 33, 18, 74} {
		fmt.Printf("find %d? node: %v\n", v, dll.FindNode(v))
	}
}

func TestDoublyDelete(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.Delete(3)
	fmt.Println("delete 3: ")
	dll.Print()

	for _, v := range []int{18, 3, 7, 91, 74} {
		dll.Insert(v)
	}
	dll.Print()

	dll.Delete(74)
	fmt.Println("delete 74: ")
	dll.Print()
}
