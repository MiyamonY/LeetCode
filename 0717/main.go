// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 20:18:03 2019
//
package main

func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 0 {
			i++
		} else {
			i += 2
		}
	}
	return i == len(bits)-1
}

func main() {

}
