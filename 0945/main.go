// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Fri Jul  5 15:58:32 2019
//
package main

import "sort"

func minIncrementForUnique(A []int) (ans int) {
	sort.Ints(A)
	num := -1
	for _, a := range A {
		if num <= a {
			num = a
		} else {
			ans += num - a
		}
		num++

	}
	return
}

func main() {

}
