/**
括号匹配
查看左右括号是否匹配
假设只存在这[{(三种括号，并且只可能有一共六种字符可能
1. 如果是左半边，直接进栈
2. 如果是右半边，先比较栈顶是不是和其对应的左半边，不是则不匹配
*/
package stack

var (
	leftHalf = map[string]struct{}{
		"(": struct{}{},
		"[": struct{}{},
		"{": struct{}{},
	}

	rightToLeft = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
)

type Parentheses struct {
	*ArrayStack
	s   string
	cnt int
}

func NewParentheses(s string, cnt int) *Parentheses {
	return &Parentheses{
		ArrayStack: NewArrayStack(cnt),
		s:          s,
		cnt:        cnt,
	}
}

func (pth *Parentheses) match() bool {
	for i := 0; i < pth.cnt; i++ {
		ele := string(pth.s[i])

		if _, ok := leftHalf[ele]; ok {
			// 左半边
			pth.Push(ele)
		} else {
			// 右半边
			top := pth.Top()

			if top != nil && top.(string) == rightToLeft[ele] {
				pth.Pop()
			} else {
				return false
			}
		}
	}

	if pth.Top() == nil {
		return true
	} else {
		return false
	}
}
