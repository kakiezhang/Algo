/**
归并排序
*/
package sort

func mergeSort(arr []int, cnt int) {
	mergeUnit(arr, 0, cnt-1)
}

func mergeUnit(arr []int, i, j int) {
	if i >= j {
		return
	}

	mid := (i + j) / 2

	mergeUnit(arr, i, mid)
	mergeUnit(arr, mid+1, j)

	mergeTwoInOne(arr, i, j, mid)
}

func mergeTwoInOne(arr []int, i, j, mid int) {
	rs := make([]int, j-i+1)

	m := i
	n := mid + 1
	var k int

	for m <= mid && n <= j {
		if arr[m] <= arr[n] {
			rs[k] = arr[m]
			m += 1
		} else {
			rs[k] = arr[n]
			n += 1
		}
		k += 1
	}

	if m <= mid {
		for x := m; x <= mid; x++ {
			rs[k] = arr[x]
			k += 1
		}
	} else if n <= j {
		for y := n; y <= j; y++ {
			rs[k] = arr[y]
			k += 1
		}
	}

	for z := i; z <= j; z++ {
		arr[z] = rs[z-i]
	}
}
