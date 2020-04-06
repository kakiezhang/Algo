/**
硬币找零
给定一组面值: c1, c2, c3（元）
求支付w元，最少需要多少个硬币
*/
package backTrack

import (
	"math"
)

type CoinChange struct {
	arr      []int // 面值
	cnt      int
	amount   int // 需要支付多少元
	minNum   int // 用的最少的硬币个数
	maxTimes int // 估算一个最大次数，超过这个次数就停止
}

func NewCoinChange(arr []int, cnt, amount int) *CoinChange {
	return &CoinChange{
		arr:      arr,
		cnt:      cnt,
		amount:   amount,
		minNum:   math.MaxInt64,
		maxTimes: 10,
	}
}

func (cc *CoinChange) find(i, j, sum int) {
	// i是第几次，也就是用掉第几个硬币，
	// j是在第i次之前加上的面值，
	// sum是到第i次之前已经达到的支付金额总数
	if cc.amount == sum {
		if cc.minNum > i {
			cc.minNum = i
		}
		return
	}

	if i >= cc.maxTimes {
		return
	}

	for k := 0; k < cc.cnt; k++ {
		cc.find(i+1, cc.arr[k], sum+cc.arr[k])
	}
}
