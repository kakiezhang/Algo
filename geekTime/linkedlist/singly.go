/**
单链表
1. 增删查
2. 反转，求中间节点(原地)
*/
package linkedlist

import "fmt"

type SinglyLinkedList struct {
	head *sNode
}

type sNode struct {
	data interface{}
	next *sNode
}

func (sn *sNode) String() string {
	return fmt.Sprintf("[%v %p][%p]", sn.data, sn, sn.next)
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: newSinglyNode(nil), // sentinel
	}
}

func newSinglyNode(v interface{}) *sNode {
	return &sNode{
		data: v,
	}
}

func (sll *SinglyLinkedList) Print() {
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

func (sll *SinglyLinkedList) Insert(v interface{}) {
	node := newSinglyNode(v)
	sll.InsertNode(node)
}

func (sll *SinglyLinkedList) InsertNode(node *sNode) {
	node.next = sll.head.next
	sll.head.next = node
}

func (sll *SinglyLinkedList) Delete(v interface{}) {
	node, prev := sll.FindNode(v)
	if prev == nil {
		return
	}
	prev.next = node.next
}

func (sll *SinglyLinkedList) FindNode(v interface{}) (*sNode, *sNode) {
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

func (sll *SinglyLinkedList) reverse() {
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

func (sll *SinglyLinkedList) findMid() *sNode {
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
