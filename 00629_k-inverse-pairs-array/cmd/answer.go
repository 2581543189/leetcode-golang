package main

import "fmt"

func min(a, b int) int {
	if a <b{
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}
func kInversePairs(n int, k int) int {
	dp := make([]int, k+1)
	sum := make([]int, k+2)
	dp[0] = 1
	sum[1]=1
	for i := 1; i <= n; i++ {
		for j := min(k, i*(i-1)/2); j >0; j-- {
			dp[j] = sum[j+1]-sum[max(0, j-i+1)]
			dp[j] %=1e9+7
		}
		for j := 1; j <= k; j++ {
			sum[j+1] = sum[j]+dp[j]
		}
	}
	return dp[k]
}

func main(){
	fmt.Println(kInversePairs(1000,1000))
}
