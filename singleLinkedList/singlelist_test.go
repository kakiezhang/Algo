package main

import (
	"testing"
)

func TestNode(t *testing.T) {
	// t.Fatal("not implemented")
	n := &Node{
		data: 1,
		next: &Node{
			data: 2,
			next: nil,
		},
	}
	t.Logf("%p", n)
	t.Logf("%p", n.next)
}

func TestInsertAfter(t *testing.T) {
	var l = &SinglyLinkedList{}
	for i := 10; i > 0; i-- {
		l.insertValueToHead(i)
	}
	t.Log(l)

	// foundNode := l.findByValue(0)
	// foundNode := l.findByValue(10)
	foundNode := l.findByValue(1)

	if foundNode == nil {
		t.Fatal("cannot find node")
		return
	}
	t.Log(foundNode)

	// insert after one node
	l.insertValueAfter(foundNode, 88)
	t.Log(l)
}

func TestInsertBefore(t *testing.T) {
	var l = &SinglyLinkedList{}
	for i := 5; i > 0; i-- {
		l.insertValueToHead(i)
	}
	t.Log(l)

	// foundNode := l.findByValue(0)
	// foundNode := l.findByValue(1)
	// foundNode := l.findByValue(3)
	foundNode := l.findByValue(5)

	if foundNode == nil {
		t.Fatal("cannot find node")
		return
	}
	t.Log(foundNode)

	l.insertValueBefore(foundNode, 88)
	t.Log(l)
}

func TestFindMiddleNodes(t *testing.T) {
	var l = &SinglyLinkedList{}
	for i := 6; i > 0; i-- {
		l.insertValueToHead(i)
	}
	t.Log(l)

	nodes := l.findMiddleNodes()
	if len(nodes) == 0 {
		t.Fatal("cannot find nodes")
		return
	}
	for _, n := range nodes {
		t.Log(n)
	}
}
