// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 10:09:58 2019
//
package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxTurbulenceSize(A []int) (ret int) {
	numE := 1
	numO := 1

	for i := range A {
		if i+1 == len(A) {
			continue
		}
		if i%2 == 0 {
			if A[i] < A[i+1] {
				numE++
				ret = max(ret, numE)
				numO = 1
			} else if A[i] > A[i+1] {
				numO++
				ret = max(ret, numO)
				numE = 1
			} else {
				numO = 1
				numE = 1
			}
		} else {
			if A[i] < A[i+1] {
				numE = 1
				numO++
				ret = max(ret, numO)
			} else if A[i] > A[i+1] {
				numO = 1
				numE++
				ret = max(ret, numE)
			} else {
				numO = 1
				numE = 1
			}
		}
	}
	ret = max(ret, numO)
	ret = max(ret, numE)

	return ret
}

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	fmt.Println(maxTurbulenceSize([]int{4, 8, 12, 16}))
	fmt.Println(maxTurbulenceSize([]int{100}))
}
