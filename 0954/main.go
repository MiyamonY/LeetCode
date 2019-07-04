// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 20:40:37 2019
//
package main

import "sort"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func canReorderDoubled(A []int) bool {
	as := []int{}
	for _, a := range A {
		as = append(as, abs(a))
	}
	sort.Ints(as)

	m := map[int]int{}
	for _, v := range as {
		m[v]++
	}

	for _, a := range as {
		if m[a] == 0 {
			continue
		}
		if m[2*a] == 0 {
			return false
		}
		m[a]--
		m[2*a]--
	}
	return true
}

func main() {

}
