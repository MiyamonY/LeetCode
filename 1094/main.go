// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 00:08:38 2019
//
package main

func carPooling(trips [][]int, capacity int) bool {
	dp := make([]int, 1002)
	for _, trip := range trips {
		num := trip[0]
		dp[trip[1]] += num
		dp[trip[2]] -= num
	}
	for i := range dp {
		if i+1 < len(dp) {
			dp[i+1] += dp[i]
		}
	}
	for i := range dp {
		if capacity < dp[i] {
			return false
		}
	}
	return true
}

func main() {

}
