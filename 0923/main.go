// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Mon Jul 22 22:28:27 2019
//
package main

import (
	"fmt"
	"sort"
)

const MOD = 1e9 + 7

func threeSumMulti(A []int, target int) int {
	sort.Ints(A)
	ans := 0
	for i := range A {
		for j := i + 1; j < len(A); j++ {
			d := target - (A[i] + A[j])
			arr := A[j+1:]
			if d >= 0 && j+1 < len(A) {
				k := sort.Search(len(arr), func(i int) bool {
					return d <= arr[i]
				})
				if k < len(arr) && arr[k] == d {
					l := sort.Search(len(arr), func(i int) bool {
						return d+1 <= arr[i]
					})
					if l-1 < len(arr) && l-1 >= 0 && arr[l-1] == d {
						ans += l - k
					} else {
						ans++
					}
					ans %= MOD
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}
