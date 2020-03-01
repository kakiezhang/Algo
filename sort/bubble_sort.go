package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	fmt.Println(arr)
	bubble(arr, len(arr))
	fmt.Println(arr)

	arr = []int{6, 5, 4, 1, 2, 3}
	fmt.Println(arr)
	bubble(arr, len(arr))
	fmt.Println(arr)
}

func bubble(arr []int, num int) {
	if num <= 1 {
		return
	}

	var cnt int

	for i := 0; i < num; i++ {
		cnt += 1
		seq := true

		for j := 0; j < num-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp

				seq = false
			}
		}

		if seq {
			break
		}
	}

	fmt.Printf("bubble count: %d\n", cnt)
}
