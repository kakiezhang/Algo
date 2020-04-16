package backTrack

import (
	"fmt"
	"testing"
)

func TestCoinChangeElegant(t *testing.T) {
	var arr []int
	var minNum, amount int

	arr = []int{9, 1, 3, 5}
	arr = []int{2, 1, 4, 7}
	// arr = []int{90, 10, 30, 50}
	amount = 9

	cce := NewCoinChangeElegant(arr, len(arr), amount)
	minNum = cce.getMinNum(amount)

	fmt.Println(minNum)
}
