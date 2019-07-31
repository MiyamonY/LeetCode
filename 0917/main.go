// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Fri Jul 26 00:29:53 2019
//
package main

func isAlpha(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func reverseOnlyLetters(S string) string {
	stack := []byte{}
	for i := range S {
		if isAlpha(S[i]) {
			stack = append(stack, S[i])
		}
	}

	ans := ""
	for i := range S {
		if isAlpha(S[i]) {
			ans += string(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			ans += string(S[i])
		}
	}
	return ans
}

func main() {

}
