package sort

import "testing"

func TestShellSort(t *testing.T) {
	var arr []int

	arr = []int{1}
	t.Logf("before: %v", arr)
	shellSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{2, 1}
	t.Logf("before: %v", arr)
	shellSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{3, 1, 2}
	t.Logf("before: %v", arr)
	shellSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{3, 2, 1}
	t.Logf("before: %v", arr)
	shellSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{3, 2, 1, 4, 7, 5, 6}
	t.Logf("before: %v", arr)
	shellSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{5, 2, 4, 1, 2, 4, 7, 5, 6}
	t.Logf("before: %v", arr)
	shellSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{
		13, 14, 94, 33, 82, 25, 59, 94,
		65, 23, 45, 27, 73, 25, 39, 10}
	t.Logf("before: %v", arr)
	shellSort(arr, len(arr))
	t.Logf("after: %v", arr)
}
