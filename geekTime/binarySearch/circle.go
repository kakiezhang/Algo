/**
从循环有序数组二分查找一个值
*/
package main

import "fmt"

func main() {
	var arr []int
	var x, pos int

	// x = -1
	// arr = []int{4, 5, 6, 1, 2, 3}
	// pos = bisearch(x, 0, len(arr)-1, arr)
	// fmt.Printf("pos: %d\n", pos)

	arr = []int{6, 7, 8, 9, 12, 16, 2, 3, 4, 5}
	// for _, x = range []int{6, 7, 8, 9, 12, 16, 2, 3, 4, 5, -1, 100} {
	for _, x = range []int{16} {
		pos = bisearch(x, 0, len(arr)-1, arr)
		fmt.Printf("x: %d, pos: %d\n", x, pos)
	}

}

func bisearch(x, i, j int, arr []int) int {
	if i > j {
		return -1
	}

	mid := (i + j) / 2
	fmt.Printf("x: %d, i: %d, j: %d, mid: %d\n", x, i, j, mid)

	if x == arr[mid] {
		return mid
	} else if x < arr[mid] {
		if x == arr[i] {
			return i
		} else if x > arr[i] {
			if arr[i] < arr[mid] {
				return bisearch(x, i+1, mid-1, arr)
			} else {
				panic("impossible")
			}
		} else {
			if arr[i] < arr[mid] {
				return bisearch(x, mid+1, j, arr)
			} else {
				return bisearch(x, i+1, mid-1, arr)
			}
		}
	} else {
		// x > arr[mid]
		if x == arr[j] {
			return j
		} else if x < arr[j] {
			if arr[j] > arr[mid] {
				return bisearch(x, mid+1, j-1, arr)
			} else {
				panic("impossible")
			}
		} else {
			// x > arr[j]
			if arr[j] > arr[mid] {
				return bisearch(x, i, mid-1, arr)
			} else {
				return bisearch(x, mid+1, j-1, arr)
			}
		}
	}
}
