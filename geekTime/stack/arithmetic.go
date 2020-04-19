/**
四则运算
分为操作数和操作符号两个栈，依次进栈
操作符这边假如等待进栈的符号优先级<=操作符栈栈顶的元素，
那就先从栈中弹出一个符号，两个数，做完操作后，再压栈
*/
package stack

import (
	"strconv"
)

var (
	// 操作符的优先级
	opPrior = map[string]int{
		"+": 1,
		"-": 1,
		"x": 2,
		"/": 2,
	}
	// 操作符的元素方式
	opCal = map[string]CalFunc{
		"+": plus,
		"-": minus,
		"x": multiply,
		"/": divide,
	}
)

type FourPronged struct {
	num *ArrayStack // 操作数
	op  *ArrayStack // 操作符
	s   string
	cnt int
}

type FPOperation struct {
	name  string
	prior int
}

func NewFourPronged(s string, cnt int) *FourPronged {
	return &FourPronged{
		num: NewArrayStack(cnt),
		op:  NewArrayStack(cnt),
		s:   s,
		cnt: cnt,
	}
}

func (fp *FourPronged) Operate() int {
	for i := 0; i < fp.cnt; {
		ele := string(fp.s[i])

		if ele == " " {
			i++
			continue
		}

		if nowOpPrior, ok := opPrior[ele]; ok {
			// 操作符
			for {
				lastOp := fp.op.Top()

				if lastOp == nil {
					fp.op.Push(&FPOperation{
						name:  ele,
						prior: nowOpPrior,
					})
					break
				} else {
					if lastOp.(*FPOperation).prior >= nowOpPrior {
						fp.OneStep()
					} else {
						fp.op.Push(&FPOperation{
							name:  ele,
							prior: nowOpPrior,
						})
						break
					}
				}
			}
			i++
		} else {
			// 操作数
			var data int

			j := i + 1
			for ; j <= fp.cnt; j++ {
				if v, err := strconv.Atoi(string(fp.s[i:j])); err == nil {
					data = v
				} else {
					break
				}
			}

			fp.num.Push(data)
			i = j - 1
		}
	}

	for {
		if fp.op.Top() == nil {
			if fp.num.Top() != nil {
				return fp.num.Pop().(int)
			} else {
				panic("error")
			}
		}

		fp.OneStep()
	}
}

func (fp *FourPronged) OneStep() {
	// fmt.Printf("[before] num: %v, op: %v\n", fp.num, fp.op)
	n1 := fp.num.Pop().(int)
	n2 := fp.num.Pop().(int)
	op := fp.op.Pop().(*FPOperation).name
	n3 := opCal[op](n1, n2)
	fp.num.Push(n3)
	// fmt.Printf("[after] num: %v, op: %v\n", fp.num, fp.op)
}

type CalFunc func(x, y int) int

func plus(x, y int) int {
	return y + x
}

func minus(x, y int) int {
	return y - x
}

func multiply(x, y int) int {
	return y * x
}

func divide(x, y int) int {
	return y / x
}
