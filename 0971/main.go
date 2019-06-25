// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 14:58:41 2019
//
package main

import "errors"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func aux(root *TreeNode, voyage []int, index int) (int, []int, error) {
	if root == nil {
		return index, []int{}, nil
	}
	if root.Val != voyage[index] {
		return index + 1, nil, errors.New("")
	}

	i, ret1, errl := aux(root.Left, voyage, index+1)
	j, ret2, errr := aux(root.Right, voyage, i)
	if errl == nil && errr == nil {
		return j, append(ret1, ret2...), nil
	}

	i, ret1, errl = aux(root.Right, voyage, index+1)
	j, ret2, errr = aux(root.Left, voyage, i)
	if errl == nil && errr == nil {
		return j, append(append(ret1, root.Val), ret2...), nil
	}
	return 0, nil, errors.New("")
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	if _, ret, err := aux(root, voyage, 0); err == nil {
		return ret
	}
	return []int{-1}
}

func main() {

}
