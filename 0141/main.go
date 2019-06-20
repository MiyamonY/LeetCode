package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := map[*ListNode]interface{}{}
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = nil
		head = head.Next
	}
	return false
}
