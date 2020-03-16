/**
bm 算法
假设模式串的字符范围为全 ascii 表 128 位
1. 坏字符
2. 好后缀
*/
package main

import "fmt"

func main() {
	// testGenerateBadCharacter()
	testGenerateGoodSuffix()
	// testGenGS()
}

func bm(a, b string, m, n int) {
	// bc := generateBadCharacter(b, n)
	// suffix, prefix := generateGoodSuffix(b, n)

	// for i := m - n; i > 0; i-- {
	// }
}

func testGenerateBadCharacter() {
	var b string
	var bc []int

	b = "abcdeab"
	bc = generateBadCharacter(b, len(b))
	fmt.Printf("bc: %v\n", bc)
}

func generateBadCharacter(b string, n int) []int {
	// b 是模式串，n 是模式串长度
	var bc = make([]int, 128)

	for i := 0; i < 128; i++ {
		bc[i] = -1
	}

	for j := 0; j < n; j++ {
		x := rune(b[j])
		bc[x] = j
	}

	return bc
}

func testGenerateGoodSuffix() {
	var b string
	var suffix []int
	var prefix []bool

	b = "abcdeab"
	b = "abaabxabcx"
	b = "bbc"
	b = "abcabcabc"
	b = "ababa"
	suffix, prefix = generateGoodSuffix(b, len(b))
	fmt.Printf("suffix: %v\n", suffix)
	fmt.Printf("prefix: %v\n", prefix)
}

func generateGoodSuffix(b string, n int) ([]int, []bool) {
	var prefix = make([]bool, n)
	var suffix = make([]int, n)
	for i := 0; i < n; i++ {
		prefix[i] = false
		suffix[i] = -1
	}

	for i := 0; i < n-1; i++ {
		var j = n - i - 2

		for ; j > 0; j-- {
			if b[n-i-1] == b[j] {
				break
			}
		}

		if b[n-i-1] == b[j] && ((suffix[i] > -1 && i > 0) || (suffix[i] == -1 && i == 0)) {
			// fmt.Println("ssssss")
			// fmt.Println(i)
			// fmt.Println(j)
			suffix[i+1] = j
			if j == 0 {
				prefix[i+1] = true
			}
		}
	}

	return suffix, prefix
}

func testGenGS() {
	var b string
	var suffix []int
	var prefix []bool

	// b = "abcdeab"
	// b = "abaabxabcx"
	// b = "bbc"
	b = "abcabcabc"
	// b = "ababa"
	suffix, prefix = genGS(b, len(b))
	fmt.Printf("suffix: %v\n", suffix)
	fmt.Printf("prefix: %v\n", prefix)
}

func genGS(b string, n int) ([]int, []bool) {
	var prefix = make([]bool, n)
	var suffix = make([]int, n)
	for i := 0; i < n; i++ {
		prefix[i] = false
		suffix[i] = -1
	}

	for i := 0; i < n-1; i++ {
		j := i
		var k int

		for j >= 0 && b[j] == b[n-1-k] {
			k += 1
			suffix[k] = j
			j -= 1
		}

		if j == -1 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}
