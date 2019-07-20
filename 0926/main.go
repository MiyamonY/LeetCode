// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jul 20 16:44:05 2019
//
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func reverse(A []int) {
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
}
func minFlipsMonoIncr(S string) int {
	num0, num1 := 0, 0
	var sums1 []int
	for i := range S {
		if S[i] == '1' {
			num1++
		}
		sums1 = append(sums1, num1)
	}
	var sums0 []int
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '0' {
			num0++
		}
		sums0 = append(sums0, num0)
	}
	reverse(sums0)

	ans := min(sums1[len(sums1)-1],
		sums0[0])
	for i := range S {
		if i+1 < len(S) {
			ans = min(ans, sums1[i]+sums0[i+1])
		}
	}
	fmt.Println(sums1, sums0)
	return ans
}

func main() {

}
