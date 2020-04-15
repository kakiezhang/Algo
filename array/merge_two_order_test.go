/**
两个有序数组合并为一个有序数组
*/
package array

import "testing"

func TestMergeTwoOrder(t *testing.T) {
	var arr1, arr2, rs []int

	arr1 = []int{3, 4, 5, 6}
	arr2 = []int{1, 8, 9}

	arr1 = []int{3}
	arr2 = []int{1}

	arr1 = []int{2, 4, 7, 10, 13, 14, 15, 16}
	arr2 = []int{3, 9, 20}

	arr1 = []int{2, 4, 7, 10}
	arr2 = []int{3, 9, 13, 14, 15, 16, 20}

	rs = mergeTwoOrder(arr1, arr2, len(arr1), len(arr2))
	t.Logf("arr1: %+v", arr1)
	t.Logf("arr2: %+v", arr2)
	t.Logf("rs: %+v", rs)
}
