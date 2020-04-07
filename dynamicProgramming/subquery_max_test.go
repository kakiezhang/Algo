package dynamicProgramming

import (
	"fmt"
	"testing"
)

func TestSubqueryMax(t *testing.T) {
	var arr []int

	arr = []int{2, 9, 3, 6, 5, 1, 7}
	arr = []int{2, 1, 2, 3, 5, 1, 7}
	arr = []int{10, 22, 9, 33, 21, 50, 41, 60, 80}

	sm := newSubqueryMax(arr, len(arr))

	max := sm.getMaxNum()
	fmt.Printf("max: %d\n", max)

	maxSub := sm.getMaxSub()
	fmt.Printf("maxSub: %d\n", maxSub)
}
