// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jun 22 18:45:27 2019
//
package main

import (
	"fmt"
)

func bitwiseComplement(N int) int {
	msb := 0
	for i := 0; i < 64; i++ {
		if N&(1<<uint8(i)) != 0 {
			msb = i
		}
	}

	tmp := 0
	for i := 0; i <= msb; i++ {
		tmp <<= 1
		tmp |= 1
	}
	return tmp ^ N
}

func main() {
	fmt.Println(bitwiseComplement(10))
}
