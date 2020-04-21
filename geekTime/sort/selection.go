/**
选择排序
分已排和未排序区间，在未排区间
找一个最小值放入已排的最后
*/
package sort

func selectionSort(arr []int, cnt int) {
	for i := 0; i < cnt-1; i++ {
		min := arr[i]
		k := i // min idx

		for j := i + 1; j < cnt; j++ {
			if min > arr[j] {
				min = arr[j]
				k = j
			}
		}

		if k != i {
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
}
