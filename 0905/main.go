// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jul  2 19:16:29 2019
//
package main

import "sort"

func sortArrayByParity(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		if A[i]%2 == 0 {
			return true
		} else if A[j]%2 == 0 {
			return false
		}
		return A[i] < A[j]
	})
	return A
}

func main() {

}
