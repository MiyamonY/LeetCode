// Package main provides
//
// File:  mian.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 12:04:13 2019
//
package main

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 0; i-- {
		if i > 1 {
			if A[i] < A[i-1]+A[i-2] {
				return A[i] + A[i-1] + A[i-2]
			}
		}
	}
	return 0
}

func main() {

}
