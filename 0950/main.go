// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 21:13:09 2019
//
package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	nums := make([]int, len(deck))
	for i := range nums {
		nums[i] = i
	}

	sim := []int{}
	for len(nums) > 0 {
		sim = append(sim, nums[0])
		nums = nums[1:]
		if len(nums) >= 2 {
			nums = append(nums[1:], nums[0])
		}
	}

	ans := make([]int, len(deck))
	for i, v := range sim {
		ans[v] = deck[i]
	}
	return ans
}

func main() {

}
