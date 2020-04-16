/**
有序排重数组的二分查找
两种方法实现：普通循环和递归
*/
package main

import "fmt"

func main() {
	// testBiLoopFind()
	testBiRecursiveFind()
}

func testBiLoopFind() {
	var arr []int
	var x int
	var pos int

	arr = []int{1, 4, 7, 12, 18, 25}
	x = 7

	pos = biLoopFind(x, len(arr), arr)
	fmt.Println(pos)
}

func biLoopFind(x, num int, arr []int) int {
	var i, j int
	j = num - 1

	for i <= j {
		mid := (i + j) / 2
		fmt.Printf("i: %d, j: %d, mid: %d\n", i, j, mid)

		if x < arr[mid] {
			j = mid - 1
		} else if x > arr[mid] {
			i = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func testBiRecursiveFind() {
	var arr []int
	var x int
	var pos int

	arr = []int{1, 4, 7, 12, 18, 25}
	x = 18

	pos = biRecursiveFind(x, 0, len(arr)-1, arr)
	fmt.Println(pos)
}

func biRecursiveFind(x, i, j int, arr []int) int {
	if i > j {
		return -1
	}

	mid := (i + j) / 2
	fmt.Printf("i: %d, j: %d, mid: %d\n", i, j, mid)

	if x < arr[mid] {
		return biRecursiveFind(x, i, mid-1, arr)
	} else if x > arr[mid] {
		return biRecursiveFind(x, mid+1, j, arr)
	} else {
		return mid
	}
}
