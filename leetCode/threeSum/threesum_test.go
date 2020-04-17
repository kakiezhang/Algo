package threeSum

import "testing"

func TestThreeSumCornered(t *testing.T) {
	var arr []int
	var rs [][3]int
	var target int

	arr = []int{-1, 0, 1, 2, -1, -4}
	target = 0

	rs = ThreeSumCornered(arr, len(arr), target)
	t.Logf("rs: %v", rs)
}
