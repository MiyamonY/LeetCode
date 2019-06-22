package main

import (
	"strings"
)

func binary(n int) string {
	s := ""
	for i := 64; i >= 0; i-- {
		if n&(1<<uint(i)) != 0 {
			s += "1"
		} else {
			s += "0"
		}
	}
	return strings.TrimLeft(s, "0")
}

func queryString(S string, N int) bool {
	end := N >> 1
	for N >= end {
		if !strings.Contains(S, binary(N)) {
			return false
		}
		N--
	}
	return true
}

func main() {
	queryString("0110", 3)
}
