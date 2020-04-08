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
}

func NewCoinChangeElegant(
	arr []int, cnt int) *CoinChangeElegant {
	return &CoinChangeElegant{
		arr: arr,
		cnt: cnt,
	}
}

func (cce *CoinChangeElegant) getMinNum(amount int) int {
	if amount == 0 {
		return 0
	}

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

	return min
}
