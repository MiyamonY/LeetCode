package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func aux(l1, l2 *ListNode, carry bool) *ListNode {
	if l1 == nil && l2 == nil {
		if carry {
			return &ListNode{1, nil}
		}
		return nil
	}

	if l1 == nil {
		if carry {
			return aux(&ListNode{1, nil}, l2, false)
		}
		return l2
	}

	if l2 == nil {
		if carry {
			return aux(&ListNode{1, nil}, l1, false)
		}
		return l1
	}

	val := l1.Val + l2.Val
	if carry {
		val++
	}

	if val/10 == 1 {
		carry = true
	} else {
		carry = false
	}
	return &ListNode{val % 10, aux(l1.Next, l2.Next, carry)}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return aux(l1, l2, false)
}

func main() {
	n := addTwoNumbers(&ListNode{1, &ListNode{0, nil}}, &ListNode{2, &ListNode{3, nil}})
	for n != nil {
		fmt.Printf("%+v\n", n.Val)
		n = n.Next
	}
}
