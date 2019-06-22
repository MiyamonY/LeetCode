package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			p := stack[len(stack)-1]
			if (p == '(' && c != ')') || (p == '{' && c != '}') || (p == '[' && c != ']') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	var S string
	fmt.Scan(&S)
	fmt.Println(isValid(S))
}
