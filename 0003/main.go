package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) (ret int) {
	start := 0
	m := map[byte]interface{}{}
	for end := range s {
		c := s[end]
		_, ok := m[c]
		for ok {
			delete(m, s[start])
			start++
			_, ok = m[c]
		}
		m[c] = nil
		if ret < len(m) {
			ret = len(m)
		}
	}
	return
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(lengthOfLongestSubstring(s))
}
