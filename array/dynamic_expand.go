/**
动态扩容的数组
*/
package array

type DynamicArray struct {
	elements []int
	length   int
	capacity int
}

func NewDynamicArray(cnt int) *DynamicArray {
	if cnt <= 0 {
		panic("array capacity must be gt 0.")
	}
	return &DynamicArray{
		elements: make([]int, cnt),
		length:   0,
		capacity: cnt,
	}
}

func (da *DynamicArray) Add(m, x int) {
	// m：加入的索引位置，x：要加入的值
	if m > da.length {
		panic("couldn't add, hollow between elements.")
	}

	if da.length+1 > da.capacity {
		da.capacity = da.capacity << 1
		tmp := make([]int, da.capacity) // 申请一个 2 倍大的数组
		for i := 0; i < da.length; i++ {
			tmp[i] = da.elements[i]
		}
		da.elements = tmp
	}

	if m < da.length {
		for i := da.length - 1; i >= m; i-- {
			da.elements[i+1] = da.elements[i]
		}
	}

	da.elements[m] = x
	da.length += 1
}
