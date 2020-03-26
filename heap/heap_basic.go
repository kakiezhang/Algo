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

	for _, v := range []int{1, 2, 3, 7, 5, 4, 6} {
		h.insert(v)
		fmt.Printf("insert: %v, heap: %+v\n", v, h)
	}

	for h.cnt > 0 {
		h.deleteTop()
		fmt.Printf("deleteTop, heap: %+v\n", h)
	}
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
		// auto expand
		h.max = h.max*2 + 1
		b := make([]int, h.max)
		copy(b, h.arr)
		h.arr = b
	}

	// fmt.Printf("hhh: %+v\n", h)

	h.cnt += 1
	h.arr[h.cnt] = x
	h.heapifyFromBottom()
}

func (h *heap) heapifyFromBottom() {
	// 从下往上堆化
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

func (h *heap) deleteTop() {
	if h.cnt == 0 {
		return
	} else if h.cnt == 1 {
		h.arr[h.cnt] = 0
		h.cnt -= 1
		return
	}

	pos := h.cnt
	h.arr[pos], h.arr[1] = h.arr[1], h.arr[pos]
	h.arr[pos] = 0

	if h.cnt == (h.max-1)/2 {
		// auto shrink
		b := make([]int, h.cnt)
		copy(b, h.arr)
		h.arr = b
		h.max = h.cnt
	}

	h.cnt -= 1

	h.heapifyFromTop()
}

func (h *heap) heapifyFromTop() {
	// 从上往下堆化
	if h.cnt == 1 {
		return
	}

	pos := 1

	for {
		i := pos * 2
		if i > h.cnt {
			break
		}

		if i < h.cnt && h.arr[i] < h.arr[i+1] {
			i = i + 1
		}

		if h.arr[pos] < h.arr[i] {
			h.arr[pos], h.arr[i] = h.arr[i], h.arr[pos]
			pos = i
		} else {
			break
		}
	}
}
