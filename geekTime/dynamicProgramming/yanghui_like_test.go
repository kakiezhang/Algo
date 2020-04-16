package dynamicProgramming

import (
	"fmt"
	"testing"
)

func TestNewYanghui(t *testing.T) {
	var arr []int
	arr = []int{5, 7, 8, 2, 3, 4, 4, 9, 6, 1, 2, 7, 9, 4, 5}
	arr = []int{5, 7, 8, 5, 3, 4, 12, 9, 9, 10, 2, 7, 9, 4, 5}
	yh := newYanghui(arr, 5)
	fmt.Printf("Yanghui: %+v\n", yh)

	fmt.Println("Origin states: ")
	yh.PrintTriangle()

	idx, path := yh.FindMinPath()
	fmt.Println("After states: ")
	yh.PrintTriangle()
	fmt.Printf("idx: %v\n", idx)
	fmt.Printf("path: %v\n", path)
}
