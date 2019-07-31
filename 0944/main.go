// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jul 17 00:17:23 2019
//
package main

func minDeletionSize(A []string) (ans int) {
	for j := range A[0] {
		invalid := false
		for i := range A {
			if i+1 < len(A) {
				if A[i][j] > A[i+1][j] {
					invalid = true
					break
				}
			}
		}
		if invalid {
			ans++
		}
	}
	return
}

func main() {

}
