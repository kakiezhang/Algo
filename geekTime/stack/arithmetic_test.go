package stack

import "testing"

func TestFourPronged(t *testing.T) {
	var s string
	var rs int

	// pass
	s = "2+3x5-7"
	// s = "2+3x5"
	// s = "9+3x8/4"
	// s = "3x8/4-10"
	s = "3x8/4-10+16x8/4"
	s = "3x8/4-10+16x3/4"

	// fail
	s = "3x8/4-10+16x3/4*10-19"

	fp := NewFourPronged(s, len(s))
	rs = fp.Operate()

	t.Logf("%s=%d", s, rs)
}
