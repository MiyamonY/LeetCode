// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 23:22:53 2019
//
package main

import "fmt"

func largestTimeFromDigits(A []int) string {
	max := 0
	ans := ""
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				h := A[i]*10 + A[j]
				m := A[k]*10 + A[6-i-j-k]
				if 0 <= h && h < 24 && 0 <= m && m < 60 {
					if max <= h*100+m {
						max = h*100 + m
						ans = fmt.Sprintf("%02d:%02d", h, m)
					}
				}
			}
		}
	}
	return ans
}

func main() {

}
