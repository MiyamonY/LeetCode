// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Fri Jul  5 19:05:17 2019
//
package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	for len(popped) > 0 {
		if len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		} else {
			if len(pushed) > 0 {
				stack = append(stack, pushed[0])
				pushed = pushed[1:]
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
