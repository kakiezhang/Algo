/**
硬币找零
给定一组面值: c1, c2, c3（元）
求支付w元，最少需要多少个硬币
minCoin(w)=min(minCoin(w-c1),..minCoin(w-cn))+1
*/
package backTrack

import (
	"math"
)

type CoinChangeElegant struct {
	arr []int // 面值
	cnt int
	mem []int // 备忘录(存储每一个可达金额需要用到的最少面值个数)
}

func NewCoinChangeElegant(
	arr []int, cnt, amount int) *CoinChangeElegant {
	return &CoinChangeElegant{
		arr: arr,
		cnt: cnt,
		mem: make([]int, amount+1), // 从第1位开始启用
	}
}

func (cce *CoinChangeElegant) getMinNum(amount int) int {
	if amount == 0 {
		return 0
	}

	if cce.mem[amount] > 0 {
		return cce.mem[amount]
	}

	// fmt.Printf("bef=> amount: %d\n", amount)

	min := math.MaxInt64

	for k := 0; k < cce.cnt; k++ {
		if amount-cce.arr[k] < 0 {
			continue
		}
		n := cce.getMinNum(amount-cce.arr[k]) + 1
		if min > n {
			min = n
		}
	}

	// fmt.Printf("aft=> amount: %d, num: %d\n", amount, min)

	cce.mem[amount] = min
	return min
}
