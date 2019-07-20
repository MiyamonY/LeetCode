// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jul 20 16:12:31 2019
//
package main

func isLongPressedName(name string, typed string) bool {
	ind := 0
	for i := range name {
		n := name[i]

		if ind >= len(typed) {
			return false
		}
		if n != typed[ind] {
			if ind == 0 || typed[ind-1] != typed[ind] {
				return false
			}

			cur := typed[ind]
			for ind < len(typed) && cur == typed[ind] {
				ind++
			}
			if ind == len(typed) || typed[ind] != n {
				return false
			}

		}
		ind++
	}
	return true
}

func main() {

}
