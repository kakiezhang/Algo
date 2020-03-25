/**
小顶堆和大顶堆
数组元素从第1个位置开始写入
以大顶堆举例
插入：将元素放入数组的最后面，然后从下往上堆化
删除堆顶：将堆顶和最后一个元素做交换，删除最后一个元素，把堆顶的元素从上往下堆化
*/
package main

import "fmt"

func main() {
	var h = newHeap(3)

	h.insert(1)
	h.insert(2)
	fmt.Printf("%+v\n", h)

	h.insert(3)
	fmt.Printf("%+v\n", h)

	h.insert(7)
	fmt.Printf("%+v\n", h)

	h.insert(4)
	fmt.Printf("%+v\n", h)

	h.insert(6)
	fmt.Printf("%+v\n", h)
}

type heap struct {
	arr []int
	cnt int
	max int
}

func newHeap(max int) *heap {
	if max <= 0 {
		panic("max cant be le zero")
	}

	return &heap{
		arr: make([]int, max),
		cnt: 0,
		max: max,
	}
}

func (h *heap) insert(x int) {
	if h.max == h.cnt+1 {
		// expand
		b := make([]int, h.max*2)
		copy(b, h.arr)
		h.arr = b
		h.max = h.max * 2
	}

	// fmt.Printf("hhh: %+v\n", h)

	h.cnt += 1
	h.arr[h.cnt] = x
	h.heapify()
}

func (h *heap) heapify() {
	if h.cnt == 1 {
		return
	}

	pos := h.cnt
	// fmt.Println(h.arr)
	// fmt.Println(pos)

	for {
		if pos <= 1 {
			break
		}

		if h.arr[pos/2] < h.arr[pos] {
			h.arr[pos], h.arr[pos/2] = h.arr[pos/2], h.arr[pos]
			pos = pos / 2
		} else {
			break
		}
	}
}
