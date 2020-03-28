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
	// testPriorityQueue()
	testMergeKSorts()
}

func testMergeKSorts() {
	var arr = [][]int{
		[]int{5, 6, 78, 83, 91},
		[]int{3},
		[]int{7, 15, 29, 37, 73},
		[]int{1, 93},
		[]int{3, 12, 22, 24, 42, 56},
		[]int{8, 9, 66, 91},
	}
	var total int
	var arrCnt = len(arr)
	var eleMaxCnt = make([]int, arrCnt)

	for i := 0; i < len(arr); i++ {
		for j := 0; j <= len(arr[i])-1; j++ {
			total += 1
			eleMaxCnt[i] += 1
		}
	}

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("total: %d\n", total)
	fmt.Printf("eleMaxCnt: %v\n", eleMaxCnt)

	mergedRes := mergeKnumSorts(arr, eleMaxCnt, total, arrCnt)
	// fmt.Printf("mergedRes: %+v\n", mergedRes)

	fmt.Print("mergedRes: \n")
	for _, v := range mergedRes {
		// fmt.Printf("%d ", v.key)
		fmt.Printf("v:%d,q:%d\n", v.key, v.value)
	}
	fmt.Println()
}

func mergeKnumSorts(arr [][]int, eleMaxCnt []int,
	total, arrCnt int) []*heapNode {
	// 数组之间可能会存在不同长度
	var eleUsedCnt = make([]int, arrCnt)
	var mergedRes = make([]*heapNode, total)

	h := newHeap(arrCnt)

	var i int // 表示从哪条数组取
	var k int
	for {
		for ; h.cnt < h.max-1 && i < arrCnt; i++ {
			// 堆中还有剩余空间(没有堆满)并且i没超过数组最大条数，就可以继续
			// fmt.Printf("i: %d\n", i)
			// fmt.Printf("arrCnt: %d\n", arrCnt)
			// fmt.Printf("total: %d\n", total)
			if total > 0 { // arr中还有元素存在
				j := eleUsedCnt[i]
				if j >= eleMaxCnt[i] { // 超过了可取的同排最大索引
					continue
				}
				h.push(newHeapNode(arr[i][j], i))
				eleUsedCnt[i] += 1
				total -= 1
			}
		}

		if total > 0 { // arr中还有元素存在
			node := h.pop()
			i = node.value.(int) // 下次还从这个取出的元素所在的数组取
			mergedRes[k] = node
			k += 1
		} else { // arr中没有元素了，把剩下堆中数据pop完
			for h.cnt > 0 {
				node := h.pop()
				mergedRes[k] = node
				k += 1
			}
			break
		}
	}

	return mergedRes
}

func testPriorityQueue() {
	h := newHeap(11)

	for _, v := range []int{
		30, 48, 73, 35, 25, 65, 60, 10, 70, 50, 20} {
		fmt.Printf("push x: %d\n", v)
		h.push(newHeapNode(v, nil))
		layerPrint(h)
	}

	for h.cnt > 0 {
		node := h.pop()
		fmt.Printf("pop x: %d\n", node.key)
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
				fmt.Printf("[%v]", h.arr[j])
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
	arr []*heapNode
	cnt int
	max int
}

type heapNode struct {
	key   int
	value interface{}
}

func newHeap(max int) *heap {
	if max == 0 {
		panic("capacity cannot be zero")
	}
	return &heap{
		arr: make([]*heapNode, max+1),
		max: max + 1, // 把第0位空出来
	}
}

func newHeapNode(k int, v interface{}) *heapNode {
	return &heapNode{
		key:   k,
		value: v,
	}
}

func (h *heap) push(node *heapNode) {
	if h.cnt == h.max-1 {
		panic("heap is full")
	}

	h.cnt += 1
	h.arr[h.cnt] = node

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

		if h.arr[start].key < h.arr[j].key {
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

		if h.arr[start].key > h.arr[j].key {
			min = j
		}

		if j+1 <= end && h.arr[min].key > h.arr[j+1].key {
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

func (h *heap) pop() *heapNode {
	if h.cnt == 0 {
		panic("heap is empty")
	} else if h.cnt == 1 {
		h.cnt -= 1
		return h.arr[1]
	}

	node := h.arr[1]
	h.arr[1], h.arr[h.cnt] = h.arr[h.cnt], h.arr[1]
	h.arr[h.cnt] = nil
	h.cnt -= 1
	h.heapifyFromTop(1, h.cnt)

	return node
}
