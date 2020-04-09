package dynamicProgramming

import "testing"

func TestCoinChange(t *testing.T) {
	var arr []int
	var amount int

	arr = []int{1, 3, 5}
	amount = 9

	cc := newCoinChange(arr, len(arr), amount)
	states := cc.getMin()
	cc.printStates(states)
}
