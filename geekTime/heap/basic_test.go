package heap

import "testing"

func TestSmallTopHeap(t *testing.T) {
	sth := NewSmallTopHeap(5)
	sth.ScoreFunc = func(hn *HeapNode) int {
		return hn.Value.(int)
	}

	sth.Add(NewHeapNode("B", 2))
	sth.Add(NewHeapNode("A", 1))
	sth.Add(NewHeapNode("C", 3))

	t.Log(sth)
	t.Logf("%+v", sth.node)
}
