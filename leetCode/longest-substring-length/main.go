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
}

func lengthOfLongestSubstring(s string) int {
	cnt := len(s)
	if cnt <= 1 {
		return cnt
	}

	m := make(map[byte]int)
	m[s[0]] = 0
	rs := len(m)

	i := 0
	j := i + 1
	for i < len(s)-1 && j < len(s) {
		next := s[j]

		if k, ok := m[next]; ok {
			i = k + 1

			m = make(map[byte]int)

			for a := k; a < j; a++ {
				m[s[a]] = a
			}
		}

		m[next] = j
		rs = max(len(m), rs)

		j++

		// fmt.Printf("i: %d, j: %d, rs: %d, m: %+v\n", i, j, rs, m)
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
