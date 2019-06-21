package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func myAtoi(str string) int {
	s := strings.TrimLeft(str, " ")
	m := regexp.MustCompile(`^([+-])??(\d+)`).FindSubmatch([]byte(s))
	if len(m) == 0 {
		return 0
	}

	d := string(m[2])
	if len(d) == 0 {
		return 0
	}
	d = strings.TrimLeft(d, "0")
	if len(m[1]) == 0 || m[1][0] == '+' {
		if len(d) >= 11 {
			return math.MaxInt32
		}
		i, _ := strconv.ParseInt(d, 10, 64)
		return int(min(i, math.MaxInt32))
	}

	if len(d) >= 11 {
		return math.MinInt32
	}
	i, _ := strconv.ParseInt(d, 10, 64)
	return int(max(-i, math.MinInt32))
}

func main() {
	m := regexp.MustCompile(`^([+-])??(\d+)`).FindSubmatch([]byte("123"))
	fmt.Printf("%+v", m)

	m = regexp.MustCompile(`^([+-])??(\d+)`).FindSubmatch([]byte("+123"))
	fmt.Printf("%+v", m)

	m = regexp.MustCompile(`^([+-])??(\d+)`).FindSubmatch([]byte("w123"))
	fmt.Printf("%+v", m)

	tests := []struct {
		input string
		val   int
	}{{"42", 42}, {"123", 123}, {"+123", 123}, {"-123", -123},
		{"w123", 0}, {"111111111111111111111111", math.MaxInt32}}

	for _, test := range tests {
		if myAtoi(test.input) != test.val {
			fmt.Errorf("error")
		}
	}

}
