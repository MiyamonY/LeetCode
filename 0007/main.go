package main

import (
	"math"
)

func reverse(x int) int {
	d := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		d = d*10 + pop
	}
	if math.MaxInt32 < d {
		return 0
	}
	if d < math.MinInt32 {
		return 0
	}

	return d
}

func main() {
}
