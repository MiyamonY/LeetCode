// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 12:19:14 2019
//
package main

import "fmt"

func subarraysDivByK(A []int, K int) (ans int) {
	sums := []int{0}
	for i := range A {
		sums = append(sums, ((sums[i]+A[i])%K+K)%K)
	}
	fmt.Printf("%+v", sums)
	m := map[int]int{}
	for _, s := range sums {
		m[s]++
	}

	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	return ans
}

func main() {

}
