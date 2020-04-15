/**
两个有序数组合并为一个有序数组
*/
package array

func mergeTwoOrder(arr1, arr2 []int, m, n int) []int {
	// m 是 arr1 的长度，n 是 arr2 的长度
	var rs = make([]int, m+n)
	var k int    // 指向 rs 的指针
	var i, j int // i 指向 arr1，j 指向 arr2

	for {
		if i >= m || j >= n {
			break
		}

		if arr1[i] < arr2[j] {
			rs[k] = arr1[i]
			i += 1
		} else {
			rs[k] = arr2[j]
			j += 1
		}

		k += 1
	}

	for x := i; x < m; x++ {
		rs[k] = arr1[x]
		k += 1
	}
	for y := j; y < n; y++ {
		rs[k] = arr2[y]
		k += 1
	}

	return rs
}
