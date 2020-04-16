/**
散列表
用链表法解决冲突
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// testHash()
	// testHashCode()
	// testPutGetDel()
	testExpand()
}

func testHash() {
	hm := NewHashMap(16, 0.75)

	var v int
	v = hm.hash("nice")
	fmt.Printf("bucket: nice: %d\n", v)

	v = hm.hash("bot")
	fmt.Printf("bucket: bot: %d\n", v)
}

func testHashCode() {
	hm := NewHashMap(16, 0.75)

	var v int
	v = hm.hashCode("nice")
	fmt.Printf("nice: %d\n", v)
}

func testPutGetDel() {
	hm := NewHashMap(16, 0.75)

	hm.put("apple", "a1")
	fmt.Println(hm)
	hm.put("banana", "b2")
	fmt.Println(hm)

	fmt.Println(hm.get("apple"))

	// hm.del("banana")
	// fmt.Println(hm)
	// hm.del("lemmon")
	// fmt.Println(hm)
}

func testExpand() {
	alphabets := [][2]string{
		[2]string{"apple", "a1"},
		[2]string{"banana", "b2"},
		[2]string{"cat", "c3"},
		[2]string{"dog", "d4"},
		[2]string{"ear", "e5"},
		[2]string{"freeze", "f6"},
		[2]string{"giggle", "g7"},
	}

	hm := NewHashMap(4, 0.75)

	for _, item := range alphabets {
		k, v := item[0], item[1]
		fmt.Printf("add k[ %s ] v[ %s ]\n\n", k, v)
		hm.put(k, v)

		fmt.Println("Origin Map: ")
		fmt.Println(hm)
		fmt.Println()
		fmt.Println("Expand Map: ")
		fmt.Println(hm.expMap)

		if k == "cat" {
			fmt.Printf("find apple: %s\n", hm.get("apple"))
		} else if k == "dog" {
			fmt.Printf("find dog: %s\n", hm.get(k))
		}

		fmt.Println("-------------------")
	}
}

func NewHashMap(capacity int, factor float64) *HashMap {
	return &HashMap{
		buckets:  make([]*Node, capacity),
		capacity: capacity,
		factor:   factor,
	}
}

type HashMap struct {
	buckets  []*Node
	cnt      int
	capacity int
	factor   float64
	exp      bool
	expMap   *HashMap
}

type Node struct {
	data Data
	next *Node
}

type Data struct {
	key string
	val string
}

func (d Data) String() string {
	return fmt.Sprintf("%s:%s", d.key, d.val)
}

func (hm *HashMap) String() string {
	if hm == nil {
		return fmt.Sprint("<nil>")
	}

	var s string

	for i := 0; i < hm.capacity; i++ {
		s += fmt.Sprintf("[%d]", i)
		ptr := hm.buckets[i]

		for {
			if ptr == nil {
				if i == hm.capacity-1 {
					s += fmt.Sprintf("[%p:%v]", ptr, ptr)
				} else {
					s += fmt.Sprintf("[%p:%v]\n", ptr, ptr)
				}
				break
			}
			s += fmt.Sprintf("[%p:%v][%p]->", ptr, ptr.data, ptr.next)
			ptr = ptr.next
		}
	}

	return s
}

func (hm *HashMap) hash(key string) int {
	return hm.hashCode(key) % hm.capacity
}

func (hm *HashMap) hashCode(key string) int {
	var rs float64
	arr := []rune(key)
	keyCnt := len(arr)
	// fmt.Println(arr)
	for i := 0; i < keyCnt; i++ {
		rs += float64(arr[i]-[]rune("a")[0]) * math.Pow(26, float64(keyCnt-i-1))
	}
	return int(rs)
}

func (hm *HashMap) get(key string) string {
	_, _, ptr := hm.getNode(key)
	if ptr != nil {
		return ptr.data.val
	}
	panic(fmt.Sprintf("no key element: %s", key))
}

func (hm *HashMap) getNode(key string) (int, *Node, *Node) {
	if hm.exp {
		bucketNum, prev, ptr := hm.expMap.getNode(key)
		if ptr != nil {
			return bucketNum, prev, ptr
		}
	}

	bucketNum := hm.hash(key)

	var prev, ptr *Node
	ptr = hm.buckets[bucketNum]

	for {
		if ptr == nil {
			break
		}
		if ptr.data.key == key {
			break
		}
		prev = ptr
		ptr = ptr.next
	}

	return bucketNum, prev, ptr
}

func (hm *HashMap) put(key, value string) {
	if (float64(hm.cnt+1)/float64(hm.capacity)) >= hm.factor || hm.exp {
		// 大于装载因子或者已经有扩容散列
		if !hm.exp {
			fmt.Println("Expanding map.")
			// 没有扩容过，扩容
			hm.exp = true
			hm.expMap = NewHashMap(hm.capacity*2, hm.factor)
		}

		// 放入新扩容的散列表
		hm.expMap.put(key, value)

		if hm.cnt > 0 {
			var i int
			for i = 0; i < hm.capacity; i++ {
				if hm.buckets[i] != nil {
					break
				}
			}

			node := hm.buckets[i]
			hm.del(node.data.key)
			hm.expMap.put(node.data.key, node.data.val)
		}

		if hm.cnt == 0 {
			fmt.Println("Moving expand map to origin.")
			*hm = *(hm.expMap)
			hm.expMap = nil
		}

		return
	}

	bucketNum, prev, ptr := hm.getNode(key)
	if ptr != nil {
		// 找到直接更新值
		ptr.data.val = value
	} else {
		hm.cnt += 1
		newNode := &Node{
			data: Data{
				key: key,
				val: value,
			},
			next: nil,
		}
		if prev != nil {
			prev.next = newNode
		} else {
			hm.buckets[bucketNum] = newNode
		}
	}
}

func (hm *HashMap) del(key string) {
	bucketNum, prev, ptr := hm.getNode(key)
	if ptr != nil {
		hm.cnt -= 1
		if prev != nil {
			prev.next = prev.next.next
		} else {
			hm.buckets[bucketNum] = ptr.next
		}
	} else {
		panic(fmt.Sprintf("no key element: %s", key))
	}
}
