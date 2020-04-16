package main

import "fmt"

func main() {
	var arr []int

	// arr = []int{7, 8, 1, 3, 2, 6, 4, 5}
	arr = []int{2, 4, 5, 3, 2, 1, 4, 3}
	fmt.Println(arr)
	merge_sort(arr, len(arr))
	fmt.Println(arr)
}

func merge_sort(arr []int, num int) {
	merge_sort_unit(arr, 0, num-1)
}

func merge_sort_unit(arr []int, p, q int) {
	if p >= q {
		return
	}

	c := (p + q) / 2

	merge_sort_unit(arr, p, c)
	merge_sort_unit(arr, c+1, q)

	merge(arr, p, q, c)
}

func merge(arr []int, p, q, c int) {
	var k int

	i, j := p, c+1

	tn := q - p + 1

	fmt.Printf("tn: %d, p: %d, q: %d, c: %d\n", tn, p, q, c)

	var tmp = make([]int, tn)

	for i <= c && j <= q {
		fmt.Printf("i: %d, j: %d\n", arr[i], arr[j])
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i += 1
		} else {
			tmp[k] = arr[j]
			j += 1
		}
		k += 1
	}

	fmt.Printf("tmp: %v\n", tmp)

	var start, end int

	if i <= c {
		start = i
		end = c
	}
	if j <= q {
		start = j
		end = q
	}
	fmt.Printf("start: %d, end: %d\n", start, end)

	for m := start; m <= end; m++ {
		tmp[k] = arr[m]
		k += 1
	}
	fmt.Printf("tmp1: %v\n", tmp)

	for n := 0; n <= q-p; n++ {
		arr[p+n] = tmp[n]
		fmt.Printf("h: %v\n", arr)
	}
}
