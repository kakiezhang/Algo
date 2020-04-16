/**
顺序栈实现游览器前进，后退功能
*/
package main

import "fmt"

const MAXNUM = 2

func main() {
	sb := newStackBrower()
	sb.visit("a")
	sb.visit("b")
	// sb.visit("c")
	sb.backward()
	sb.backward()
	sb.foreward()
	sb.visit("d")
}

func newStackBrower() *stackBrowser {
	return &stackBrowser{
		sfore: newStackA("fore", MAXNUM),
		sback: newStackA("back", MAXNUM),
	}
}

func newStackA(name string, num int) *stackA {
	return &stackA{
		name: name,
		num:  num,
		cnt:  0,
		arr:  make([]string, num),
	}
}

type stackBrowser struct {
	sfore *stackA
	sback *stackA
}

func (sb *stackBrowser) visit(x string) {
	fmt.Printf("visit page[%s]\n", x)

	sb.sback.clear()
	sb.sfore.push(x)

	fmt.Printf("%s\n", sb.sfore)
	fmt.Printf("%s\n", sb.sback)
	fmt.Println()
}

func (sb *stackBrowser) backward() {
	fmt.Println("backward page")

	x := sb.sfore.pop()
	sb.sback.push(x)

	fmt.Printf("%s\n", sb.sfore)
	fmt.Printf("%s\n", sb.sback)
	fmt.Println()
}

func (sb *stackBrowser) foreward() {
	fmt.Println("foreward page")

	x := sb.sback.pop()
	sb.sfore.push(x)

	fmt.Printf("%s\n", sb.sfore)
	fmt.Printf("%s\n", sb.sback)
	fmt.Println()
}

type stackA struct {
	name string
	num  int
	cnt  int
	arr  []string
}

func (sa *stackA) String() string {
	var s string
	s = fmt.Sprintf("[%sstack bottom] ", sa.name)
	for i := 0; i < sa.cnt; i++ {
		s += fmt.Sprintf("[%s] ", sa.arr[i])
	}
	s += " top"
	return s
}

func (sa *stackA) clear() {
	sa.cnt = 0
}

func (sa *stackA) push(x string) {
	if sa.cnt == sa.num {
		panic("brower max page")
	}

	sa.arr[sa.cnt] = x
	sa.cnt += 1
}

func (sa *stackA) pop() string {
	if sa.cnt == 0 {
		panic("brower no page")
	}

	x := sa.arr[sa.cnt-1]
	sa.cnt -= 1
	return x
}
