package main

import "fmt"

func generateParenthesis(n int) (ret []string) {
	for i := 0; i < 1<<uint8(2*n); i++ {
		valid := true
		parens := ""
		sum := 0
		for j := 0; j < 2*n; j++ {
			if i&(1<<uint8(j)) != 0 {
				sum++
				parens += "("
			} else {
				sum--
				parens += ")"
			}

			if sum < 0 {
				valid = false
			}
		}
		if sum > 0 || !valid {
			continue
		}

		ret = append(ret, parens)
	}
	return ret
}

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Println(generateParenthesis(n))
}
