/**
插入排序
分已排和未排序区间，在未排序区间顺序
取一个数，倒序的在已排区间中找出第一
个比其小的值
*/
package sort

func insertionSort(arr []int, cnt int) {
	for i := 1; i < cnt; i++ {
		k := i

		for j := i - 1; j >= 0; j-- {
			if arr[k] >= arr[j] {
				break
			} else {
				arr[k], arr[j] = arr[j], arr[k]
				k = j
			}
		}
	}
}
