package tree

import (
	"fmt"
	"testing"
)

func TestIsContainValue(t *testing.T) {
	tree := NewBinaryTree()
	tree.CreateTree(2)
	tree.CreateTree(4)
	tree.CreateTree(6)
	tree.CreateTree(3)
	tree.CreateTree(1)
	fmt.Println(tree.IsContainValue(9))
}
