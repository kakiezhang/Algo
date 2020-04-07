/**
序列最长递增子序列及其长度
*/
package dynamicProgramming

import "fmt"

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

func (sm *SubqueryMax) getMaxNum() int {
	var states = make([]int, sm.cnt)
	states[0] = 1 // 第一个节点状态长度是1

	for i := 1; i < sm.cnt; i++ {
		states[i] = 1
		for j := 0; j < i; j++ {
			if sm.arr[j] < sm.arr[i] && states[i] < states[j]+1 {
				// 方程：maxlen(i) = max(maxlen(0),..,maxlen(i-1))+1
				states[i] = states[j] + 1
			}
		}
	}

	fmt.Println(states)

	var max int
	for i := 0; i < sm.cnt; i++ {
		if max < states[i] {
			max = states[i]
		}
	}
	return max
}

func (sm *SubqueryMax) getMaxSub() []int {
	sub := make([][]int, sm.cnt)
	for i := 0; i < sm.cnt; i++ {
		sub[i] = make([]int, 0, sm.cnt)
	}
	sub[0] = sub[0][:1]
	sub[0][0] = sm.arr[0]

	for i := 1; i < sm.cnt; i++ {
		for j := 0; j < i; j++ {
			if sm.arr[j] < sm.arr[i] && len(sub[i]) < len(sub[j])+1 {
				fmt.Printf("bef=> i: %d, j: %d, sub: %v, %d, subj: %v\n",
					i, j, sub[i], len(sub[i]), sub[j])

				sub[i] = sub[i][:len(sub[j])]
				copy(sub[i], sub[j])

				fmt.Printf("aft=> i: %d, j: %d, sub: %v, %d, subj: %v\n",
					i, j, sub[i], len(sub[i]), sub[j])
			}
		}
		sub[i] = append(sub[i], sm.arr[i])
	}

	fmt.Println(sub)

	maxSub := sub[0]

	for _, v := range sub {
		if len(v) > len(maxSub) {
			maxSub = v
		}
	}

	return maxSub
}
