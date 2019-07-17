// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul 18 00:48:32 2019
//
package main

import "fmt"

const MOD = 1e9 + 7

func knightDialer(N int) int {
	dp := make([][10]int, N)
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 0; i < N-1; i++ {
		dp[i+1][0] = dp[i][4] + dp[i][6]
		dp[i+1][1] = dp[i][6] + dp[i][8]
		dp[i+1][2] = dp[i][7] + dp[i][9]
		dp[i+1][3] = dp[i][4] + dp[i][8]
		dp[i+1][4] = dp[i][0] + dp[i][3] + dp[i][9]
		dp[i+1][5] = 0
		dp[i+1][6] = dp[i][0] + dp[i][1] + dp[i][7]
		dp[i+1][7] = dp[i][2] + dp[i][6]
		dp[i+1][8] = dp[i][1] + dp[i][3]
		dp[i+1][9] = dp[i][2] + dp[i][4]
		for j := range dp[i+1] {
			dp[i+1][j] %= MOD
		}
	}

	ans := 0
	for i := range dp[N-1] {
		ans += dp[N-1][i]
		ans %= MOD
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(knightDialer(n))
}
