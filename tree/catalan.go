/**
给定一组值，计算可以构建多少种二叉树
1, 3, 5, 6, 9, 10
卡特兰数
*/
package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6, 9, 10}
	fmt.Println(catalan(arr, len(arr)-1)) // 按住第一个节点不变
}

func catalan(arr []int, cnt int) int {
	if cnt == 0 || cnt == 1 {
		return 1
	}

	var rs int

	for i := 0; i < cnt; i++ {
		rs += catalan(arr, i) * catalan(arr, cnt-i-1)
	}

	return rs
}
