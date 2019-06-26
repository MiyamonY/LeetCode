// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jun 27 00:35:32 2019
//
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numerize(root *TreeNode, n int) {
	if root == nil {
		return
	}
	root.Val = n
	numerize(root.Left, 2*n)
	numerize(root.Right, 2*n+1)
}

func bfs(root *TreeNode, arr *[]int) {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		*arr = append(*arr, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
}

func isCompleteTree(root *TreeNode) bool {
	numerize(root, 1)
	arr := []int{}
	bfs(root, &arr)
	for i, v := range arr {
		if v != i+1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isCompleteTree(&TreeNode{1, &TreeNode{2, nil, nil}, nil}))
}
