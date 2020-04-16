package backTrack

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	var arr []int
	arr = []int{1, 3, 5}

	cc := NewCoinChange(arr, len(arr), 9)
	cc.find(0, 0, 0)

	fmt.Println(cc.minNum)
}
