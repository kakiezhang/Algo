/**
序列最长递增子序列长度
*/
package backTrack

type SubqueryMax struct {
	arr []int
	cnt int
}

func newSubqueryMax(arr []int, cnt int) *SubqueryMax {
	return &SubqueryMax{
		arr: arr,
		cnt: cnt,
	}
}

func (sm *SubqueryMax) getMax(i int) int {
	if i == 0 {
		return 1
	}

	var length int

	if sm.arr[i] >= sm.arr[i-1] {
		length = sm.getMax(i-1) + 1
	} else {
		length = sm.getMax(i - 1)
	}

	return length
}
