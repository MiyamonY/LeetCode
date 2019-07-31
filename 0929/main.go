// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sat Jul 20 06:51:16 2019
//
package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := map[string]map[string]int{}
	for i := range emails {
		email := strings.Split(emails[i], "@")
		domain := email[1]
		local := strings.Split(email[0], "+")[0]
		trimed := strings.Replace(local, ".", "", -1)
		if _, ok := m[domain]; !ok {
			m[domain] = map[string]int{}
		}
		m[domain][trimed]++
	}
	fmt.Printf("%+v", m)
	ans := 0
	for _, v := range m {
		ans += len(v)
	}
	return ans
}

func main() {

}
