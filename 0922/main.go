// Package main provides
//
// File:  main.go
// Author: ymiyamoto
//
// Created on Mon Jul 22 22:24:56 2019
//
package main

func sortArrayByParityII(A []int) []int {
    ans := make([]int, len(A))
    even := 0
    odd := 1
    for _, a := range A{
        if a %2 == 0{
            ans[even] = a
            even+=2
        }else {
            ans[odd] = a
            odd += 2
        }
    }
    return ans
}


func main(){

}
