package main 

// 定义一个二叉树类
type binaryTree struct {
	// 根结点
	root *treeNode
	// 二叉树的高度
	h int
}

func NewBinaryTree() *binaryTree {
	return &binaryTree{
		root: nil,
		h:    0,
	}
}

// 获取根节点
func (t *binaryTree) Root() *treeNode {
	return t.root
}

func (t *binaryTree) SetRoot(val treeNode) {
	t.root = &val
}

// 获取深度
func (t *binaryTree) H() int {
	return t.h
}

func (t *binaryTree) SetH(h int) {
	t.h = h
}

func (t *binaryTree) CreateTree(val int) {
	if t.Root() == nil {
		t.SetRoot(treeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		})
	} else {
		root := t.Root()
		for root != nil {
			if val <= root.Val && root.Left == nil {
				fmt.Println("插入左节点")
				root.Left = &treeNode{val, nil, nil}
				break
			} else if val > root.Val && root.Right == nil {
				fmt.Println("插入右节点")
				root.Right = &treeNode{val, nil, nil}
				break
			} else if val <= root.Val && root.Left != nil {
				root = root.Left
			} else if val > root.Val && root.Right != nil {
				root = root.Right
			}
		}
	}
}