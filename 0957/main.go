// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jun 27 00:52:19 2019
//
package main

import (
	"fmt"
	"reflect"
)

func cycle(cells []int) []int {
	ret := make([]int, 8)
	for j := 1; j <= 6; j++ {
		if cells[j-1] == cells[j+1] {
			ret[j] = 1
		} else {
			ret[j] = 0
		}
	}
	return ret
}

func prisonAfterNDays(cells []int, N int) []int {
	cells1 := cycle(cells)
	c := 0
	next := cells1
	for i := 0; i < 2<<6; i++ {
		next = cycle(next)
		if reflect.DeepEqual(cells1, next) {
			c = i + 1
			break
		}
	}
	m := (N - 1) % c
	ret := cells1
	for i := 0; i < m; i++ {
		ret = cycle(ret)
	}
	return ret
}

func main() {
	fmt.Printf("%+v", prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}
