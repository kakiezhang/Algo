/**
0-1背包（物品不可分割）
一共有n个物品，背包最大承载maxWkg，要求在
满足背包能承受的最大重量情况下，保证物品价值
最高
*/
package backTrack

import "fmt"

type MostWeightValue struct {
	weight   []int // 物品重量
	value    []int // 物品价值
	num      int   // 物品个数
	packMaxW int   // 背包最大承重
	goodMaxW int   // 背包内装的物品最大重量
	goodMaxV int   // 背包内装的物品最高价值
}

func newMostWeightValue(num, packMaxW int,
	weight, value []int) *MostWeightValue {
	return &MostWeightValue{
		weight:   weight,
		value:    value,
		num:      num,
		packMaxW: packMaxW,
		goodMaxW: 0,
		goodMaxV: 0,
	}
}

func (mwv *MostWeightValue) recurFind(i, cw, cv int) {
	fmt.Printf("i: %d, cw: %d, cv: %d\n", i, cw, cv)
	if i == mwv.num || cw == mwv.packMaxW { // 物品放完或者承重最大
		if cw > mwv.goodMaxW {
			mwv.goodMaxW = cw
		}
		if cv > mwv.goodMaxV {
			mwv.goodMaxV = cv
		}
		return
	}

	mwv.recurFind(i+1, cw, cv) // 不装第i个物品
	if cw+mwv.weight[i] <= mwv.packMaxW {
		mwv.recurFind(i+1, cw+mwv.weight[i], cv+mwv.value[i]) // 装第i个物品
	}
}
