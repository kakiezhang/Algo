package twoSum

import (
	"testing"
)

func TestTwoSumBruteForce(t *testing.T) {
	var arr []int
	var target int
	var rs [][2]int

	arr = []int{2, 7, 6, 15, 3}
	target = 9

	rs = TwoSumBruteForce(arr, len(arr), target)
	t.Logf("arr: %v, rs: %v", arr, rs)
}

func TestTwoSumMemory(t *testing.T) {
	var arr []int
	var target int
	var rs [][2]int

	arr = []int{2, 7, 6, 15, 3}
	target = 9

	rs = TwoSumMemory(arr, len(arr), target)
	t.Logf("arr: %v, rs: %v", arr, rs)
}
