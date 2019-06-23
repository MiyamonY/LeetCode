// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 10:14:48 2019
//
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func depth(root *TreeNode, val int, dep int) int {
	if root == nil {
		return -1
	}
	if root.Val == val {
		return dep
	}
	return max(depth(root.Left, val, dep+1), depth(root.Right, val, dep+1))
}

func parent(root *TreeNode, val int, par *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return par
	}

	parentL := parent(root.Left, val, root)
	parentR := parent(root.Right, val, root)
	if parentL != nil {
		return parentL
	}
	return parentR
}

func isCousins(root *TreeNode, x int, y int) bool {
	depx := depth(root, x, 0)
	depy := depth(root, y, 0)
	if depx != depy {
		return false
	}
	parentx := parent(root, x, nil)
	parenty := parent(root, y, nil)
	if parentx == parenty {
		return false
	}
	return true
}

func main() {

}
