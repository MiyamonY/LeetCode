// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jul 20 16:35:52 2019
//
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFallingPathSum(A [][]int) int {
	sums := make([][]int, len(A))
	for i := range sums {
		sums[i] = make([]int, len(A[0]))
	}

	for i := len(A) - 1; i >= 0; i-- {
		if i == len(A)-1 {
			for j := range sums[i] {
				sums[i][j] = A[i][j]
			}
		} else {
			for j := range sums[i] {
				if j == 0 {
					sums[i][j] = A[i][j] + min(sums[i+1][j], sums[i+1][j+1])
				} else if j == len(sums[i])-1 {
					sums[i][j] = A[i][j] + min(sums[i+1][j-1], sums[i+1][j])
				} else {
					sums[i][j] = A[i][j] + min(sums[i+1][j-1], min(sums[i+1][j], sums[i+1][j+1]))
				}
			}
		}
	}
	ans := 1 << 60
	for j := range sums[0] {
		ans = min(ans, sums[0][j])
	}
	return ans
}

func main() {

}
