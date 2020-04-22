/**
找第K大数
借用快排分区和二分查找比大小的思想
*/
package sort

func findKBigNum(arr []int, cnt, k int) int {
	if k <= 0 {
		panic("kth. num cannot le 0.")
	}
	if k > cnt {
		panic("kth. num cannot gt arr capacity.")
	}

	var p_idx int
	i := 0
	j := cnt - 1

	for {
		p_idx = partBigToSmall(arr, i, j)
		if p_idx == k-1 {
			return arr[p_idx]
		} else if p_idx < k-1 {
			i = p_idx + 1
		} else {
			j = p_idx - 1
		}
	}
}

func partBigToSmall(arr []int, i, j int) int {
	pivot := arr[j]
	m, n := i, i

	for n < j {
		if arr[n] < pivot {
			n += 1
		} else {
			if m != n {
				arr[m], arr[n] = arr[n], arr[m]
			}
			m += 1
			n += 1
		}
	}

	if m != n {
		arr[m], arr[n] = arr[n], arr[m]
	}

	return m
}
