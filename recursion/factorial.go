/**
阶乘

数字最小是从 0 开始的
*/
package main

import "fmt"

func main() {
	num := factorial(0)
	fmt.Println(num)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return factorial(n-1) * n
}
