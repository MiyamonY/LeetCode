// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 12:42:10 2019
//
package main

import (
	"math"
	"sort"
)

type point struct {
	dist  float64
	point []int
}

func kClosest(points [][]int, K int) (ret [][]int) {
	ps := []point{}
	for _, p := range points {
		ps = append(ps, point{math.Sqrt(float64(p[0]*p[0] + p[1]*p[1])), p})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].dist < ps[j].dist
	})
	for i := 0; i < K; i++ {
		ret = append(ret, ps[i].point)
	}
	return
}

func main() {

}
