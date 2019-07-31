// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jul 16 22:24:13 2019
//
package main

func diStringMatch(S string) []int {
	max := 0
	min := 0
	arr := []int{0}
	for i := range S {
		if S[i] == 'D' {
			min--
			arr = append(arr, min)
		} else {
			max++
			arr = append(arr, max)
		}
	}
	for i := range arr {
		arr[i] += -min
	}
	return arr
}

func main() {

}
