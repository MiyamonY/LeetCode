// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jul 17 00:20:49 2019
//
package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func bagOfTokensScore(tokens []int, P int) (ans int) {
	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i] < tokens[j]
	})
	left := 0
	right := len(tokens) - 1
	score := 0
	for left <= right {
		oldLeft, oldRight := left, right
		for left <= right && tokens[left] <= P {
			P -= tokens[left]
			score++
			left++
		}
		ans = max(ans, score)
		if score > 0 && left <= right {
			score--
			P += tokens[right]
			right--
		}
		if oldLeft == left && right == oldRight {
			break
		}
	}
	return
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
