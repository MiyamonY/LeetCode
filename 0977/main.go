// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 11:59:02 2019
//
package main

import "sort"

func sortedSquares(A []int) (ret []int) {
	for _, a := range A {
		ret = append(ret, a*a)
	}
	sort.Ints(ret)
	return ret
}

func main() {

}
