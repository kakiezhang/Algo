/**
固定长度的有序数组，动态增删改
*/
package array

type FixedOrderArray struct {
	elements []int
	length   int
	capacity int
}

func NewFixedOrderArray(cnt int) *FixedOrderArray {
	return &FixedOrderArray{
		elements: make([]int, cnt),
		capacity: cnt,
	}
}

func (foa *FixedOrderArray) add(x int) {
}

func (foa *FixedOrderArray) del() {
}

func (foa *FixedOrderArray) update() {
}
