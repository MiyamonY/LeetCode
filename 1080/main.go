package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// max ...
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func sum(root *TreeNode, init int) *Node {
	if root == nil {
		return nil
	}

	left := sum(root.Left, root.Val+init)
	right := sum(root.Right, root.Val+init)

	var val int
	if left == nil && right == nil {
		val = root.Val + init
	} else if left == nil {
		val = right.Val
	} else if right == nil {
		val = left.Val
	} else {
		val = max(left.Val, right.Val)
	}

	return &Node{Val: val, Left: left, Right: right}
}

func truncate(root *TreeNode, root1 *Node, limit int) *TreeNode {
	if root1 == nil || root1.Val < limit {
		return nil
	}

	return &TreeNode{root.Val,
		truncate(root.Left, root1.Left, limit),
		truncate(root.Right, root1.Right, limit)}
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	sumNode := sum(root, 0)
	return truncate(root, sumNode, limit)
}

func main() {
}
