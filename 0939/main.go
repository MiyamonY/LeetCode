// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Wed Jul 17 01:17:50 2019
//
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minAreaRect(points [][]int) int {
	ans := 1 << 60
	m := map[int][]int{}
	for i := range points {
		x, y := points[i][0], points[i][1]
		if _, ok := m[x]; ok {
			m[x] = append(m[x], y)
		} else {
			m[x] = []int{y}
		}
	}

	for x0, ys0 := range m {
		for x1, ys1 := range m {
			if x1 >= x0 {
				continue
			}
			for i, y00 := range ys0 {
				for j := i + 1; j < len(ys0); j++ {
					y01 := ys0[j]
					found := 0
					for _, y1 := range ys1 {
						if y1 == y00 || y1 == y01 {
							found++
						}
						if found == 2 {
							ans = min(ans, abs(x0-x1)*abs(y00-y01))
						}
					}
				}
			}
		}
	}

	if ans == 1<<60 {
		return 0
	}
	return ans
}

func main() {

}
