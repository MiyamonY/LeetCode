// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 16:12:06 2019
//
package main

func powerfulIntegers(x int, y int, bound int) []int {
	xs := []int{1}
	if x != 1 {
		acc := x
		for {
			xs = append(xs, acc)
			acc *= x
			if acc > bound {
				break
			}

		}
	}

	ys := []int{1}
	if y != 1 {
		acc := y
		for {
			ys = append(ys, acc)
			acc *= y
			if acc > bound {
				break
			}
		}
	}

	m := map[int]interface{}{}
	for _, x := range xs {
		for _, y := range ys {
			if x+y <= bound {
				m[x+y] = nil
			}
		}
	}

	ret := []int{}
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	powerfulIntegers(2, 3, 10)
}
