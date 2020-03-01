/**
用散列表+双向链表实现LRU
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	test1()
}

func test1() {
	alphabets := [][2]string{
		[2]string{"apple", "a1"},
		[2]string{"banana", "b2"},
		[2]string{"cat", "c3"},
		[2]string{"dog", "d4"},
		[2]string{"ear", "e5"},
		[2]string{"freeze", "f6"},
		[2]string{"giggle", "g7"},
		[2]string{"hahah", "h8"},
	}

	lhm := NewLinkedHashMap(4, 0.75)

	for _, item := range alphabets {
		k, v := item[0], item[1]
		fmt.Printf("add k[ %s ] v[ %s ]\n\n", k, v)
		lhm.Put(k, v)

		fmt.Println("Origin Map: ")
		fmt.Println(lhm)
		fmt.Println()
		// fmt.Println("Expand Map: ")
		// fmt.Println(hm.expMap)

		// if k == "cat" {
		// 	fmt.Printf("find apple: %s\n", lhm.Get("apple"))
		// } else if k == "dog" {
		// 	fmt.Printf("find dog: %s\n", lhm.Get(k))
		// }

		fmt.Println("-------------------")
	}

	var k, v string

	// // k := "apple"
	// k := "banana"
	// // k := "freeze"
	// fmt.Printf("del k[ %s ]\n\n", k)
	// lhm.Del(k)

	// fmt.Println("Origin Map: ")
	// fmt.Println(lhm)
	// fmt.Println()
	// fmt.Println("-------------------")

	// fmt.Println(lhm.head)

	// k = "apple"
	// k = "banana"
	k = "freeze"
	fmt.Printf("visit k[ %s ]\n\n", lhm.Get(k))

	fmt.Println("Origin Map: ")
	fmt.Println(lhm)
	fmt.Println()
	fmt.Println("-------------------")

	k, v = "ear", "e0"
	fmt.Printf("put k[ %s ] v[ %s ]\n\n", k, v)
	lhm.Put(k, v)

	fmt.Println("Origin Map: ")
	fmt.Println(lhm)
	fmt.Println()
	fmt.Println("-------------------")
}

type Data struct {
	key string
	val string
}

func (d Data) String() string {
	return fmt.Sprintf("%s:%s", d.key, d.val)
}

func newNode(key, val string) *Node {
	return &Node{
		data: Data{
			key: key,
			val: val,
		},
	}
}

type Node struct {
	data  Data
	prev  *Node
	next  *Node
	hnext *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("data: %s, ptr: %p, prev: %p, next: %p, hnext: %p",
		n.data, n, n.prev, n.next, n.hnext)
}

func NewLinkedHashMap(capacity int, factor float64) *LinkedHashMap {
	return &LinkedHashMap{
		buckets:  make([]*Node, capacity),
		capacity: capacity,
		factor:   factor,
	}
}

type LinkedHashMap struct {
	buckets []*Node
	head    *Node
	tail    *Node

	cnt      int
	capacity int
	factor   float64
}

func (lhm *LinkedHashMap) String() string {
	if lhm == nil {
		return fmt.Sprint("<nil>")
	}

	var s string

	for i := 0; i < lhm.capacity; i++ {
		s += fmt.Sprintf("[%d]", i)
		ptr := lhm.buckets[i]

		for {
			if ptr == nil {
				if i == lhm.capacity-1 {
					s += fmt.Sprintf("[%p:%v]", ptr, ptr)
				} else {
					s += fmt.Sprintf("[%p:%v]\n", ptr, ptr)
				}
				break
			}
			s += fmt.Sprintf("[%p][%p:%v][%p][%p]->", ptr.prev, ptr, ptr.data, ptr.next, ptr.hnext)
			ptr = ptr.hnext
		}
	}

	return s
}

func (lhm *LinkedHashMap) hash(key string) int {
	return lhm.hashCode(key) % lhm.capacity
}

func (lhm *LinkedHashMap) hashCode(key string) int {
	var rs float64
	arr := []rune(key)
	keyCnt := len(arr)
	// fmt.Println(arr)
	for i := 0; i < keyCnt; i++ {
		rs += float64(arr[i]-[]rune("a")[0]) * math.Pow(26, float64(keyCnt-i-1))
	}
	return int(rs)
}

func (lhm *LinkedHashMap) getNode(key string) (int, *Node, *Node) {
	bucketNum := lhm.hash(key)

	var hprev, ptr *Node
	ptr = lhm.buckets[bucketNum]
	for {
		if ptr == nil {
			break
		}
		if ptr.data.key == key {
			break
		}

		hprev = ptr
		ptr = ptr.hnext
	}

	return bucketNum, hprev, ptr
}

func (lhm *LinkedHashMap) removeDoublyNode(ptr, hprev *Node, bucketNum int) {
	lhm.removeZipperNode(ptr, hprev, bucketNum)

	if ptr.prev != nil {
		ptr.prev.next = ptr.next
	} else {
		lhm.buckets[bucketNum] = ptr.hnext
		lhm.head = ptr.next
	}

	if ptr.next != nil {
		ptr.next.prev = ptr.prev
	}
}

func (lhm *LinkedHashMap) removeZipperNode(ptr, hprev *Node, bucketNum int) {
	if hprev == nil {
		lhm.buckets[bucketNum] = ptr.hnext
		return
	}

	if hprev != nil && hprev.hnext != nil {
		hprev.hnext = hprev.hnext.hnext
	}
}

func (lhm *LinkedHashMap) appendZipperNode(ptr, hprev *Node, bucketNum int) {
	if hprev != nil {
		hprev.hnext = ptr
	} else {
		lhm.buckets[bucketNum] = ptr
	}
}

func (lhm *LinkedHashMap) removeHead() {
	if lhm.head != nil {
		headBucket := lhm.hash(lhm.head.data.key)
		lhm.removeDoublyNode(lhm.head, nil, headBucket)
	}
}

func (lhm *LinkedHashMap) moveToTail(ptr, hprev *Node, bucketNum int) {
	ptr.next = nil
	ptr.hnext = nil
	ptr.prev = lhm.tail

	if lhm.tail != nil {
		lhm.tail.next = ptr
	}
	lhm.tail = ptr
}

func (lhm *LinkedHashMap) Put(key, val string) {
	bucketNum, hprev, ptr := lhm.getNode(key)
	if ptr != nil {
		// 找到一个相同的
		ptr.data.val = val
		lhm.removeDoublyNode(ptr, hprev, bucketNum)

		s := lhm.buckets[bucketNum]
		if s != nil {
			for {
				if s.hnext == nil {
					s.hnext = ptr
					break
				}

				s = s.hnext
			}
		} else {
			lhm.buckets[bucketNum] = ptr
		}

	} else {
		ptr = newNode(key, val)
		if lhm.cnt == 0 {
			lhm.head = ptr
		}

		if lhm.cnt+1 <= lhm.capacity {
			lhm.cnt += 1
		} else {
			lhm.removeHead()
		}

		lhm.appendZipperNode(ptr, hprev, bucketNum)
		// fmt.Printf("hprev: %v\n", hprev)
		// fmt.Printf("ptr: %v\n", ptr)
	}

	// fmt.Printf("bucketNum: %d\n", bucketNum)
	lhm.moveToTail(ptr, hprev, bucketNum)
	// fmt.Printf("ptr: %v\n", ptr)
}

func (lhm *LinkedHashMap) Del(key string) {
	bucketNum, hprev, ptr := lhm.getNode(key)
	if ptr == nil {
		panic(fmt.Sprintf("key[ %s ] not found", key))
	}

	lhm.cnt -= 1
	fmt.Printf("bucketNum: %d\n", bucketNum)
	lhm.removeDoublyNode(ptr, hprev, bucketNum)
}

func (lhm *LinkedHashMap) Get(key string) string {
	bucketNum, hprev, ptr := lhm.getNode(key)
	if ptr == nil {
		panic(fmt.Sprintf("key[ %s ] not found", key))
	}

	lhm.removeDoublyNode(ptr, hprev, bucketNum)

	s := lhm.buckets[bucketNum]
	if s != nil {
		for {
			if s.hnext == nil {
				s.hnext = ptr
				break
			}

			s = s.hnext
		}
	} else {
		lhm.buckets[bucketNum] = ptr
	}

	lhm.moveToTail(ptr, hprev, bucketNum)

	return ptr.data.val
}
