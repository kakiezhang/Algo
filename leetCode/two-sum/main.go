package main

import "fmt"

func main() {
	var nums []int
	var target int

	nums = []int{2, 7, 11, 15}
	target = 9

	nums = []int{3, 2, 4}
	target = 6

	nums = []int{3, 3}
	target = 6

	rs := twoSum(nums, target)
	fmt.Printf("twoSum(%+v, %d) = %+v\n", nums, target, rs)
}

func twoSum(nums []int, target int) []int {
	rvidx := make(map[int]int)
	rests := make([]int, 0, len(nums))
	for k, v := range nums {
		rvidx[v] = k
		rests = append(rests, target-v)
	}

	var rs []int

	for k, v := range rests {
		if k1, ok := rvidx[v]; ok && k != k1 {
			rs = append(rs, k, k1)
			break
		}
	}

	return rs
}
