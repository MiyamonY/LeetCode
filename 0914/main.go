package main

func hasGroupsSizeX(deck []int) bool {
	m := map[int]int{}
	for _, v := range deck {
		m[v]++
	}
	nums := make([]int, 0)

	for _, v := range m {
		nums = append(nums, v)
	}

	for i := 2; i <= 1000/2; i++ {
		valid := true
		for _, num := range nums {
			if num%i != 0 {
				valid = false
			}
		}
		if valid {
			return true
		}
	}
	return false
}

func main() {

}
