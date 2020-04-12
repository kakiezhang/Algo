package bitmap

import (
	"testing"
)

func TestBitMap(t *testing.T) {
	// var a = [7]byte{100, 13, 14, 15}
	// fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))

	bm := NewBitMap(8)
	t.Logf("init: %v", bm.bytes)

	for _, v := range []int{0, 1, 2, 3, 4, 5, 6, 7} {
		bm.Set(v)
		t.Logf("after set %d: bytes: %v", v, bm.bytes)
	}

	t.Logf("get 1: %t", bm.Get(1))
	t.Logf("get 0: %t", bm.Get(0))
}
