/**
三数之和
1. 排序
2. 固定一个值，然后从两边夹逼
*/
package threeSum

import "sort"

func ThreeSumCornered(arr []int, cnt, target int) [][3]int {
	sort.Ints(arr)

	if cnt < 3 {
		panic("couldn't find sum equal the target.")
	}

	var rs [][3]int

	for i := 0; i < cnt-2; i++ {
		if arr[i] > target {
			return rs
		}
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left := i + 1
		right := cnt - 1

		for left < right {
			total := arr[i] + arr[left] + arr[right]

			if total == target {
				rs = append(rs, [3]int{arr[i], arr[left], arr[right]})

				for left < right && arr[left] == arr[left+1] {
					left = left + 1
				}
				for left < right && arr[right] == arr[right-1] {
					right = right - 1
				}

				left = left + 1
				right = right - 1
			} else if total > target {
				right = right - 1
			} else if total < target {
				left = left + 1
			}
		}
	}

	return rs
}
