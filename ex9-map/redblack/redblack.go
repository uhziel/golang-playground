package bst

type AVLTree struct {
	root *node
}

type ColorType string

const (
	black ColorType = "black"
	red   ColorType = "red"
)

type node struct {
	Val         int
	Left, Right *node
	Color       ColorType
}

func New() *AVLTree {
	return &AVLTree{}
}

func rightRotate(root *node) *node {
	child := root.Left

	root.Left = child.Right
	child.Right = root
	child.Color = root.Color
	root.Color = red

	return child
}

func leftRotate(root *node) *node {
	child := root.Right

	root.Right = child.Left
	child.Left = root
	child.Color = root.Color
	root.Color = red

	return child
}

func flipColors(root *node) {
	root.Color = red
	root.Left.Color = black
	root.Right.Color = black
}

func isRed(root *node) bool {
	if root == nil {
		return false
	}

	return root.Color == red
}

func (t *AVLTree) String() string {
	return "TODO"
}

func (t *AVLTree) Search(target int) bool {
	curNode := t.root
	for curNode != nil {
		if target < curNode.Val {
			curNode = curNode.Left
		} else if target > curNode.Val {
			curNode = curNode.Right
		} else {
			return true
		}
	}
	return false
}

func (t *AVLTree) Add(num int) bool {
	added := false
	var dfs func(root *node, num int) *node
	dfs = func(root *node, num int) *node {
		if root == nil {
			added = true
			return &node{
				Val:   num,
				Color: red,
			}
		}

		if num < root.Val {
			root.Left = dfs(root.Left, num)
		} else if num > root.Val {
			root.Right = dfs(root.Right, num)
		}

		if added {
			if isRed(root.Right) && !isRed(root.Left) {
				root = leftRotate(root)
			}
			if isRed(root.Left) && isRed(root.Left.Left) {
				root = rightRotate(root)
			}
			if isRed(root.Left) && isRed(root.Right) {
				flipColors(root)
			}
		}

		return root
	}

	t.root = dfs(t.root, num)
	if t.root.Color == red {
		t.root.Color = black
	}
	return added
}

func (t *AVLTree) Remove(num int) bool {
	removed := false
	var dfs func(root *node, num int) *node
	dfs = func(root *node, num int) *node {
		if root == nil {
			return nil
		}

		if num < root.Val {
			root.Left = dfs(root.Left, num)
		} else if num > root.Val {
			root.Right = dfs(root.Right, num)
		} else {
			removed = true
			if root.Left == nil {
				root = root.Right
			} else if root.Right == nil {
				root = root.Left
			} else {
				mostLeft := root.Right
				for mostLeft.Left != nil {
					mostLeft = mostLeft.Left
				}

				root.Val = mostLeft.Val
				root.Right = dfs(root.Right, mostLeft.Val) // 这里用递归减少复杂重复的 updateHeight、rotate
			}
		}

		if removed {
			updateHeight(root)
			root = rotate(root)
		}

		return root
	}

	t.root = dfs(t.root, num)
	return removed
}
