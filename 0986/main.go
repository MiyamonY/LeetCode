// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 17:37:53 2019
//
package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func intervalIntersection(A [][]int, B [][]int) (ret [][]int) {
	for i := range A {
		startA, endA := A[i][0], A[i][1]
		for j := range B {
			startB, endB := B[j][0], B[j][1]
			if startB <= endA && endA <= endB {
				ret = append(ret, []int{max(startA, startB), endA})
			} else if startA <= endB && endB <= endA {
				ret = append(ret, []int{max(startA, startB), endB})
			}
		}
	}
	return
}

func main() {

}
