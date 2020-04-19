/**
数组实现顺序栈
*/
package stack

type ArrayStack struct {
	ele []interface{}
	top int // 栈顶位置
	cnt int // 栈内元素最大数量
}

func NewArrayStack(cnt int) *ArrayStack {
	return &ArrayStack{
		ele: make([]interface{}, cnt),
		top: -1,
		cnt: cnt,
	}
}

func (as *ArrayStack) Push(x interface{}) {
	if as.top == as.cnt-1 {
		panic("stack is full.")
	}

	as.top += 1
	as.ele[as.top] = x
}

func (as *ArrayStack) Pop() interface{} {
	if as.cnt == 0 {
		panic("stack is empty.")
	}

	x := as.ele[as.top]
	as.ele[as.top] = nil
	as.top -= 1
	return x
}

func (as *ArrayStack) Top() interface{} {
	if as.top < 0 {
		return nil
	}
	return as.ele[as.top]
}
