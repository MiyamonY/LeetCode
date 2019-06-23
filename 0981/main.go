// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Sun Jun 23 21:34:18 2019
//
package main

import "sort"

type TimeMap struct {
	values     map[string][]string
	timestamps map[string][]int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	values := map[string][]string{}
	timestamps := map[string][]int{}
	return TimeMap{values, timestamps}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.values[key]; ok {
		this.values[key] = append(this.values[key], value)
		this.timestamps[key] = append(this.timestamps[key], timestamp)
	} else {
		this.values[key] = []string{value}
		this.timestamps[key] = []int{timestamp}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timestamps := this.timestamps[key]
	index := sort.Search(len(timestamps), func(i int) bool {
		return timestamp < timestamps[i]
	})
	if index == 0 {
		return ""
	}
	return this.values[key][index-1]
}

func main() {

}
