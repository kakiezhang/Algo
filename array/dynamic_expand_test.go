package array

import "testing"

func TestDynamicArrayAdd(t *testing.T) {
	t.Log("add element at first.")
	da1 := NewDynamicArray(1)
	for _, v := range []int{10, 15, 20, 18, 29, 30, 34, 27} {
		da1.Add(0, v)
		t.Logf("add: %d, array: %+v\n", v, da1)
	}

	t.Log("add element at end.")
	da2 := NewDynamicArray(1)
	for _, v := range []int{10, 15, 20, 18, 29, 30, 34, 27} {
		da2.Add(da2.length, v)
		t.Logf("add: %d, array: %+v\n", v, da2)
	}

	da2.Add(3, 100)
	t.Logf("add: %d, array: %+v\n", 100, da2)
	da2.Add(2, 99)
	t.Logf("add: %d, array: %+v\n", 99, da2)
	da2.Add(20, 89)
	t.Logf("add: %d, array: %+v\n", 89, da2)
}
