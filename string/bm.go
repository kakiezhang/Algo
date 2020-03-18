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
	// testGenerateGoodSuffix()
	// testGenGS()
	testBm()
}

func testBm() {
	var a, b string
	var pos int

	// // test 1st
	// a = "addkaamask"
	// b = "ama"
	// pos = bm(a, b, len(a), len(b))
	// fmt.Printf("pos: %d\n", pos)

	// // test 2nd
	// a = "aaoahemaoabemao"
	// b = "aoabemao"
	// pos = bm(a, b, len(a), len(b))
	// fmt.Printf("pos: %d\n", pos)

	// // test 2nd
	// a = "eaohoabemao"
	// b = "mao"
	// pos = bm(a, b, len(a), len(b))
	// fmt.Printf("pos: %d\n", pos)

	a = "heaohoabemao"
	b = "mmao"
	pos = bm(a, b, len(a), len(b))
	fmt.Printf("pos: %d\n", pos)
}

func bm(a, b string, m, n int) int {
	// a 是主串，m 是主串长度
	// b 是模式串，n 是模式串长度
	bc := generateBadCharacter(b, n)
	fmt.Printf("bc: %v\n", bc)
	suffix, prefix := genGS(b, n)
	fmt.Printf("suffix: %v\n", suffix)
	fmt.Printf("prefix: %v\n", prefix)

	i := 0
	for i <= m-n {
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

		var paceNo = -1

		fmt.Printf("num: num[%d]\n", num)

		if num > 0 { // 有好后缀
			if suffix[num] > -1 { // 有后缀子串匹配上
				fmt.Printf("1st: match child tail: bc[%d] sf[%d]\n", j, suffix[num])
				paceNo = j - suffix[num] + 1
			} else {
				// 没有后缀子串，去找一下好后缀当中有没有可匹配的前缀子串
				x := -1
				y := i + n - num + 1 // 找到好后缀的后缀子串的在主串中第一个点的下标

				fmt.Printf("2nd: first child tail index of GS in main str: %d\n", y)
				for x < n && b[x+1] == a[y] {
					x++
					y++
					if prefix[x] {
						break
					}
				}

				fmt.Printf("2nd: x[%d], y[%d]\n", x, y)
				if x < n && x > -1 { // 找到了对应的前缀
					paceNo = y
				} else {
					// 没有找到，直接滑倒最后一位
					fmt.Println("3rd: direct move to next seg")
					paceNo = n
				}
			}
		}

		fmt.Printf("paceNo: %v\n", paceNo)
		fmt.Printf("j-bc[a[i+j]]: %v\n", j-bc[a[i+j]])

		paceNo = int(math.Max(float64(paceNo), float64(j-bc[a[i+j]])))
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
