/**
求一个数的平方根，精确到小数后 6 位
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x float64

	x = biSearchSquare(0.9, 6)
	fmt.Println(x)

	x = biSearchSquare(4, 6)
	fmt.Println(x)

	x = biSearchSquare(15, 1)
	fmt.Println(x)
}

func biSearchSquare(x float64, d int) float64 {
	if x < 0 {
		panic("square num cannot lt 0")
	} else if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	ds := "%." + fmt.Sprint(d) + "f"
	xd, _ := strconv.ParseFloat(fmt.Sprintf(ds, x), 64)
	if x > 0 && x < 1 {
		return biRecursiveSquare(xd, x, 1, ds)
	} else {
		// x > 1
		return biRecursiveSquare(xd, 1, x, ds)
	}
}

func biRecursiveSquare(x, i, j float64, ds string) float64 {
	mid := (i + j) / 2
	y, _ := strconv.ParseFloat(fmt.Sprintf(ds, mid*mid), 64)

	fmt.Printf("x: %f, i: %f, j: %f, mid: %f, mid*mid: %f\n",
		x, i, j, mid, y)

	if x < y {
		return biRecursiveSquare(x, i, mid, ds)
	} else if x > y {
		return biRecursiveSquare(x, mid, j, ds)
	} else {
		return mid
	}
}
