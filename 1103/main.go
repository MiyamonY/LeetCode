// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 01:52:46 2019
//
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func distributeCandies(candies int, num_people int) []int {
	arr := make([]int, num_people)
	num := 0
	for candies > 0 {
		n := min(num+1, candies)
		arr[num%num_people] += n
		num++
		candies -= n
	}
	return arr
}

func main() {

}
