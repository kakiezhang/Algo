/**
插入排序
分为已排序和未排序两个区间

在已排序空间使用倒叙：
每次从未排序区间获取一个值，然后在已排序区间中按倒叙的方式
去寻找第一个大于的位置，并且在其后面将元素插入（每次比较的
时候，如果已排序空间的元素比给定值大，就将已排序空间的元素
往后移一个位置）

时间复杂度：
最好 O(n)，{1, 2, 3, 4, 5, 6}
最差 O(n2)，{6, 5, 4, 3, 2, 1}

稳定性：稳定

空间复杂度：原地排序
*/
package main

import "fmt"

func main() {
	var arr []int

	// arr = []int{8, 5, 7, 1, 3, 2}
	// insert_desc(arr, len(arr))
	// fmt.Println(arr)

	// arr = []int{8, 5, 7, 7, 2, 2}
	// insert_desc(arr, len(arr))
	// fmt.Println(arr)

	// arr = []int{1, 2, 3, 4, 5, 6}
	// insert_desc(arr, len(arr))
	// fmt.Println(arr)

	arr = []int{6, 5, 4, 3, 2, 1}
	insert_desc(arr, len(arr))
	fmt.Println(arr)
}

func insert_desc(arr []int, num int) {
	if num <= 1 {
		return
	}

	for i := 1; i <= num-1; i++ {
		value := arr[i]

		j := i - 1
		for ; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1] = arr[j]
				fmt.Println(arr)
			} else {
				break
			}
		}

		arr[j+1] = value
	}
}
