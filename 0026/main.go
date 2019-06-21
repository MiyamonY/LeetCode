package main

func removeDuplicates(nums []int) int {
	old := -1 << 60
	index := 0
	num := 0
	for _, n := range nums {
		if n != old {
			nums[index] = n
			num++
			index++
		}
		old = n
	}

	if index > 0 && nums[index-1] != nums[len(nums)-1] {
		nums[index] = nums[len(nums)-1]
		num++
	}
	return num
}

func main() {
}
