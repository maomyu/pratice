package tree

import "fmt"

// 前序遍历
func PreTraverseBTree(root *treeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	PreTraverseBTree(root.Left)
	PreTraverseBTree(root.Right)
}

// 中序遍历
func MidTraverseTree(root *treeNode) {
	if root == nil {
		return
	}
	PreTraverseBTree(root.Left)
	fmt.Print(root.Val, " ")
	PreTraverseBTree(root.Right)
}

// 后序遍历
func PostTraverseTree(root *treeNode) {
	if root == nil {
		return
	}
	PreTraverseBTree(root.Left)
	PreTraverseBTree(root.Right)
	fmt.Print(root.Val, " ")
}

// 判断B是否是A的子结构
func HasSubtree(pRoot1 *treeNode, pRoot2 *treeNode) bool {
	result := false
	if pRoot1 != nil && pRoot2 != nil {
		if pRoot1.Val == pRoot2.Val {
			result = DoesTreeHHaveTree2(pRoot1, pRoot2)
		}
		if !result {
			result = HasSubtree(pRoot1.Left, pRoot2)
		}
		if !result {
			result = HasSubtree(pRoot1.Right, pRoot2)
		}
	}
	return result
}
func DoesTreeHHaveTree2(pRoot1 *treeNode, pRoot2 *treeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return DoesTreeHHaveTree2(pRoot1.Left, pRoot2.Left) && DoesTreeHHaveTree2(pRoot1.Right, pRoot2.Right)
}

// 二叉树的镜像
func MirrorRecursively(root *treeNode) {

}
