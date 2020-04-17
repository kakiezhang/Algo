package majorityElement

import "testing"

func TestMajorNumMemo(t *testing.T) {
	var arr []int
	var rs int

	arr = []int{3, 2, 3}
	arr = []int{2, 2, 1, 1, 1, 2, 2}

	rs = MajorNumMemo(arr, len(arr))
	t.Logf("rs: %d", rs)
}

func TestMajorNumMiddle(t *testing.T) {
	var arr []int
	var rs int

	arr = []int{3, 2, 3}
	arr = []int{2, 2, 1, 1, 1, 1, 2}

	rs = MajorNumMiddle(arr, len(arr))
	t.Logf("rs: %d", rs)
}
