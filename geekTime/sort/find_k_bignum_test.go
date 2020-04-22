package sort

import "testing"

func TestFindKBigNum(t *testing.T) {
	var arr []int
	var knum int

	arr = []int{1}
	knum = findKBigNum(arr, len(arr), 1)
	t.Logf("arr: %v, knum: %d", arr, knum)

	arr = []int{2, 1}
	knum = findKBigNum(arr, len(arr), 2)
	t.Logf("arr: %v, knum: %d", arr, knum)

	arr = []int{3, 1, 2}
	knum = findKBigNum(arr, len(arr), 1)
	t.Logf("arr: %v, knum: %d", arr, knum)

	arr = []int{3, 2, 1}
	knum = findKBigNum(arr, len(arr), 2)
	t.Logf("arr: %v, knum: %d", arr, knum)

	arr = []int{3, 2, 1, 4, 7, 5, 6}
	knum = findKBigNum(arr, len(arr), 5)
	t.Logf("arr: %v, knum: %d", arr, knum)

	arr = []int{5, 2, 4, 1, 2, 4, 7, 5, 6}
	knum = findKBigNum(arr, len(arr), 8)
	t.Logf("arr: %v, knum: %d", arr, knum)
}
