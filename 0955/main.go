// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jul  6 10:13:30 2019
//
package main

import "fmt"

func allSorted(A []string) bool {
	for i := range A {
		if i+1 < len(A) {
			if A[i] > A[i+1] {
				return false
			}
		}
	}

	return true
}

func minDeletionSize(A []string) (ans int) {
	cols := make([]string, len(A))
	for i := 0; i < len(A[0]); i++ {
		cols2 := make([]string, len(A))
		copy(cols2, cols)
		for j, a := range A {
			cols2[j] = cols2[j] + string(a[i])
		}

		if allSorted(cols2) {
			cols = cols2
		} else {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(minDeletionSize([]string{"xga", "xfb", "yfa"}))
}
