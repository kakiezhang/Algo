package dynamicProgramming

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	var arr []int
	var amount int

	arr = []int{1, 3, 5}
	arr = []int{1, 5, 6}
	arr = []int{2, 1, 3}
	amount = 9

	cc := newCoinChange(arr, len(arr), amount)
	num, states, coins := cc.getMin()
	cc.printStates(states)
	fmt.Printf("minNum: %d\n", num)
	fmt.Println("coins: ")
	fmt.Println(coins)
}
