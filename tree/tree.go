package main

// 定义一个数的结点类
type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}
type Tree interface {
	// 动态创建一个二叉树（默认为二叉查找树）
	CreateTree(val interface{})
}

func main() {
	root := NewBinaryTree()
	root.CreateTree(6)
	root.CreateTree(4)
	root.CreateTree(3)
	root.CreateTree(5)
	root.CreateTree(3)
	root.CreateTree(4)
	root.CreateTree(10)
	root.CreateTree(7)
	root.CreateTree(8)
	PreTraverseBTree(root.Root())
}
