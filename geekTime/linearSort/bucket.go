/**
桶排序

需要排序的数据存在天然的顺序
比如：给定 10 个考生的考试成绩，求给这 10 个考生的考
试成绩排序（考试分为 3-10 分）

上面的情况就可以拿分数（从 3 - 10 分）作为 8 个桶，
然后分别往桶里面插入符合要求的


*/
package main

import "fmt"

func main() {
	var arr []int

	arr = []int{9, 7, 5, 10, 6, 8, 3, 5, 5, 7}
	low := 0
	high := 10
	rs := bucketCjd(arr, len(arr), low, high)
	fmt.Println(rs)

}

func bucketCjd(arr []int, num, low, high int) []int {
	var rs = make([]int, high-low+1)
	for i := 0; i <= num-1; i++ {
		rs[arr[i]-low] += 1
	}
	return rs
}
