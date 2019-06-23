// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 20:48:48 2019
//
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, len(days)+1)
	for i := range dp {
		dp[i] = 1 << 30
	}
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		d := days[i-1]
		min7 := 1 << 30
		min30 := 1 << 30
		for j := 0; j < i; j++ {
			if d-7+1 <= days[j] {
				min7 = min(min7, dp[j])
			}

			if d-30+1 <= days[j] {
				min30 = min(min30, dp[j])
			}
		}
		dp[i] = min(dp[i], min(costs[0]+dp[i-1], min(min7+costs[1], min30+costs[2])))
	}

	return dp[len(dp)-1]
}

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}))
}
