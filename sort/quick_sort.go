/**
快速排序

首先需要找分区点，一般选择数组的最后一个元素
；将比分区点小的元素放在分区点前面，比分区点
大的放在后面；(这部分会写成一个分区函数，最
终会返回分区点在数组中的位置）

然后在分区点前后的值分别组成两个数组，分别使
用分区函数，一直到分区点前后数组为空为止，即
所有的数都被排序完成

由上可知，快排是一个自顶向下的过程



*/
package main

import "fmt"

func main() {
	var arr []int

	// arr = []int{3, 6, 1, 4, 7, 2, 5}
	// arr = []int{2, 1, 4, 5, 3}
	arr = []int{7, 8, 4, 7, 3, 2, 1, 6, 5, 3}
	// i := partition(arr, 0, len(arr)-1)
	// fmt.Println(i)

	quick_sort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quick_sort(arr []int, p, q int) {
	if p < 0 {
		return
	}
	if p >= q {
		return
	}

	i := partition(arr, p, q)
	fmt.Println(i)
	quick_sort(arr, p, i-1)
	quick_sort(arr, i+1, q)
}

func partition(arr []int, p, q int) int {
	var i = p
	var j = p

	pivot := arr[q]

	for ; j <= q; j++ {
		if arr[j] < pivot {
			tmp := arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
			i += 1
		}
	}

	arr[i], arr[j-1] = arr[j-1], arr[i]

	return i
}
