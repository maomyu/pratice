package tree

import "fmt"

type heap struct {
	val      []int
	capacity int
	count    int
}

func NewHeap(capacity int) *heap {
	return &heap{
		val:      make([]int, capacity),
		capacity: capacity,
		count:    0,
	}
}

// 创建一个堆
func (h *heap) CreateMinHeap(val int) {
	if h.count == h.capacity {
		return
	}

	h.val[h.count] = val
	i := h.count
	parent := (i - 1) / 2
	for parent >= 0 && (i-1) >= 0 {
		if h.val[parent] > h.val[i] {
			h.val[parent], h.val[i] = h.val[i], h.val[parent]
		}
		i = parent
		parent = (i - 1) / 2
	}
	h.count++
}

// 堆排序

// 堆删除根节点
func (h *heap) MoveRoot() {
	// 最后一个结点的元素
	h.val[h.count-1], h.val[0] = h.val[0], h.val[h.count-1]
	h.val[h.count-1] = 0
	h.count--
	h.AdjustHeap()
}

func (h *heap) AdjustHeap() {
	fmt.Println(h.val)
	top := 0
	// 左孩子结点
	j := 2*top + 1
	for j < h.count {
		if j+1 < h.count && h.val[j+1] < h.val[j] {
			j++
		}
		if h.val[top] > h.val[j] {
			h.val[top], h.val[j] = h.val[j], h.val[top]
			top = j
			j = 2*top + 1
		} else {
			break
		}
	}
	// h.val[top] = temp
}
