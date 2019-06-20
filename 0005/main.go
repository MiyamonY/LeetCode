package main

import "fmt"

func isPalidrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) (ret string) {
	maxLen := 0

	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			if isPalidrome(s[i:j]) {
				n := len(s[i:j])
				if maxLen < n {
					maxLen = n
					ret = s[i:j]
				}
			}
		}
	}
	return
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestPalindrome(s))
}
