package heap

import (
	"testing"
)

func TestSmallTopHeap(t *testing.T) {
	sth := NewSmallTopHeap(5)
	sth.ScoreFunc = func(hn *HeapNode) int {
		return hn.Value.(int)
	}

	sth.Add(NewHeapNode("B", 20))
	sth.Add(NewHeapNode("A", 1))
	sth.Add(NewHeapNode("C", 3))
	sth.Add(NewHeapNode("D", 5))

	t.Log("init: \n", sth)
	t.Log("===============\n\n")

	for sth.cnt > 0 {
		hn := sth.Poll()
		t.Logf("poll: node[%v]", hn)
		t.Log("tree: \n", sth)
		t.Log("nodes: ", sth.node)
		t.Log("===============\n\n")
	}
}

func TestBigTopHeap(t *testing.T) {
	bth := NewBigTopHeap(5)
	bth.ScoreFunc = func(hn *HeapNode) int {
		return hn.Value.(int)
	}

	bth.Add(NewHeapNode("B", 2))
	bth.Add(NewHeapNode("A", 1))
	bth.Add(NewHeapNode("C", 3))
	bth.Add(NewHeapNode("D", 5))

	t.Log("init: \n", bth)
	t.Log("===============\n\n")

	for bth.cnt > 0 {
		hn := bth.Poll()
		t.Logf("poll: node[%v]", hn)
		t.Log("tree: \n", bth)
		t.Log("nodes: ", bth.node)
		t.Log("===============\n\n")
	}
}
