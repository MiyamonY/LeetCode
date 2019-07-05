// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 20:53:35 2019
//
package main

func sorted(w1, w2 string, m map[byte]int) bool {
	for i := range w1 {
		if i == len(w2) {
			return false
		}
		if m[w1[i]] < m[w2[i]] {
			return true
		} else if m[w1[i]] > m[w2[i]] {
			return false
		}
	}
	return true
}

func isAlienSorted(words []string, order string) bool {
	m := map[byte]int{}

	for i := range order {
		m[order[i]] = i
	}

	for i := 0; i+1 < len(words); i++ {
		if !sorted(words[i], words[i+1], m) {
			return false
		}
	}
	return true
}

func main() {

}
