package main

import "fmt"

func main() {
	var arr []int

	arr = []int{8, 2, 4, 5}
	fullArrange(arr, 0, len(arr)-1)
}

func fullArrange(arr []int, p, q int) {
	if p == q {
		for j := 0; j <= q; j++ {
			fmt.Printf("%d ", arr[j])
		}
		fmt.Println()
		return
	}

	for i := p; i <= q; i++ {
		swap(&arr[i], &arr[p])
		fullArrange(arr, p+1, q)
		swap(&arr[i], &arr[p])
	}
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
