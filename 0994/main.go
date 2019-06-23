// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 09:58:49 2019
//
package main

type pos struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	var queue []pos
	oranges := make([][]int, len(grid))

	for i := range oranges {
		oranges[i] = make([]int, len(grid[0]))
		for j := range oranges[i] {
			oranges[i][j] = 1 << 30
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, pos{i, j})
				oranges[i][j] = 0
			}
		}
	}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		for _, d := range []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			s := top.x + d.x
			t := top.y + d.y
			if 0 <= s && s < len(grid) && 0 <= t && t < len(grid[0]) && grid[s][t] == 1 {
				if oranges[top.x][top.y]+1 < oranges[s][t] {
					oranges[s][t] = oranges[top.x][top.y] + 1
					queue = append(queue, pos{s, t})
				}
			}
		}
	}
	ret := 0
	for i := range oranges {
		for j := range oranges[i] {
			if grid[i][j] == 1 && oranges[i][j] == 1<<30 {
				return -1
			}
			if oranges[i][j] != 1<<30 {
				if ret < oranges[i][j] {
					ret = oranges[i][j]
				}
			}
		}
	}
	return ret
}

func main() {

}
