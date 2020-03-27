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

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{10, 20, 30, 40, 70, 80, 60}
	// var arr = []int{10, 20, 30, 40}
	h := newHeap(arr, len(arr))
	// h.heapifyFromBottom()
	h.heapifyFromTop()
	layerPrint(h)
	// fmt.Println(arr)
}

func layerPrint(h *heap) {
	if h.cnt == 0 {
		return
	}

	var i int // 表示深度

	for i = 0; ; i++ {
		j := int(math.Pow(2, float64(i))) - 1   // 每一层的起始点
		k := int(math.Pow(2, float64(i+1))) - 1 // 每一层下一层的起始点

		if j >= h.cnt {
			break
		}

		fmt.Printf("[depth: %d]", i)

		for ; j < k; j++ {
			if j <= h.cnt {
				fmt.Printf("[%d]", h.arr[j])
			} else {
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
	// max int // 这里基于原地排序，不需要考虑最大容量
}

func newHeap(arr []int, cnt int) *heap {
	return &heap{
		arr: arr,
		cnt: cnt,
	}
}

func (h *heap) heapifyFromBottom() {
	// 第一种的堆化：从前往后遍历数组，从下往上堆化
	if h.cnt == 1 {
		return
	}
	// fmt.Println(h)

	i := 1

	for i < h.cnt {
		j := i

		for {
			if j <= 0 {
				break
			}

			k := (j - 1) / 2
			if h.arr[j] > h.arr[k] {
				h.arr[j], h.arr[k] = h.arr[k], h.arr[j]
				j = k
			} else {
				break
			}
		}

		// fmt.Println(h)

		i += 1
	}
}

func (h *heap) heapifyFromTop() {
	// 第二种的堆化：是从倒数第二层的非叶节点开始，
	// 从后往前遍历数组，从上往下堆化
	if h.cnt == 1 {
		return
	}

	i := (h.cnt - 2) / 2

	for i >= 0 {
		k := i

		for {
			j := k*2 + 1
			if j >= h.cnt {
				break
			}

			max := k

			if h.arr[k] < h.arr[j] {
				max = j
			}

			if j+1 < h.cnt && h.arr[max] < h.arr[j+1] {
				max = j + 1
			}

			// fmt.Println(h.arr)

			if max != k {
				h.arr[k], h.arr[max] = h.arr[max], h.arr[k]
				k = max
			} else {
				break
			}
		}

		i -= 1
	}
}

// 0
// 12
// 3456

// cnt j i
// 7 6 2
// 6 5 2
// 5 4 1
// 4 3 1
// 3 2 0

func (h *heap) sort() {
}
