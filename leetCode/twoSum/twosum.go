/**
两数之和
*/
package twoSum

func TwoSumBruteForce(arr []int, cnt, target int) [][2]int {
	var rs = make([][2]int, 0)

	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if arr[i]+arr[j] == target {
				rs = append(rs, [2]int{i, j})
			}
		}
	}

	return rs
}

func TwoSumMemory(arr []int, cnt, target int) [][2]int {
	var mem = make(map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		mem[arr[i]] = i // 倒置备忘录 k: 值，v: 索引
	}

	// fmt.Println(mem)

	var rs = make([][2]int, 0)
	var uni = make(map[int]int) // 去重, 更好的做法可以用 bitmap

	for i := 0; i < cnt; i++ {
		j := target - arr[i]

		if v, ok := mem[j]; ok && v != i {
			min, max := i, v
			if min > max {
				min, max = max, min
			}

			k := max*10 + min

			if _, ok := uni[k]; !ok {
				uni[k] = k
				rs = append(rs, [2]int{i, v})
			}
		}
	}

	return rs
}
