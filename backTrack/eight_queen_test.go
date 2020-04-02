package backTrack

import (
	"testing"
)

func TestEightQueen(t *testing.T) {
	eq := newEightQueen(8)
	eq.PrintArr()
	eq.Find(0)
	eq.PrintArr()
}
