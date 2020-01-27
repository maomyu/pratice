package tree

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
