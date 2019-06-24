// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Tue Jun 25 00:13:14 2019
//
package main

type pos struct {
	x int
	y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	dists := make([][]int, len(grid))
	for i := range dists {
		dists[i] = make([]int, len(grid[0]))
		for j := range dists[i] {
			dists[i][j] = 1 << 30
		}
	}

	dists[0][0] = 0
	queue := []pos{{0, 0}}
	for 0 < len(queue) {
		top := queue[0]
		queue = queue[1:]
		for _, x := range []int{1, 0, -1} {
			for _, y := range []int{1, 0, -1} {
				if x == 0 && y == 0 {
					continue
				}
				dx := top.x + x
				dy := top.y + y
				if 0 <= dx && dx < len(grid) && 0 <= dy && dy < len(grid[0]) &&
					grid[dx][dy] == 0 && dists[top.x][top.y]+1 < dists[dx][dy] {
					dists[dx][dy] = dists[top.x][top.y] + 1
					queue = append(queue, pos{dx, dy})
				}
			}
		}
	}

	if grid[0][0] == 1 ||
		grid[len(grid)-1][len(grid[0])-1] == 1 || dists[len(grid)-1][len(grid[0])-1] == 1<<30 {
		return -1
	}
	return dists[len(grid)-1][len(grid[0])-1] + 1
}

func main() {

}
