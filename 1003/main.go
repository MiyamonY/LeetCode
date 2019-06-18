package main

import (
	"fmt"
	"strings"
)

func isValid(S string) bool {
	old := S
	for {
		new := strings.Replace(old, "abc", "", -1)
		if old == new {
			return false
		}
		if len(new) == 0 {
			return true
		}
		old = new
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(isValid(s))
}
