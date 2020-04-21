/**
冒泡排序
俩俩比较，前比后大，交换
*/
package sort

func bubbleSort(arr []int, cnt int) {
	for i := 0; i < cnt-1; i++ {
		for j := i + 1; j < cnt; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
