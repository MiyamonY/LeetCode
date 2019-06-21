package main

func removeDuplicates(nums []int) int {
	n := len(nums)

	old := -1
	num := 0
	duplicates := 0
	index := 0
	for _, n := range nums {
		if old != n {
			if 2 < num {
				duplicates += num - 2
				nums[index] = old
				index++
				nums[index] = old
				index++
			} else {
				for i := 0; i < num; i++ {
					nums[index] = old
					index++
				}
			}
			num = 0
			old = n
		}
		num++
	}

	if 2 < num {
		duplicates += num - 2
		nums[index] = old
		index++
		nums[index] = old
		index++
	} else {
		for i := 0; i < num; i++ {
			nums[index] = old
			index++
		}
	}

	return n - duplicates
}

func main() {
}
