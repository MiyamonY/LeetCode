package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func totalFruit(tree []int) int {
	ans, i := 0, 0

	m := map[int]int{}
	for j, f := range tree {
		m[f]++
		for len(m) >= 3 {
			m[tree[i]]--
			if m[tree[i]] == 0 {
				delete(m, tree[i])
			}
			i++
		}
		ans = max(ans, j-i+1)

	}
	return ans
}

func main() {
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}
