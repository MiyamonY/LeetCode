// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 13:59:36 2019
//
package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode) (ret []string) {
	c := string(root.Val + 'a')
	if root.Left == nil && root.Right == nil {
		return []string{c}
	}
	if root.Left != nil {
		lefts := traverse(root.Left)
		for i := range lefts {
			ret = append(ret, lefts[i]+c)
		}
	}
	if root.Right != nil {
		rights := traverse(root.Right)
		for i := range rights {
			ret = append(ret, rights[i]+c)
		}
	}
	return
}

func smallestFromLeaf(root *TreeNode) string {
	strs := traverse(root)
	sort.Strings(strs)
	return strs[0]
}

func main() {

}
