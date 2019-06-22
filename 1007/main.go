// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jun 22 19:28:39 2019
//
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDominoRotations(A []int, B []int) int {
	ret := 1 << 30
	for n := 1; n <= 6; n++ {
		rotNumA := 0
		validA := true
		rotNumB := 0
		validB := true
		for i := range A {
			if A[i] == n {
			} else if A[i] != n && B[i] == n {
				rotNumA++
			} else {
				validA = false
			}

			if B[i] == n {
			} else if B[i] != n && A[i] == n {
				rotNumB++
			} else {
				validB = false
			}
		}
		if validA {
			ret = min(ret, rotNumA)
		}
		if validB {
			ret = min(ret, rotNumB)
		}
	}

	if ret == 1<<30 {
		return -1
	}
	return ret
}

func main() {

}
