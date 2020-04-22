/**
快速排序
利用分区点，将数组分成比其小和比其大两块区域，
然后再各自找分区继续分
*/
package sort

func quickSort(arr []int, cnt int) {
	quickUnit(arr, 0, cnt-1)
}

func quickUnit(arr []int, i, j int) {
	if i >= j {
		return
	}

	p_idx := partition(arr, i, j)
	quickUnit(arr, i, p_idx-1)
	quickUnit(arr, p_idx+1, j)
}

func partition(arr []int, i, j int) int {
	pivot := arr[j]
	m, n := i, i

	for n < j {
		if arr[n] > pivot {
			n += 1
		} else {
			if m != n {
				arr[m], arr[n] = arr[n], arr[m]
			}
			m += 1
			n += 1
		}
	}

	if m != j {
		arr[m], arr[n] = arr[n], arr[m]
	}

	return m
}
