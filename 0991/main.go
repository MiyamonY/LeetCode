// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jun 22 19:37:18 2019
//
package main

import "fmt"

func pow2(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret *= 2
	}
	return ret
}

func brokenCalc(X int, Y int) int {
	if X == Y {
		return 0
	} else if X > Y {
		return X - Y
	}

	n := 0
	for X < Y {
		X *= 2
		n++
	}

	num := n
	p := pow2(n)
	diff := X - Y
	for diff > 0 {
		//fmt.Println(diff)
		num += diff / p
		diff %= p
		p /= 2
	}
	return num
}

func main() {
	fmt.Println(brokenCalc(4, 10))
}
