/**
选择排序：
分为已排序和未排序两个区间，
每次从未排序中获取一个最小值，
插入到已排序区间的末尾
稳定性：否，反例：{2, 5, 6, 6, 4}
最好，最坏，平均时间复杂度：O(n2)
*/
package main

import "fmt"

func main() {
	var arr []int
	// arr = []int{6, 2, 5, 3, 1, 4}
	arr = []int{1, 2, 3, 4, 5, 6}
	select_sort(arr, len(arr))
	fmt.Println(arr)
}

func select_sort(arr []int, num int) {
	if num <= 1 {
		return
	}

	for i := 0; i < num-1; i++ {
		fmt.Println(arr)
		j := i + 1
		for ; j < num; j++ {
			if arr[i] > arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
}
