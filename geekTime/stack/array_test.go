package stack

import "testing"

func TestArrayStack(t *testing.T) {
	as := NewArrayStack(2)

	for _, v := range []string{"apple", "banana"} {
		as.Push(v)
		t.Logf("[push] stack: %+v, x: %v", as, v)
	}

	x := as.Pop()
	t.Logf("[pop] stack: %+v, x: %v", as, x)

	t.Logf("stack top: %v", as.Top())

	y := as.Pop()
	t.Logf("[pop] stack: %+v, x: %v", as, y)

	as.Pop()
}
