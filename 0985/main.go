// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 10:28:30 2019
//
package main

func sumEvenAfterQueries(A []int, queries [][]int) (ret []int) {
	sum := 0
	for _, val := range A {
		if val%2 == 0 {
			sum += val
		}
	}
	for _, qs := range queries {
		val := qs[0]
		index := qs[1]
		if A[index]%2 == 0 {
			sum -= A[index]
		}
		A[index] += val
		if A[index]%2 == 0 {
			sum += A[index]
		}
		ret = append(ret, sum)
	}
	return
}

func main() {

}
