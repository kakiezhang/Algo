package backTrack

import (
	"fmt"
	"testing"
)

func TestCoinChangeElegant(t *testing.T) {
	var arr []int
	arr = []int{1, 3, 5}

	cce := NewCoinChangeElegant(arr, len(arr))
	minNum := cce.getMinNum(9)

	fmt.Println(minNum)
}
