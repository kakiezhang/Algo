/**
位图
1 字节的上限是 256，即 2^8
也就是同时可以表示 8 个数的二进制位
*/
package bitmap

type BitMap struct {
	bitnum int    // 可存储的二进制位数
	bytes  []byte // 需要用到的字节
}

func NewBitMap(num int) *BitMap {
	return &BitMap{
		bitnum: num,
		bytes:  make([]byte, num/8+1), // 这里在 8 的倍数时会多算一位，whatever
	}
}

func (bm *BitMap) Set(x int) {
	if x >= bm.bitnum {
		panic("exceed")
	}

	bitIndex := uint(x % 8)
	byteIndex := x / 8
	bm.bytes[byteIndex] |= (1 << bitIndex)
}

func (bm *BitMap) Get(x int) bool {
	if x >= bm.bitnum {
		panic("exceed")
	}

	bitIndex := uint(x % 8)
	byteIndex := x / 8
	return (bm.bytes[byteIndex] & (1 << bitIndex)) != 0
}
