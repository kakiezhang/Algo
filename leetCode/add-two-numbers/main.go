package main

import (
	"fmt"
)

func main() {
	var l1, l2, rs *ListNode

	{
		l1 = &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		}
		l2 = &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 4,
				},
			},
		}
		rs = addTwoNumbers(l1, l2)
		fmt.Printf("addTwoNumbers(%+v, %+v) = %+v\n", l1, l2, rs)
	}

	{
		l1 = &ListNode{
			Val: 0,
		}
		l2 = &ListNode{
			Val: 0,
		}
		rs = addTwoNumbers(l1, l2)
		fmt.Printf("addTwoNumbers(%+v, %+v) = %+v\n", l1, l2, rs)
	}

	{
		l1 = &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
								},
							},
						},
					},
				},
			},
		}
		l2 = &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		}
		rs = addTwoNumbers(l1, l2)
		fmt.Printf("addTwoNumbers(%+v, %+v) = %+v\n", l1, l2, rs)
	}
}

func (n *ListNode) String() (s string) {
	for p := n; p != nil; p = p.Next {
		s += fmt.Sprintf("%d", p.Val)
	}
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rs := &ListNode{}
	p := rs
	x3 := 0

	i := 0
	quit1 := false
	quit2 := false

	for {
		x1, x2 := 0, 0

		if l1 != nil {
			x1 = l1.Val
			l1 = l1.Next
		} else {
			quit1 = true
		}

		if l2 != nil {
			x2 = l2.Val
			l2 = l2.Next
		} else {
			quit2 = true
		}

		if quit1 && quit2 && x3 == 0 {
			p.Next = nil
			break
		}

		total := x1 + x2 + x3

		x3 = total / 10

		node := &ListNode{
			Val:  total % 10,
			Next: &ListNode{},
		}

		if i == 0 {
			rs = node
			p = rs
		}

		if p != nil {
			p.Next = node
			p = p.Next
		}

		i++
	}

	return rs
}
