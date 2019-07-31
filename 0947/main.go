// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jul  6 15:56:41 2019
//
package main

import "fmt"

type UF struct {
	parents []int
}

func newUF(n int) *UF {
	uf := UF{}
	uf.parents = make([]int, n)
	for i := range uf.parents {
		uf.parents[i] = i
	}
	return &uf
}

func (uf *UF) unite(i, j int) {
	uf.parents[uf.find(j)] = uf.find(i)
}

func (uf *UF) find(i int) int {
	for uf.parents[i] != i {
		i = uf.parents[i]
	}
	return i

}

func removeStones(stones [][]int) (ans int) {
	uf := newUF(len(stones))
	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			si := stones[i]
			sj := stones[j]
			if si[0] == sj[0] || si[1] == sj[1] {
				uf.unite(i, j)
			}
		}
	}
	m := map[int]interface{}{}
	for i := range uf.parents {
		m[uf.find(i)] = nil
	}
	return len(stones) - len(m)
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 0}}))
	fmt.Println(removeStones([][]int{{0, 1}, {1, 0}, {1, 1}}))
}
