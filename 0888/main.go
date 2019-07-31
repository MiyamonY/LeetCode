// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jul 31 19:50:34 2019
//
package main

import "sort"

func total(ar []int) (ans int) {
	for _, a := range ar {
		ans += a
	}
	return
}

func fairCandySwap(A []int, B []int) (ans []int) {
	t0 := total(A)
	t1 := total(B)
	half := (t0 + t1) / 2

	sort.Ints(B)
	for _, a := range A {
		i := sort.Search(len(B), func(i int) bool {
			return half <= t0-a+B[i]
		})
		if i < len(B) && t0-a+B[i] == half {
			return []int{a, B[i]}
		}
	}
	return []int{}
}

func main() {

}
