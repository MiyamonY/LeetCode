package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	area := 0

	left, right := 0, len(height)-1
	for left < right {
		area = max(area, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func main() {
}
