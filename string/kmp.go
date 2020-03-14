/**
kmp 算法
按照好前缀来求模式串是否存在于主串
*/
package main

import (
	"fmt"
)

func main() {
	testKmp()
	// testGetNexts()
}

func testKmp() {
	var a, b string
	var pos int

	a = "abaabcaabaabcax"
	// a = "abaabdaabaabcax"
	b = "abaabc"
	pos = kmp(a, b, len(a), len(b))
	fmt.Printf("pos: %d\n", pos)
}

func kmp(a, b string, m, n int) int {
	// a 是主串，m 是其长度，b 是模式串，n 是其长度
	next := getNexts(b, n)
	fmt.Printf("next: %v\n", next)
	var j int

	for i := 0; i < m; i++ {
		for j > 0 && a[i] != b[j] {
			j = next[j-1] + 1
		}

		if a[i] == b[j] {
			j += 1
		}

		if j == n {
			return i - j + 1
		}
	}

	return -1
}

func getNexts(b string, n int) []int {
	var next = make([]int, n)
	next[0] = -1
	var j = -1

	for i := 1; i < n; i++ {
		for j > -1 && b[j+1] != b[i] {
			fmt.Printf("2: b[%d]: %v, b[%d]: %d\n", i, b[i], j, b[j])
			j = next[j]
		}

		fmt.Printf("1: b[%d]: %v, b[%d]: %d\n", i, b[i], j+1, b[j+1])
		if b[j+1] == b[i] {
			j += 1
		}
		next[i] = j
	}

	return next
}

func testGetNexts() {
	var b string
	var next []int

	b = "abaabc"
	b = "abaabcaabaabcax"
	next = getNexts(b, len(b))
	fmt.Printf("next: %v\n", next)
}
