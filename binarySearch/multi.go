/**
假设给定的一组数据是从小到大排列的
1. 二分查找第一个值等于给定值的元素
2. 二分查找最后一个值等于给定值的元素
3. 查找第一个大于等于给定值的元素
4. 查找最后一个小于等于给定值的元素
*/
package main

import "fmt"

func main() {
	// TestFindFirstEqual()
	// TestFindLastEqual()
	// TestFindFirstGreaterEqual()
	TestFindLastLessEqual()
}

func TestFindFirstEqual() {
	var arr []int
	var pos int

	arr = []int{3, 4, 5, 5, 7, 7, 7, 7, 8}
	for _, x := range []int{5, 7, 3, 8, 10, 6, 1} {
		pos = findFirstEqual(x, 0, len(arr)-1, arr)
		fmt.Printf("x: %d, pos: %d\n", x, pos)
	}
}

func findFirstEqual(x, i, j int, arr []int) int {
	var mid int

	for {
		if i > j {
			break
		}

		mid = (i + j) / 2

		if arr[mid] < x {
			i = mid + 1
		} else {
			if arr[mid] == x {
				if mid == 0 || (mid > 0 && arr[mid-1] < x) {
					return mid
				}
			}

			j = mid - 1
		}
	}

	return -1
}

func TestFindLastEqual() {
	var arr []int
	var pos int

	arr = []int{3, 4, 5, 5, 7, 7, 7, 7, 8}
	for _, x := range []int{5, 7, 3, 8, 10, 6, 1} {
		pos = findLastEqual(x, 0, len(arr)-1, len(arr)-1, arr)
		fmt.Printf("x: %d, pos: %d\n", x, pos)
	}
}

func findLastEqual(x, i, j, lastnum int, arr []int) int {
	var mid int

	if i > j {
		return -1
	}

	mid = (i + j) / 2

	if arr[mid] > x {
		return findLastEqual(x, i, mid-1, lastnum, arr)
	} else {
		if arr[mid] == x {
			if mid == lastnum || (mid < lastnum && arr[mid+1] > x) {
				return mid
			}
		}

		return findLastEqual(x, mid+1, j, lastnum, arr)
	}
}

func TestFindFirstGreaterEqual() {
	var arr []int
	var pos int

	arr = []int{3, 4, 5, 5, 7, 7, 7, 7, 8}
	for _, x := range []int{5, 7, 3, 8, 10, 6, 1} {
		pos = findFirstGreaterEqual(x, 0, len(arr)-1, arr)
		fmt.Printf("x: %d, pos: %d\n", x, pos)
	}
}

func findFirstGreaterEqual(x, i, j int, arr []int) int {
	var mid int

	if i > j {
		return -1
	}

	mid = (i + j) / 2

	if arr[mid] < x {
		return findFirstGreaterEqual(x, mid+1, j, arr)
	} else {
		if arr[mid] >= x {
			if mid == 0 || (mid > 0 && arr[mid-1] < x) {
				return mid
			}
		}

		return findFirstGreaterEqual(x, i, mid-1, arr)
	}
}

func TestFindLastLessEqual() {
	var arr []int
	var pos int

	arr = []int{3, 4, 5, 5, 7, 7, 7, 7, 8}
	for _, x := range []int{5, 7, 3, 8, 10, 6, 1} {
		pos = findLastLessEqual(x, 0, len(arr)-1, len(arr)-1, arr)
		fmt.Printf("x: %d, pos: %d\n", x, pos)
	}
}

func findLastLessEqual(x, i, j, lastnum int, arr []int) int {
	var mid int

	for {
		if i > j {
			break
		}

		mid = (i + j) / 2

		if arr[mid] > x {
			j = mid - 1
		} else {
			if arr[mid] <= x {
				if mid == lastnum || (mid < lastnum && arr[mid+1] > x) {
					return mid
				}
			}

			i = mid + 1
		}
	}

	return -1
}
