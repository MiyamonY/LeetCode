// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Fri Jul 26 01:43:06 2019
//
package main

func min(as ...int) int {
	ret := as[0]
	for i := range as {
		if as[i] < ret {
			ret = as[i]
		}
	}
	return ret
}

func max(as ...int) int {
	ret := as[0]
	for i := range as {
		if ret < as[i] {
			ret = as[i]
		}
	}
	return ret
}

func smallestRangeI(A []int, K int) int {
	minVal := min(A...)
	maxVal := max(A...)
	if minVal+K < maxVal-K {
		return maxVal - K - (minVal + K)
	}
	return 0
}

func main() {

}
