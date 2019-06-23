// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 17:54:34 2019
//
package main

import (
	"fmt"
	"strings"
)

func strWithout3a3b(A int, B int) (ret string) {
	c0 := "a"
	c1 := "b"
	if A < B {
		A, B = B, A
		c0, c1 = c1, c0
	}

	diff := A - B
	extras := []string{}
	if diff > 0 {
		if B <= diff {
			if diff-B == 2 {
				extras = append(extras, c0+c0)
				diff -= 2
			} else if diff-B == 1 {
				extras = append(extras, c0)
				diff--
			}
		}
		for 0 < diff {
			extras = append([]string{c0}, extras...)
			diff--
		}

	}

	ret = strings.Join(extras, c0+c1)
	for len(ret) < A+B {
		ret = c0 + c1 + ret
	}
	return
}

func main() {
	fmt.Println(strWithout3a3b(2, 3))
}
