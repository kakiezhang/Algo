package main

import "fmt"

func main() {
	{
		s := "abcd"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "abcabcbb"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "bbbbb"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "pwwkew"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "aab"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "ca"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "aa"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "dvdf"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}

	{
		s := "dvgcvgdf"
		fmt.Printf("lengthOfLongestSubstring(%s) = %d\n", s, lengthOfLongestSubstring(s))
	}
}

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	cnt := len(s)
	if cnt <= 1 {
		return cnt
	}

	m := make(map[byte]int) // key: element value: index
	m[s[0]] = 0
	rs := len(m)

	// 保持找到的最新窗口的左侧，移动窗口右侧进行比较
	// 每次找到重复元素时，左侧往右收缩
	h := 0
	i := 0
	j := i + 1
	for i < len(s)-1 && j < len(s) {
		next := s[j]

		if k, ok := m[next]; ok {
			i = k + 1

			for a := h; a <= k; a++ {
				delete(m, s[a])
			}

			h = i
		}

		m[next] = j
		rs = max(len(m), rs)

		j++

		// fmt.Printf("h: %d, i: %d, j: %d, rs: %d, m: %+v\n", h, i, j, rs, m)
	}
	return rs
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
