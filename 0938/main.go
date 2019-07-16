// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jul 17 01:12:02 2019
//
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := dfs(root.Left)
		if L <= root.Val && root.Val <= R {
			sum += root.Val
		}
		sum += dfs(root.Right)
		return sum
	}

	return dfs(root)
}

func main() {

}
