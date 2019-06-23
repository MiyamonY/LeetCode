// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Mon Jun 24 02:05:12 2019
//
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var num int

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := dfs(root.Left), dfs(root.Right)
	num += abs(l) + abs(r)
	return root.Val + l + r - 1
}

func distributeCoins(root *TreeNode) int {
	num = 0
	dfs(root)
	return num
}

func main() {
	fmt.Println(distributeCoins(&TreeNode{0, &TreeNode{3, nil, nil}, &TreeNode{0, nil, nil}}))
}
