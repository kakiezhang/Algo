/**
希尔排序
步长从 n/2 递减到 1
*/
package sort

func shellSort(arr []int, cnt int) {
	for step := cnt / 2; step >= 1; {
		for i := step; i < cnt; i++ {
			k := i
			j := k - step

			for j >= 0 && j < cnt-step {
				if arr[k] >= arr[j] {
					break
				} else {
					arr[k], arr[j] = arr[j], arr[k]
					k = j
					j = j - step
				}
			}
		}

		step = step / 2
	}
}
