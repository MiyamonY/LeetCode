package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func aux(head *ListNode, visited []*ListNode) *ListNode {
	for _, v := range visited {
		if head == v {
			return head
		}
	}

	next := head.Next
	if next == nil {
		return nil
	}

	visited = append(visited, head)
	return aux(next, visited)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	visited := make([]*ListNode, 0)
	return aux(head, visited)
}

func main() {
	node := &ListNode{2, &ListNode{0, &ListNode{-4, nil}}}
	node.Next.Next.Next = node
	fmt.Println(detectCycle(&ListNode{3, node}) == node)
}
