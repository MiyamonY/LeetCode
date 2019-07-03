// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul  4 01:58:40 2019
//
package main

import "sort"

func sampleStats(count []int) []float64 {
	m := map[int]int{}

	for i, c := range count {
		m[i] = c
	}

	keys := []int{}
	total := 0
	counts := 0
	max := 0
	mode := 0
	for k, v := range m {
		if v > 0 {
			keys = append(keys, k)
		}
		total += k * v
		if max < v {
			mode = k
			max = v
		}
		counts += v
	}
	sort.Ints(keys)

	median := 0.0
	sum := 0
	for i, c := range count {
		sum += c
		if counts/2 < sum {
			median = float64(i)
			break
		} else if counts/2 == sum {
			for j, v := range count[i+1:] {
				if v > 0 {
					median = float64(2*i+j+1) / 2.0
					break
				}
			}
			break
		}
	}
	return []float64{float64(keys[0]), float64(keys[len(keys)-1]), float64(total) / float64(counts),
		median, float64(mode)}
}

func main() {

}
