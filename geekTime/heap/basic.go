/**
二叉堆
小顶堆，大顶堆
cnt 从 1 开始计算
score 是作为建堆的依据，拼接在 value 内
*/
package heap

type compareOperater interface {
	opCmp(c, p *HeapNode) bool
}

type SmallTopHeap struct {
	*Heap
}

func (sth *SmallTopHeap) opCmp(c, p *HeapNode) bool {
	if c.Score() < p.Score() {
		c, p = p, c
		return true
	} else {
		return false
	}
}

type BigTopHeap struct {
	*Heap
}

func (bth *BigTopHeap) opCmp(c, p *HeapNode) bool {
	if c.Score() > p.Score() {
		c, p = p, c
		return true
	} else {
		return false
	}
}

type Heap struct {
	node []*HeapNode
	cnt  int
	max  int
}

func (h *Heap) opCmp(c, p *HeapNode) bool

type HeapNode struct {
	key   string
	value interface{}
}

func (hn *HeapNode) Score() int

func (h *Heap) Add(hn *HeapNode) {
	if h.cnt == h.max {
		panic("reach max, couldn't add more node.")
	}

	h.node[h.cnt] = hn
	h.cnt++

	i := h.cnt

	for {
		if i == 1 {
			break
		}

		if h.opCmp(h.node[i], h.node[i/2]) {
			i = i / 2
		} else {
			break
		}
	}
}

func (h *Heap) Poll() *HeapNode {
}
