/**
简化版的循环有序数组查找
原理：每次将数组一分为二，两个子数组分别为有序数组和循环有序数组
如果 x 出现在有序数组范围中，那就直接使用二分查找，否则就继续在
将数组一分为二

这种方式虽然实现起来优雅，但是实际的时间复杂度却要比简陋版的那个
要高
*/
package main

import "fmt"

func main() {
	var arr []int
	var x, pos int

	// x = 3
	// arr = []int{4, 5, 6, 1, 2, 3}
	// pos = bicircle(x, 0, len(arr)-1, arr)
	// fmt.Printf("pos: %d\n", pos)

	arr = []int{6, 7, 8, 9, 12, 16, 2, 3, 4, 5}
	for _, x = range []int{6, 7, 8, 9, 12, 16, 2, 3, 4, 5, -1, 100} {
		// for _, x = range []int{16} {
		pos = bicircle(x, 0, len(arr)-1, arr)
		fmt.Printf("x: %d, pos: %d\n", x, pos)
	}
}

func bisearch(x, i, j int, arr []int) int {
	if i > j {
		return -1
	}

	mid := (i + j) / 2
	// fmt.Printf("tm: x: %d, i: %d, j: %d, mid: %d\n", x, i, j, mid)

	if x == arr[mid] {
		return mid
	} else if x > arr[mid] {
		return bisearch(x, mid+1, j, arr)
	} else {
		return bisearch(x, i, mid-1, arr)
	}
}

func bicircle(x, i, j int, arr []int) int {
	if arr[i] <= arr[j] {
		if x >= arr[i] && x <= arr[j] {
			// fmt.Printf("bis: x: %d, i: %d, j: %d\n", x, i, j)
			return bisearch(x, i, j, arr)
		}
	}

	if arr[i] > arr[j] {
		mid := (i + j) / 2
		// fmt.Printf("bic: x: %d, i: %d, j: %d, mid: %d\n", x, i, j, mid)

		a := bicircle(x, i, mid, arr)
		if a >= 0 {
			return a
		}
		b := bicircle(x, mid+1, j, arr)
		if b >= 0 {
			return b
		}
	}

	return -1
}
