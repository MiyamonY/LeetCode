// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 20:59:05 2019
//
package main

import (
	"fmt"
	"strconv"
)

func aux(n int, top int, k int, a []int) []string {
	a = append(a, top)
	if n == len(a) {
		s := ""
		for i := range a {
			c := strconv.Itoa(a[i])
			s += c
		}
		return []string{s}
	}

	var ret0, ret1 []string
	if top+k <= 9 {
		ret0 = aux(n, top+k, k, a)
	}
	if top+k != top-k && top-k >= 0 {
		ret1 = aux(n, top-k, k, a)
	}
	return append(ret0, ret1...)
}

func numsSameConsecDiff(N int, K int) (ret []int) {
	tmp := []string{}
	if N == 1 {
		ret = append(ret, 0)
	}
	for i := 1; i <= 9; i++ {
		tmp = append(tmp, aux(N, i, K, []int{})...)
	}
	for i := range tmp {
		num, _ := strconv.Atoi(tmp[i])
		ret = append(ret, num)
	}
	return
}

func main() {
	fmt.Println(numsSameConsecDiff(2, 1))
}
