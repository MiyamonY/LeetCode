package main

import (
	"encoding/json"
	"fmt"
)

func rec(nums []int, num int) (ret [][]int) {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	ret = rec(nums[1:], num)
	n := nums[0]
	if num <= n {
		for _, r := range rec(nums[1:], n) {
			ret = append(ret, append([]int{n}, r...))
		}
	}
	return ret
}

func filter(nums [][]int) (ret [][]int) {
	for _, num := range nums {
		if len(num) >= 2 {
			ret = append(ret, num)
		}

	}
	return
}

func unique(nums [][]int) (ret [][]int) {
	m := map[string]interface{}{}

	for _, num := range nums {
		j, _ := json.Marshal(num)
		m[string(j)] = nil
	}
	for k := range m {
		var num []int
		json.Unmarshal([]byte(k), &num)
		ret = append(ret, num)
	}

	return
}

func findSubsequences(nums []int) (ret [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}

	for _, r := range rec(nums[1:], nums[0]) {
		ret = append(ret, append([]int{nums[0]}, r...))
	}

	for _, r := range findSubsequences(nums[1:]) {
		ret = append(ret, r)
	}
	return unique(filter(ret))
}

// main ...
func main() {
	fmt.Println(findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
}
