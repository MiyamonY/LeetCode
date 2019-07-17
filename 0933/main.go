// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Thu Jul 18 01:17:07 2019
//
package main

type RecentCounter struct {
	fifo []int
}

func Constructor() RecentCounter {
	return RecentCounter{fifo: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.fifo = append(this.fifo, t)
	for t-3000 > this.fifo[0] {
		this.fifo = this.fifo[1:]
	}
	return len(this.fifo)
}

func main() {

}
