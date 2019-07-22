// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Mon Jul 22 23:26:04 2019
//
package main

func minAddToMakeValid(S string) int {
    ans := 0
    count := 0
    for i := range S{
        if S[i] == '('{
            count++
        } else {
            if count > 0{
                count--
            }
        }
    }
    ans += count
    count = 0
    for i:= len(S)-1;i>=0;i--{
        if S[i] == ')'{
            count++
        } else {
            if count > 0{
                count--
            }
        }
    }
    return ans + count
}

func main(){

}
