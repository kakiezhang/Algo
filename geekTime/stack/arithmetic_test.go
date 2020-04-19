package stack

import "testing"

func TestFourPronged(t *testing.T) {
	var s string
	var rs int

	s = "2+3x5- 7"

	fp := NewFourPronged(s, len(s))
	rs = fp.Operate()

	t.Logf("%s=%d", s, rs)
}
