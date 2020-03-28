package linkedlist

import (
	"fmt"
	"testing"
)

func TestSinglyInsert(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.Print()
	for _, v := range []int{18, 3, 7, 91, 74} {
		fmt.Printf("insert x: %d\n", v)
		sll.Insert(v)
		sll.Print()
	}
}

func TestSinglyFindMid(t *testing.T) {
	sll := NewSinglyLinkedList()
	for _, v := range []int{18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28} {
		fmt.Printf("insert x: %d\n", v)
		sll.Insert(v)
		sll.Print()
		mid := sll.findMid()
		fmt.Printf("mid: %v", mid)
	}
}

func TestSinglyReverse(t *testing.T) {
	var sll *singlyLinkedList
	sll = NewSinglyLinkedList()
	for _, v := range []int{18, 3, 7, 91, 74} {
		sll.Insert(v)
	}
	fmt.Println("origin: ")
	sll.Print()

	fmt.Println("reverse: ")
	sll.reverse()
	sll.Print()

	fmt.Println()

	sll = NewSinglyLinkedList()
	for _, v := range []int{18} {
		sll.Insert(v)
	}
	fmt.Println("origin: ")
	sll.Print()

	fmt.Println("reverse: ")
	sll.reverse()
	sll.Print()
}
