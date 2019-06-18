package main

import "fmt"

func rec(candidates []int, acc []int, n int) (ret [][]int) {
	if n == 0 {
		return [][]int{acc}
	}

	if n < 0 || len(candidates) == 0 {
		return [][]int{}
	}

	acc0 := append(acc, candidates[0])
	ac := make([]int, len(acc0))
	copy(ac, acc0)
	ret1 := rec(candidates, ac, n-candidates[0])
	if len(candidates) > 0 {
		ret2 := rec(candidates[1:], acc, n)
		ret = append(ret1, ret2...)
	}

	return ret
}

func combinationSum(candidates []int, target int) [][]int {
	return rec(candidates, []int{}, target)
}

func main() {
	vs := make([]int, 3)
	for i := range vs {
		fmt.Scan(&vs[i])
	}
	var n int
	fmt.Scan(&n)
	fmt.Printf("%+v\n", combinationSum(vs, n))
}
