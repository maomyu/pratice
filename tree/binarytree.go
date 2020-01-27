package tree

// 定义一个二叉查找树类
type binaryTree struct {
	// 根结点
	root *treeNode
	// 结点的数量
	count int
}

// 二叉查找树
func NewBinaryTree() *binaryTree {
	return &binaryTree{
		root: nil,
	}
}

// 获取根节点
func (t *binaryTree) Root() *treeNode {
	return t.root
}

func (t *binaryTree) SetRoot(val treeNode) {
	t.root = &val
}

// 创建一个二叉查找树
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
				root.Left = &treeNode{val, nil, nil}
				t.count++
				break
			} else if val > root.Val && root.Right == nil {
				root.Right = &treeNode{val, nil, nil}
				t.count++
				break
			} else if val <= root.Val && root.Left != nil {
				root = root.Left
			} else if val > root.Val && root.Right != nil {
				root = root.Right
			}
		}
	}
}
