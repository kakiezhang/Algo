/**
求众数
即出现次数大于 n/2 的数
*/
package majorityElement

import "sort"

func MajorNumMemo(arr []int, cnt int) int {
	var mem = make(map[int]int)
	var rs int

	for i := 0; i < cnt; i++ {
		x := arr[i]

		if _, ok := mem[x]; ok {
			mem[x] += 1
			// fmt.Println(mem)

			if mem[x]*2 > cnt {
				rs = x
				break
			}
		} else {
			mem[x] = 1
		}
	}

	return rs
}

func MajorNumMiddle(arr []int, cnt int) int {
	// 前提：假设一个数组内，总是存在众数
	sort.Ints(arr)
	return arr[cnt/2]
}
