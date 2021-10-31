package main

import (
	"fmt"
)

// 动态规划
func lengthOfLIS3(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//贪心法
func lengthOfLIS4(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		if len(dp) ==0 || dp[len(dp) - 1] < num {
			dp = append(dp, num)
		} else {
			l, r := 0, len(dp) - 1
			pos := r
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] >= num {
					pos = mid;
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = num
		}//二分查找
	}
	return len(dp)
}
func main(){
	nums:=[]int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS3(nums))
	fmt.Println(lengthOfLIS4(nums))
}
