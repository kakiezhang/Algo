/**
堆排序
这里以从小到大排序为例
并且从 0index 开始存储元素以减少数组来回 copy
parent = (i-1)/2, left = 2*i+1, right = 2*i+2
1. 建堆（大顶堆）
有两种方式：
第一种是从前往后遍历数组，从下往上堆化
第二种是从倒数第二层的非叶节点开始，从后往前遍历数组，从上往下堆化
2. 排序
类似于删除元素，从堆顶开始，和最后一个元素做交换，cnt-1，从上往下堆化
*/
package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3}
	h := newHeap(arr, 3)
	h.arr[0] = 6
	fmt.Println(h)
	fmt.Println(arr)
}

type heap struct {
	arr []int
	cnt int
	// max int // 这里基于原地排序，不需要考虑最大容量
}

func newHeap(arr []int, cnt int) *heap {
	return &heap{
		arr: arr,
		cnt: cnt,
	}
}

func (h *heap) heapifyFromBottom() {
	if h.cnt == 1 {
		return
	}

	i := 1

	for {
		if h.arr[i] > h.arr[(i-1)/2] {

		}
	}
}
