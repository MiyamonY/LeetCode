package main

import "fmt"

func dfs(grid [][]byte, visited [][]int, x, y int, no int) {
	if visited[x][y] != 0 || grid[x][y] == '0' {
		return
	}

	visited[x][y] = no
	for _, d := range []struct {
		x int
		y int
	}{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
		s := x + d.x
		t := y + d.y
		if 0 <= s && s < len(grid) && 0 <= t && t < len(grid[0]) {
			dfs(grid, visited, s, t, no)
		}
	}
}

func numIslands(grid [][]byte) int {
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
	}

	no := 1
	for i := range grid {
		for j := range grid[i] {
			dfs(grid, visited, i, j, no)
			no++
		}
	}

	m := map[int]interface{}{}
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] != 0 {
				m[visited[i][j]] = nil
			}

		}
	}
	return len(m)
}

func main() {
	fmt.Println(numIslands(
		[][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}))
}
