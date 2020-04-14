package array

import "testing"

func TestArray(t *testing.T) {
	a := NewArray(1)
	for _, v := range []int{10, 15, 20, 18, 29, 30, 34, 27} {
		a.Add(v)
		t.Logf("add: %d, array: %+v\n", v, a)
	}
}
