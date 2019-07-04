// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jul  2 19:50:14 2019
//
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, cb func(int)) {
	if root == nil {
		return
	}
	inOrder(root.Left, cb)
	cb(root.Val)
	inOrder(root.Right, cb)
}

func increasingBST(root *TreeNode) *TreeNode {
	top := &TreeNode{0, nil, nil}
	cur := top

	cb := func(val int) {
		cur.Right = &TreeNode{val, nil, nil}
		cur = cur.Right
	}
	inOrder(root, cb)
	return top.Right
}

func main() {
}
