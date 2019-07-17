// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul 18 00:24:08 2019
//
package main

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	letters := []string{}
	digits := []string{}
	for i := range logs {
		ss := strings.Split(logs[i], " ")
		if '0' <= ss[1][0] && ss[1][0] <= '9' {
			digits = append(digits, logs[i])
		} else {
			letters = append(letters, logs[i])
		}
	}
	sort.Slice(letters, func(i, j int) bool {
		si := strings.SplitN(letters[i], " ", 2)
		sj := strings.SplitN(letters[j], " ", 2)
		if si[1] == sj[1] {
			return si[0] < sj[0]
		}
		return si[1] < sj[1]
	})
	return append(letters, digits...)
}

func main() {

}
