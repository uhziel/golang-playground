package bst

import (
	"strconv"
	"strings"
)

type BinarySearchTree struct {
	root *node
}

type node struct {
	Val         int
	Left, Right *node
}

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

func getHeight(root *node) int {
	if root == nil {
		return -1
	}

	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b *BinarySearchTree) String() string {
	height := getHeight(b.root)
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
	dfs(b.root, 0, (n-1)/2)

	var builder strings.Builder
	builder.WriteString("----------------------------------------\n")
	for i := range ans {
		builder.Write(ans[i])
		builder.WriteByte('\n')
	}

	return builder.String()
}

func (b *BinarySearchTree) Search(target int) bool {
	curNode := b.root
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

func (b *BinarySearchTree) Add(num int) bool {
	newNode := &node{
		Val: num,
	}

	if b.root == nil {
		b.root = newNode
		return true
	}

	curNode := b.root
	for curNode != nil {
		if num < curNode.Val {
			if curNode.Left == nil {
				curNode.Left = newNode
				return true
			}
			curNode = curNode.Left
		} else if num > curNode.Val {
			if curNode.Right == nil {
				curNode.Right = newNode
				return true
			}
			curNode = curNode.Right
		} else {
			return false
		}
	}

	return false
}

func (b *BinarySearchTree) find(target int) (curNode, lastNode *node, left bool) {
	curNode = b.root
	for curNode != nil {
		if target < curNode.Val {
			left = true
			lastNode = curNode
			curNode = curNode.Left
		} else if target > curNode.Val {
			left = false
			lastNode = curNode
			curNode = curNode.Right
		} else {
			return
		}
	}
	curNode, lastNode = nil, nil
	return
}

func (b *BinarySearchTree) Remove(num int) bool {
	curNode, lastNode, left := b.find(num)
	if curNode == nil {
		return false
	}

	var newRoot *node
	if curNode.Left == nil {
		newRoot = curNode.Right
	} else if curNode.Right == nil {
		newRoot = curNode.Left
	} else {
		var last *node
		mostRight := curNode.Left
		for mostRight.Right != nil {
			last = mostRight
			mostRight = mostRight.Right
		}

		newRoot = mostRight
		if last == nil {
			mostRight.Right = curNode.Right
		} else {
			last.Right = mostRight.Left
			mostRight.Left = curNode.Left
			mostRight.Right = curNode.Right
		}
	}

	if lastNode == nil {
		b.root = newRoot
	} else if left {
		lastNode.Left = newRoot
	} else {
		lastNode.Right = newRoot
	}

	return true
}
