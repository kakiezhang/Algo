package sort

import "testing"

func TestQuickSort(t *testing.T) {
	var arr []int

	arr = []int{1}
	t.Logf("before: %v", arr)
	quickSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{2, 1}
	t.Logf("before: %v", arr)
	quickSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{3, 1, 2}
	t.Logf("before: %v", arr)
	quickSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{3, 2, 1}
	t.Logf("before: %v", arr)
	quickSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{3, 2, 1, 4, 7, 5, 6}
	t.Logf("before: %v", arr)
	quickSort(arr, len(arr))
	t.Logf("after: %v", arr)

	arr = []int{5, 2, 4, 1, 2, 4, 7, 5, 6}
	t.Logf("before: %v", arr)
	quickSort(arr, len(arr))
	t.Logf("after: %v", arr)
}
