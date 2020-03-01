/**
归并排序

首先，对数组进行分组，每次取数组的中点，对数组
中点前后区间再分别取中点，最后的颗粒度可以很小
，这里暂定为 2 个为一组

然后，每一组内部已经排序，排序完之后，组俩俩之间
再进行合并顺序排序，直至所有的数组被合并排序完

*/
package main

import "fmt"

func main() {
	var arr []int

	// arr = []int{3, 2, 8, 5, 1, 6}
	arr = []int{5, 2, 6, 3, 4, 3, 7}
	merge_sort(arr, 0, len(arr)-1)
}

func merge_sort(arr []int, p, q int) {
	tmp, _ := merge_sort_unit(arr, p, q)
	fmt.Println(tmp)
}

func merge_sort_unit(arr []int, p, q int) ([]int, int) {
	if p > q {
		return []int{}, 0
	} else if p == q {
		return []int{arr[p]}, 1
	} else if p+1 == q {
		if arr[p] <= arr[q] {
			return []int{arr[p], arr[q]}, 2
		} else {
			return []int{arr[q], arr[p]}, 2
		}
	}

	r := (p + q) / 2

	var s1, s2 []int
	var n1, n2 int

	s1, n1 = merge_sort_unit(arr, p, r)
	fmt.Printf("s1: %v, n1: %d\n", s1, n1)

	s2, n2 = merge_sort_unit(arr, r+1, q)
	fmt.Printf("s2: %v, n2: %d\n", s2, n2)

	var tmp = make([]int, n1+n2)
	var i, j, k int

	if n1 > 0 && n2 > 0 {
		for i < n1 && j < n2 {
			// fmt.Println(s1)
			// fmt.Println(i)
			if s1[i] <= s2[j] {
				// fmt.Println(s1[i])
				tmp[k] = s1[i]
				i += 1
			} else {
				tmp[k] = s2[j]
				j += 1
			}

			k += 1
		}
	}

	if i < n1 {
		for ; i <= n1-1; i++ {
			tmp[k] = s1[i]
			k += 1
		}
	} else if j < n2 {
		for ; j <= n2-1; j++ {
			tmp[k] = s2[j]
			k += 1
		}
	}

	return tmp, k
}
