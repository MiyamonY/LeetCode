package main

import (
	"fmt"
	"strings"
)

func sum(num1 string, num2 string, carry byte) string {
	if len(num1) == 0 {
		if carry > 0 {
			return sum(num2, string(carry+'0'), 0)
		}
		return num2
	}
	if len(num2) == 0 {
		if carry > 0 {
			return sum(num1, string(carry+'0'), 0)
		}
		return num1
	}

	n := num1[0] - '0' + num2[0] - '0' + carry
	return string(n%10+'0') + sum(num1[1:], num2[1:], n/10)
}

func mul(num1 string, n byte) string {
	ret := "0"
	for i := byte(0); i < n; i++ {
		ret = sum(ret, num1, 0)
	}
	return ret
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1 = reverse(num1)
	num2 = reverse(num2)
	ans := "0"
	for i := range num2 {
		n := num2[i] - '0'
		if n == 0 {
			continue
		}
		tmp := strings.Repeat("0", i) + mul(num1, n)
		ans = sum(ans, tmp, 0)
	}

	return reverse(ans)
}

func main() {
	fmt.Printf("%s\n", multiply("0", "52"))
}
