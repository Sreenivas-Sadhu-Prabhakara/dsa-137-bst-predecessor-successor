package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// PredSucc returns the inorder predecessor (largest value < key) and
// successor (smallest value > key) in the BST; either is nil when absent.
// The key itself need not be present. Two independent descents, each
// keeping the best candidate seen while being forced past it.
func PredSucc(root *TreeNode, key int) (pred, succ *TreeNode) {
	for n := root; n != nil; {
		if n.Val < key {
			pred = n // candidate; something bigger may still be below-right
			n = n.Right
		} else {
			n = n.Left
		}
	}
	for n := root; n != nil; {
		if n.Val > key {
			succ = n
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return pred, succ
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(8, tn(3, tn(1, nil, nil), tn(6, tn(4, nil, nil), tn(7, nil, nil))), tn(10, nil, tn(14, tn(13, nil, nil), nil)))
	p, s := PredSucc(root, 8)
	fmt.Println(p.Val, s.Val)
}
