// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 12:10:15 2019
//
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var parents map[*TreeNode]*TreeNode

func dfs(parent *TreeNode, root *TreeNode) {
	if root == nil {
		return
	}

	parents[root] = parent
	dfs(root, root.Left)
	dfs(root, root.Right)
}

type node struct {
	dist int
	node *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) (ret []int) {
	parents = map[*TreeNode]*TreeNode{}
	dfs(nil, root)

	queue := []node{node{0, target}}
	visited := map[*TreeNode]interface{}{}
	visited[target] = nil
	for len(queue) > 0 {
		top := queue[0].node
		dist := queue[0].dist
		queue = queue[1:]
		if dist == K {
			ret = append(ret, top.Val)
		}
		visited[top] = true
		for _, n := range []*TreeNode{top.Left, top.Right, parents[top]} {
			if _, ok := visited[n]; ok || top == nil || n == nil {
				continue
			}
			queue = append(queue, node{dist + 1, n})
		}
	}
	return ret
}

func main() {

}
