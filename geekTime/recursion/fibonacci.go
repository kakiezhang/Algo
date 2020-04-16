/**
Fibonacci

f(n)=f(n-1)+f(n-2)
可以用散列表来判断元素是否已经存在，避免重复计算

*/
package main

import "fmt"

var hs map[int]int

func main() {
	hs = make(map[int]int)
	num := fib(1)
	fmt.Println(num)
}

func fib(n int) int {
	var v int
	var ok bool
	if v, ok = hs[n]; ok {
		return v
	}

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	v = fib(n-1) + fib(n-2)
	hs[n] = v
	return v
}
