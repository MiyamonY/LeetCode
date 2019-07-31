// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jul 16 22:32:30 2019
//
package main

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 0
	for i+1 < len(A) && A[i] < A[i+1] {
		i++
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for i+1 < len(A) && A[i] > A[i+1] {
		i++
	}
	return i == len(A)-1
}

func main() {

}
