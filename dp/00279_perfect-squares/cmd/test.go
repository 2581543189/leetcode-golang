package main

import (
	"fmt"
	"math"
)

func numSquares1(n int) int {
	cache:= make([]int,n+1)
	var dfs func(n int)int
	dfs = func(n int) int{
		if n == 0{
			return 0
		}
		if cache[n] > 0 {
			return cache[n]
		}
		min := math.MaxInt32
		maxi:= 1
		for ;maxi*maxi <= n; maxi++{

		}
		if maxi*maxi > n{
			maxi--
		}
		for i:=maxi; i >= 1; i--{
			tmp := 1 + dfs(n - i*i)
			if tmp < min{
				min = tmp
			}
		}
		cache[n] = min
		return min
	}
	return dfs(n)
}

func main(){
	fmt.Println(numSquares1(48))
}
