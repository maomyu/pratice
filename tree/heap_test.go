package tree

import (
	"fmt"
	"testing"
)

func TestCreateHeap(t *testing.T) {
	h := NewHeap(7)
	h.CreateMinHeap(1)
	h.CreateMinHeap(3)
	h.CreateMinHeap(2)
	h.CreateMinHeap(0)
	h.CreateMinHeap(1)
	h.CreateMinHeap(5)
	h.CreateMinHeap(9)
	fmt.Println(h.val)
	h.MoveRoot()
	fmt.Println(h.val)
}
