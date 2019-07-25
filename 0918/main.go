// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Fri Jul 26 01:13:11 2019
//
package main

import "fmt"

func max(as ...int) int {
	ret := as[0]
	for i := range as {
		if ret < as[i] {
			ret = as[i]
		}
	}
	return ret
}

func reverse(A []int) {
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
}

func maxSubarraySumCircular(A []int) int {
	oneInterval := []int{A[0]}
	for i := 1; i < len(A); i++ {
		oneInterval = append(oneInterval, max(A[i], A[i]+oneInterval[len(oneInterval)-1]))
	}
	ans := max(oneInterval...)
	fmt.Println(oneInterval)
	halfIntervalA := []int{A[0]}
	for i := 1; i < len(A); i++ {
		halfIntervalA = append(halfIntervalA, A[i]+halfIntervalA[len(halfIntervalA)-1])
	}

	for i := range halfIntervalA {
		if i > 0 {
			halfIntervalA[i] = max(halfIntervalA[i], halfIntervalA[i-1])
		}
	}

	halfIntervalB := []int{A[len(A)-1]}
	for i := len(A) - 2; i >= 0; i-- {
		halfIntervalB = append(halfIntervalB, A[i]+halfIntervalB[len(halfIntervalB)-1])
	}
	for i := range halfIntervalB {
		if i > 0 {
			halfIntervalB[i] = max(halfIntervalB[i], halfIntervalB[i-1])
		}
	}

	reverse(halfIntervalB)
	fmt.Println(halfIntervalB)
	for i := range A {
		if i+1 < len(A) {
			n := halfIntervalA[i] + halfIntervalB[i+1]
			if ans < n {
				ans = n
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
	fmt.Println(maxSubarraySumCircular([]int{3, -1, 2, -1}))
	fmt.Println(maxSubarraySumCircular([]int{-2, 2, -2, 9}))
}
