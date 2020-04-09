/**
硬币找零
状态节点为(amount+1)^2矩阵
一维：次数
二维：累加金额
从第1位开始用
*/
package dynamicProgramming

import "fmt"

type CoinChange struct {
	arr    []int // 面值（元）
	cnt    int   // 硬币个数
	amount int   // 需要支付的金额
}

func newCoinChange(
	arr []int, cnt, amount int) *CoinChange {
	return &CoinChange{
		arr:    arr,
		cnt:    cnt,
		amount: amount,
	}
}

func (cc *CoinChange) printStates(states [][]bool) {
	for i := 1; i < cc.amount; i++ {
		fmt.Printf("%v\n", states[i])
	}
}

func (cc *CoinChange) getMin() [][]bool {
	var states = make([][]bool, cc.amount+1)
	for i := 0; i < cc.amount+1; i++ {
		states[i] = make([]bool, cc.amount+1)
	}
	for i := 0; i < cc.cnt; i++ {
		states[1][cc.arr[i]] = true
	}

	for i := 2; i < cc.amount+1; i++ {
		for j := 1; j < cc.amount+1; j++ {
			if !states[i-1][j] {
				continue
			}

			for k := 0; k < cc.cnt; k++ {
				if j+cc.arr[k] <= cc.amount && !states[i][j+cc.arr[k]] {
					states[i][j+cc.arr[k]] = true

					if j+cc.arr[k] == cc.amount {
						goto LoopEnd
					}
				}
			}
		}
	}
LoopEnd:

	return states
}
