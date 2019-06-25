// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 21:47:00 2019
//
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}

	return check(root.Left, val) && check(root.Right, val)
}

func isUnivalTree(root *TreeNode) bool {
	return check(root, root.Val)
}

func main() {

}
