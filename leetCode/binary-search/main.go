package main

import "fmt"

func main() {
	x := 20
	arr := []int{0, 20, 22, 81, 100}
	fmt.Printf("arr: %+v, x: %d, loc: %d\n", arr, x, binary_search_2(x, arr))
}

func binary_search_1(target int, arr []int) int {
	totalnum := len(arr)
	start := 0
	end := totalnum - 1

	for start <= end {
		mid := (start + end) / 2

		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func binary_search_2(target int, arr []int) int {
	totalnum := len(arr)
	start := 0
	end := totalnum - 1

	return recursion(target, arr, start, end)

}

func recursion(target int, arr []int, start, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if target == arr[mid] {
		return mid
	} else if target > arr[mid] {
		return recursion(target, arr, mid+1, end)
	} else {
		return recursion(target, arr, start, mid-1)
	}
}
