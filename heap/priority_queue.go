/**
优先级队列（以小顶堆为例）&合并K个有序数组
元素从数组的第1位开始存储
push: 往队列中插入一个数据
pop: 从队列中取出最小的
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	h := newHeap(11)

	for _, v := range []int{
		30, 48, 73, 35, 25, 65, 60, 10, 70, 50, 20} {
		fmt.Printf("push x: %d\n", v)
		h.push(v)
		layerPrint(h)
	}

	for h.cnt > 0 {
		x := h.pop()
		fmt.Printf("pop x: %d\n", x)
		layerPrint(h)
	}
}

func layerPrint(h *heap) {
	if h.cnt == 0 {
		return
	}

	var i int // 表示深度

	for i = 0; ; i++ {
		j := int(math.Pow(2, float64(i)))   // 每一层的起始点
		k := int(math.Pow(2, float64(i+1))) // 每一层下一层的起始点

		if j > h.cnt {
			break
		}

		fmt.Printf("[depth: %d]", i)

		for ; j < k; j++ {
			if j <= h.cnt {
				fmt.Printf("[%d]", h.arr[j])
			} else {
				fmt.Println()
				goto endLoop
			}
		}

		fmt.Println()
	}
endLoop:
	fmt.Println()
}

type heap struct {
	arr []int
	cnt int
	max int
}

func newHeap(max int) *heap {
	if max == 0 {
		panic("capacity cannot be zero")
	}
	return &heap{
		arr: make([]int, max+1),
		max: max + 1, // 把第0位空出来
	}
}

func (h *heap) push(x int) {
	if h.cnt == h.max-1 {
		panic("heap is full")
	}

	h.cnt += 1
	h.arr[h.cnt] = x

	if h.cnt == 1 {
		return
	}

	h.heapifyFromBottom(h.cnt)
}

func (h *heap) heapifyFromBottom(start int) {
	for {
		if start == 1 {
			break
		}
		j := start / 2

		if h.arr[start] < h.arr[j] {
			h.arr[start], h.arr[j] = h.arr[j], h.arr[start]
			start = j
		} else {
			break
		}
	}
}

func (h *heap) heapifyFromTop(start, end int) {
	for {
		min := start
		j := start * 2
		if j > end {
			break
		}

		if h.arr[start] > h.arr[j] {
			min = j
		}

		if j+1 <= end && h.arr[min] > h.arr[j+1] {
			min = j + 1
		}

		if min != start {
			h.arr[start], h.arr[min] = h.arr[min], h.arr[start]
			start = min
		} else {
			break
		}
	}
}

func (h *heap) pop() int {
	if h.cnt == 0 {
		panic("heap is empty")
	} else if h.cnt == 1 {
		h.cnt -= 1
		return h.arr[1]
	}

	x := h.arr[1]
	h.arr[1], h.arr[h.cnt] = h.arr[h.cnt], h.arr[1]
	h.arr[h.cnt] = 0
	h.cnt -= 1
	h.heapifyFromTop(1, h.cnt)

	return x
}
