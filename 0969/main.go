// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 16:53:33 2019
//
package main

import (
	"fmt"
)

// rev ...
func rev(A []int) {
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
}

func linear_search(A []int, n int) int {
	for i, a := range A {
		if a == n {
			return i
		}
	}
	return -1
}

func pancakeSort(A []int) (ret []int) {
	for i := len(A); i >= 1; i-- {
		index := linear_search(A, i)
		ret = append(ret, index+1)
		rev(A[:index+1])
		ret = append(ret, i)
		rev(A[:i])
	}
	return
}

func main() {
	a := []int{3, 2, 4, 1}
	fmt.Println(pancakeSort(a))
	fmt.Println(a)
}
