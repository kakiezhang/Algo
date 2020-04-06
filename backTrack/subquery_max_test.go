package backTrack

import (
	"fmt"
	"testing"
)

func TestSubqueryMax(t *testing.T) {
	var arr []int

	arr = []int{2, 9, 3, 6, 5, 1, 7}

	sm := newSubqueryMax(arr, len(arr))

	n := sm.getMax(len(arr) - 1)

	fmt.Println(n)
}
