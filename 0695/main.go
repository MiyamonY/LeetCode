package main

func dfs(grid [][]int, visited [][]int, x, y, no int) {
	if visited[x][y] != 0 {
		return
	}
	visited[x][y] = no

	for _, d := range []struct {
		x int
		y int
	}{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
		s := x + d.x
		t := y + d.y
		if 0 <= s && s < len(grid) && 0 <= t && t < len(grid[0]) && grid[s][t] == 1 {
			dfs(grid, visited, s, t, no)
		}
	}
}

func maxAreaOfIsland(grid [][]int) (ret int) {
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
	}

	no := 1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && visited[i][j] == 0 {
				dfs(grid, visited, i, j, no)
				no++
			}
		}
	}

	m := map[int]int{}
	for i := range visited {
		for j := range visited[i] {
			m[visited[i][j]]++
		}
	}

	for k, v := range m {
		if k != 0 {
			if ret < v {
				ret = v

			}
		}
	}
	return
}

func main() {
}
