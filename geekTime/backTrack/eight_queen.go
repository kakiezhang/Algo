/**
八皇后
8x8(nxn)的棋盘每行放置一个棋子
棋子间不同行，列，对角线
*/
package backTrack

import "fmt"

type EightQueen struct {
	arr [][]int
	cnt int
}

func newEightQueen(cnt int) *EightQueen {
	eq := &EightQueen{
		arr: make([][]int, cnt),
		cnt: cnt,
	}

	for i := 0; i < cnt; i++ {
		eq.arr[i] = make([]int, cnt)
	}

	return eq
}

func (eq *EightQueen) PrintArr() {
	fmt.Println("[")
	for i := 0; i < eq.cnt; i++ {
		for j := 0; j < eq.cnt; j++ {
			fmt.Printf("%d ", eq.arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println("]")
}

func (eq *EightQueen) Find(i int) {
	if i == eq.cnt { // 行i
		return
	}

	// 逐行分配位置，每次看是否与前面的位置互斥即可
	var ex bool
	var j int
	for ; j < eq.cnt; j++ { // 列j
		if eq.arr[i][j] == 1 {
			// 在下一行不匹配的情况下，
			// 当前行的占位需要重新调整
			eq.arr[i][j] = 0
			continue
		}
		ex = eq.ExistMutex(j, i)
		if !ex {
			break
		}
	}
	if !ex {
		// 没有互斥，填位，跳到下一行
		eq.arr[i][j] = 1
		i += 1
	} else {
		// 存在互斥，回到上一行
		i -= 1
	}

	eq.Find(i)
}

func (eq *EightQueen) ExistMutex(x, y int) bool {
	// 存在互斥元素？列x，行y
	var rs bool
	j := 1

	for i := y - 1; i >= 0; i-- {
		if eq.arr[i][x] == 1 {
			// 列上有没有互斥的
			rs = true
			break
		}
		if x >= j && eq.arr[i][x-j] == 1 {
			// 对角线上有没有互斥的
			rs = true
			break
		}
		if x+j < eq.cnt && eq.arr[i][x+j] == 1 {
			// 斜对角线上有没有互斥的
			rs = true
			break
		}
		j += 1
	}

	return rs
}
