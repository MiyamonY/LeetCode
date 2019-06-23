// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 14:10:31 2019
//
package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type position struct {
	val int
	y   int
}

type positions []position

func (poss positions) Len() int {
	return len(poss)
}

func (poss positions) Less(i, j int) bool {
	if poss[i].y == poss[j].y {
		return poss[i].val < poss[j].val
	}
	return poss[i].y < poss[j].y
}

func (poss positions) Swap(i, j int) {
	poss[i], poss[j] = poss[j], poss[i]
}

var m map[int]positions

func traverse(root *TreeNode, pos int, y int) {
	if root == nil {
		return
	}

	if _, ok := m[pos]; ok {
		m[pos] = append(m[pos], position{root.Val, y})
	} else {
		m[pos] = positions{{root.Val, y}}
	}

	traverse(root.Left, pos-1, y+1)
	traverse(root.Right, pos+1, y+1)
}

func verticalTraversal(root *TreeNode) (ret [][]int) {
	m = map[int]positions{}
	traverse(root, 0, 0)
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		sort.Sort(m[k])
		vals := []int{}
		for _, v := range m[k] {
			vals = append(vals, v.val)
		}
		ret = append(ret, vals)
	}
	return
}

func main() {

}
