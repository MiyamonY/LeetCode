package main

func findDuplicate(nums []int) int {
	found := make([]bool, len(nums))
	for _, n := range nums {
		if found[n] {
			return n
		}
		found[n] = true
	}
	return 0
}
