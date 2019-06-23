// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 10:33:54 2019
//
package main

import "fmt"

type unionFind struct {
	roots  []int
	depths []int
}

func create(n int) *unionFind {
	roots := make([]int, n)
	depths := make([]int, n)
	for i := range roots {
		roots[i] = i
	}
	return &unionFind{roots, depths}
}

func (uf *unionFind) union(x, y int) {
	rx := uf.root(x)
	ry := uf.root(y)
	if uf.depths[rx] < uf.depths[ry] {
		uf.roots[rx] = ry
		uf.depths[ry]++
	} else if uf.depths[rx] < uf.depths[ry] {
		uf.roots[rx] = ry
	} else {
		uf.roots[ry] = rx
	}
}

func (uf *unionFind) find(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf *unionFind) root(x int) int {
	root := uf.roots[x]
	for root != uf.roots[root] {
		root = uf.roots[root]
	}
	return root
}

func equationsPossible(equations []string) bool {
	uf := create(26)
	for _, eq := range equations {
		left := int(eq[0] - 'a')
		right := int(eq[3] - 'a')
		if eq[1] == '=' {
			uf.union(left, right)
		}
	}

	for _, eq := range equations {
		left := int(eq[0] - 'a')
		right := int(eq[3] - 'a')
		if eq[1] == '!' {
			if uf.find(left, right) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
}
