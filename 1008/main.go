package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	rest := preorder[1:]
	i := sort.Search(len(preorder[1:]), func(i int) bool {
		return val < rest[i]
	})

	left := bstFromPreorder(rest[:i])
	var right *TreeNode
	if i < len(rest) {
		right = bstFromPreorder(rest[i:])
	}

	return &TreeNode{Val: val,
		Left: left, Right: right}
}

func main() {
	fmt.Printf("%+v", bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
}
