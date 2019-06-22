package main

import "sort"

func twoCitySchedCost(costs [][]int) (ret int) {
	var diff [][]int
	for _, cost := range costs {
		diff = append(diff, []int{cost[0] - cost[1], cost[0], cost[1]})
	}
	sort.Slice(diff, func(i, j int) bool {
		return diff[i][0] < diff[j][0]
	})
	for i, _ := range costs {
		if i < len(costs)/2 {
			ret += diff[i][1]
		} else {
			ret += diff[i][2]
		}
	}
	return
}

func main() {
}
