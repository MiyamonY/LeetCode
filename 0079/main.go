package main

import (
	"fmt"
)

func dfs(board [][]byte, visited [][]bool, x, y int, word string) bool {
	if len(word) == 0 || (len(word) == 1 && word[0] == board[x][y]) {
		return true
	}

	if visited[x][y] || word[0] != board[x][y] {
		return false
	}

	visited[x][y] = true
	ret := false
	for _, d := range []struct {
		x int
		y int
	}{{1, 0}, {-1, 0}, {0, -1}, {0, 1}} {
		s := x + d.x
		t := y + d.y
		if 0 <= s && s < len(board) && 0 <= t && t < len(board[0]) && !visited[s][t] {
			ret = ret || dfs(board, visited, s, t, word[1:])
		}
	}
	visited[x][y] = false
	return ret
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			visited := make([][]bool, len(board))
			for i := range visited {
				visited[i] = make([]bool, len(board[0]))
			}
			if dfs(board, visited, i, j, word) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'}}, "abcced"))
}
