// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jun 26 21:44:32 2019
//
package main

func repeatedNTimes(A []int) int {
	m := map[int]int{}
	for i := range A {
		m[A[i]]++
	}
	for k, v := range m {
		if v > 1 {
			return k
		}
	}
	return 0
}

func main() {

}
