package stack

import "testing"

func TestNewParentheses(t *testing.T) {
	var s string
	var rs bool

	s = "{({[()]})}"
	s = "{(()})}"
	s = "{[]()}{[[]]}([{}()]{[[[]]]})"

	pth := NewParentheses(s, len(s))
	rs = pth.match()

	t.Logf("s: %s, match: %t", s, rs)
}
