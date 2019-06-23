// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 11:24:23 2019
//
package main

import "fmt"

func rev(A []int) []int {
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
	return A
}

func aux(A []int, K int, acc []int) []int {
	if K == 0 && len(A) == 0 {
		return acc
	}

	k := K % 10
	if len(A) == 0 {
		acc := append(acc, k)
		return aux(A, K/10, acc)
	}
	n := (A[0] + k) % 10
	m := (A[0] + k) / 10
	acc = append(acc, n)
	return aux(A[1:], K/10+m, acc)
}

func addToArrayForm(A []int, K int) (ret []int) {
	a := aux(rev(A), K, ret)
	return rev(a)

}

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 3}, 0))
}
