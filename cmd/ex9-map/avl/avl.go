package bst

import (
	"strconv"
	"strings"
)

type AVLTree struct {
	root *node
}

type node struct {
	Val         int
	Height      int
	Left, Right *node
}

func New() *AVLTree {
	return &AVLTree{}
}

func getHeight(root *node) int {
	if root == nil {
		return -1
	}

	return root.Height
}

func updateHeight(root *node) {
	if root == nil {
		return
	}

	lh, rh := getHeight(root.Left), getHeight(root.Right)
	root.Height = max(lh, rh) + 1
}

func getBalanceFactor(root *node) int {
	if root == nil {
		return 0
	}
	return getHeight(root.Left) - getHeight(root.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rightRotate(root *node) *node {
	child := root.Left

	root.Left = child.Right
	child.Right = root
	updateHeight(root)
	updateHeight(child)

	return child
}

func leftRotate(root *node) *node {
	child := root.Right

	root.Right = child.Left
	child.Left = root
	updateHeight(root)
	updateHeight(child)

	return child
}

func rotate(root *node) *node {
	if root == nil {
		return nil
	}

	bf := getBalanceFactor(root)
	newRoot := root

	if bf > 1 {
		bfChild := getBalanceFactor(root.Left)
		if bfChild < 0 {
			root.Left = leftRotate(root.Left)
		}
		newRoot = rightRotate(root)
	} else if bf < -1 {
		bfChild := getBalanceFactor(root.Right)
		if bfChild > 0 {
			root.Right = rightRotate(root.Right)
		}
		newRoot = leftRotate(root)
	}

	return newRoot
}

func (t *AVLTree) String() string {
	height := getHeight(t.root)
	if height == -1 {
		return "empty Tree()"
	}
	m := height + 1
	n := 2<<(height+1) - 1
	storage := make([]byte, m*n)
	for i := range storage {
		storage[i] = ' '
	}
	ans := make([][]byte, m)
	for i := range ans {
		ans[i] = storage[0:n]
		storage = storage[n:]
	}
	var dfs func(root *node, r, c int)
	dfs = func(root *node, r, c int) {
		ans[r][c] = strconv.Itoa(root.Val)[0]
		if root.Left != nil {
			dfs(root.Left, r+1, c-2<<(height-r-1))
		}
		if root.Right != nil {
			dfs(root.Right, r+1, c+2<<(height-r-1))
		}
	}
	dfs(t.root, 0, (n-1)/2)

	var builder strings.Builder
	builder.WriteString("----------------------------------------\n")
	for i := range ans {
		builder.Write(ans[i])
		builder.WriteByte('\n')
	}

	return builder.String()
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
				Val: num,
			}
		}

		if num < root.Val {
			root.Left = dfs(root.Left, num)
		} else if num > root.Val {
			root.Right = dfs(root.Right, num)
		}

		if added {
			updateHeight(root)
			root = rotate(root)
		}

		return root
	}

	t.root = dfs(t.root, num)
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
