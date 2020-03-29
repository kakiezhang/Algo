/**
单链表
1. 增删查
2. 反转，求中间节点(原地)
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
	return fmt.Sprintf("[%v %p][%p]", sn.data, sn, sn.next)
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

func (sll *singlyLinkedList) Delete(v interface{}) {
	node, parent := sll.findNode(v)
	if parent == nil {
		return
	}

	parent.next = node.next
}

func (sll *singlyLinkedList) findNode(v interface{}) (*sNode, *sNode) {
	p := sll.head
	if p == nil {
		return nil, nil
	}

	for {
		if p.next == nil {
			return nil, nil
		}

		if p.next.data == v {
			return p.next, p
		}
		p = p.next
	}
}

func (sll *singlyLinkedList) reverse() {
	if sll.head == nil {
		return
	}

	p := sll.head.next
	var t *sNode // 反转后的尾节点

	for {
		if p == nil {
			break
		}

		n := p.next // 当前的next先保存起来
		p.next = t  // 把当前的next赋值为t
		t = p       // t更新成当前节点

		if n == nil {
			break
		}
		p = n // 移到原来的后面那个节点
	}

	sll.head.next = p
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
