package backTrack

import (
	"fmt"
	"testing"
)

func TestMostWeightValue(t *testing.T) {
	mwv := newMostWeightValue(8, 20, []int{
		3, 8, 6, 8, 1, 3, 5, 4,
	}, []int{
		25, 31, 11, 15, 22, 14, 12, 10,
	})
	// fmt.Printf("mwv: %+v\n", mwv)
	mwv.recurFind(0, 0, 0)
	fmt.Printf("goodMaxW: %d\n", mwv.goodMaxW)
	fmt.Printf("goodMaxV: %d\n", mwv.goodMaxV)
}
