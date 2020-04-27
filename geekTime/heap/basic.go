/**
二叉堆
小顶堆，大顶堆
idx 从 1 开始计算
score 是作为建堆的依据，拼接在 value 内
*/
package heap

type heaper interface {
	scorer
	compareOperater
}

type scorer interface {
	ScoreFunc(*HeapNode) int
}

type compareOperater interface {
	opCmpFunc(*HeapNode, *HeapNode) bool
}

func NewSmallTopHeap(max int) *SmallTopHeap {
	sth := &SmallTopHeap{
		Heap: NewHeap(max),
	}
	sth.Heap.opCmpFunc = sth.opCmpFunc
	return sth
}

type SmallTopHeap struct {
	*Heap
}

func (sth *SmallTopHeap) opCmpFunc(c, p *HeapNode) bool {
	if sth.ScoreFunc(c) < sth.ScoreFunc(p) {
		return true
	} else {
		return false
	}
}

func NewBigTopHeap(max int) *BigTopHeap {
	bth := &BigTopHeap{
		Heap: NewHeap(max),
	}
	bth.Heap.opCmpFunc = bth.opCmpFunc
	return bth
}

type BigTopHeap struct {
	*Heap
}

func (bth *BigTopHeap) opCmpFunc(c, p *HeapNode) bool {
	if bth.ScoreFunc(c) > bth.ScoreFunc(p) {
		return true
	} else {
		return false
	}
}

func NewHeap(max int) *Heap {
	return &Heap{
		node: make([]*HeapNode, max+1), // 0idx不占位
		cnt:  0,
		max:  max,
	}
}

type Heap struct {
	node []*HeapNode
	cnt  int
	max  int
	opCmpFunc
	ScoreFunc
}

type opCmpFunc func(*HeapNode, *HeapNode) bool
type ScoreFunc func(*HeapNode) int

func NewHeapNode(
	key string, value interface{}) *HeapNode {
	return &HeapNode{
		Key:   key,
		Value: value,
	}
}

type HeapNode struct {
	Key   string
	Value interface{}
}

func (h *Heap) Add(hn *HeapNode) {
	if h.cnt == h.max {
		panic("reach max, couldn't add more node.")
	}

	h.cnt++
	h.node[h.cnt] = hn

	i := h.cnt

	for {
		if i == 1 {
			break
		}

		if h.opCmpFunc(h.node[i], h.node[i/2]) {
			h.node[i], h.node[i/2] = h.node[i/2], h.node[i]
			i = i / 2
		} else {
			break
		}
	}
}

func (h *Heap) Poll() *HeapNode {
	if h.cnt == 0 {
		panic("no more node.")
	}

	hn := h.node[1]

	if h.cnt > 1 {
		h.node[h.cnt], h.node[1] = h.node[1], h.node[h.cnt]

		i := 1
		for {
			if i >= h.cnt {
				break
			}

			k := i * 2 // the mini/max idx
			m := k + 1

			if !h.opCmpFunc(h.node[k], h.node[i]) {
				k = i
			}

			if !h.opCmpFunc(h.node[k], h.node[m]) {
				k = m
			}

			if k != i {
				h.node[k], h.node[m] = h.node[m], h.node[k]
			}
		}
	}

	h.cnt--

	return hn
}
