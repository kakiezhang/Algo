/**
动态扩容的数组
*/
package array

type Array struct {
	elements []int
	length   int
	capacity int
}

func NewArray(cnt int) *Array {
	if cnt <= 0 {
		panic("array capacity must be gt 0.")
	}
	return &Array{
		elements: make([]int, cnt),
		length:   0,
		capacity: cnt,
	}
}

func (a *Array) Add(x int) {
	if a.length+1 > a.capacity {
		a.capacity = a.capacity << 1
		tmp := make([]int, a.capacity) // 申请一个 2 倍大的数组
		for i := 0; i < a.length; i++ {
			tmp[i] = a.elements[i]
		}
		a.elements = tmp
	}

	a.elements[a.length] = x
	a.length += 1
}
