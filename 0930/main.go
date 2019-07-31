// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jul 20 07:11:54 2019
//
package main

import (
	"fmt"
	"sort"
)

func numSubarraysWithSum(A []int, S int) (ans int) {
	for i := range A {
		if i+1 < len(A) {
			A[i+1] += A[i]
		}
	}
	A = append([]int{0}, A...)

	for i, a := range A {
		if S <= a {
			d := a - S
			lower := sort.Search(i+1, func(i int) bool {
				return d <= A[i]
			})
			upper := sort.Search(i+1, func(i int) bool {
				return d+1 <= A[i]

			})
			if upper == i+1 {
				upper--
			}
			ans += upper - lower
			fmt.Println(a, d, lower, upper)
		}
	}

	return ans
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, 0))
}
