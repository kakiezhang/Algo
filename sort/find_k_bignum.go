/**
寻找第 K 大元素

数组 arr 长度为 [0, n-1]
根据分区点返回的位置来决定
如果 i+1 == k，那 arr[i] 就是要找的元素
如果 i+1 > k，那在区间数组 [0, i-1] 中继续找
如果 i+1 < k，那在区间数组 [i+1, n-1] 中继续找

*/
package main

import "fmt"

func main() {
	k := 1
	var arr []int

	arr = []int{2, 4, 5, 1, 3}
	find(arr, 0, len(arr)-1, k)

	// i := partition(arr, 0, len(arr)-1)
	// fmt.Println(i)

	fmt.Println(arr)
}

func find(arr []int, p, q, k int) {
	if p < 0 || p > q {
		fmt.Println("cannot be found")
		return
	}

	i := partition(arr, p, q)
	if i+1 == k {
		fmt.Printf("found: num[%d]\n", arr[i])
	} else if i+1 > k {
		find(arr, 0, i-1, k)
	} else {
		find(arr, i+1, q, k)
	}
}

func partition(arr []int, p, q int) int {
	i := p
	j := p

	pivot := arr[q]

	for ; j <= q; j++ {
		if arr[j] > pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}

	arr[i], arr[j-1] = arr[j-1], arr[i]
	fmt.Printf("pos[%d]\n", i)

	return i
}
