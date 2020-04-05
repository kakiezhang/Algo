package dynamicProgramming

import (
	"fmt"
	"testing"
)

func TestNewYanghui(t *testing.T) {
	arr := []int{5, 7, 8, 2, 3, 4, 4, 9, 6, 1, 2, 7, 9, 4, 5}
	yh := newYanghui(arr, 5)
	fmt.Printf("%+v\n", yh)

	minSum := yh.FindMinPath()
	fmt.Printf("%+v\n", yh.states)
	fmt.Printf("minSum: %d\n", minSum)
}
