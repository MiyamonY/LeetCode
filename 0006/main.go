package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	charss := make([][]byte, numRows)
	for i := range charss {
		charss[i] = make([]byte, 0)
	}

	for i := range s {
		c := s[i]
		row := i % (numRows + numRows - 2)
		if row >= numRows {
			row = 2*numRows - 2 - row
		}
		charss[row] = append(charss[row], c)
	}

	ret := ""
	for _, chars := range charss {
		for _, c := range chars {
			ret += string(c)
		}
	}
	return ret
}

func main() {
	var s string
	var n int
	fmt.Scan(&s)
	fmt.Scan(&n)
	fmt.Println(convert(s, n))
}
