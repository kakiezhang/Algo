/**
bm 算法
假设模式串的字符范围为全 ascii 表 128 位
1. 坏字符
2. 好后缀
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// testGenerateBadCharacter()
	testGenerateGoodSuffix()
	// testGenGS()
}

func bm(a, b string, m, n int) int {
	// a 是主串，m 是主串长度
	// b 是模式串，n 是模式串长度
	bc := generateBadCharacter(b, n)
	suffix, prefix := generateGoodSuffix(b, n)

	i := 0
	for i < m-n {

		num := 0   // 好后缀都长度
		j := n - 1 // 坏字符在模式串中的位置
		k := i + n - 1
		for j >= 0 && b[j] == a[k] {
			j--
			k--
			num++
		}

		if j < 0 { // 所有都匹配上
			return i
		}

		var paceNo int

		if num > 0 && suffix[num] > -1 { // 有后缀子串匹配上
			paceNo = j - suffix[num] + 1
		} else {
			// 没有后缀子串，去找一下后缀当中有没有可匹配的前缀子串
			x := 0
			y := i + n - num + 1 // 找到好后缀的后缀子串的在主串中第一个点的下标
			for x < n && b[x] == a[y] {
				x++
				y++
				if prefix[x-1] {
					break
				}
			}

			if x < n { // 找到了对应的前缀
				paceNo = y
			} else { // 没有找到了，直接滑倒最后一位
				paceNo = n
			}
		}

		paceNo = int(math.Max(float64(paceNo), float64(bc[j]+1)))
		i += paceNo
	}

	return -1
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
