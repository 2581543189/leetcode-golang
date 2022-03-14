package main

import (
	"fmt"
	"math"
)


// 动态规划
// cache[i] 表示 以i开头的 最长升序子串的长度
// 如果 nums[i - 1] < nums[i] 那么 nums[i - 1] = nums[i] + 1
// 如果 nums[i - 1] >= nums[i] 那么 需要从 i ~ n-1 去找，有没有 nums[i - 1] < nums[i] 的。获取最大的一项，作为 nums[i - 1]
func lengthOfLIS1(nums []int) int {
	n:= len(nums)
	if n<=1{
		return 1
	}
	cache:= make([]int,n)
	cache[n-1] = 1
	ans:=math.MinInt32
	for i:= n-2; i>=0 ; i--{
		max:=math.MinInt32
		for j:=i+1; j< n; j++{
			if nums[j] > nums[i]{
				max = max2(max,1 +  cache[j])
			}else{
				max = max2(max,1)
			}
		}
		cache[i] = max
		ans = max2(ans,max)
	}
	return ans
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//贪心法
func lengthOfLIS2(nums []int) int {
	n:= len(nums)
	if n<=1{
		return 1
	}
	cache:=[]int{nums[n-1]}
	for i:= n-2; i>=0; i--{
		if nums[i] < cache[len(cache) - 1]{
			cache=append(cache,nums[i])
		}else{
			l,r := 0,len(cache)-1
			for l < r{
				mid := l + (r - l) >> 1
				if cache[mid] > nums[i]{
					l = mid +1
				}else {
					r = mid
				}
			}
			cache[l] = nums[i]
		}
	}
	return  len(cache)
}
func main(){
	nums:=[]int{1,3,6,7,9,4,10,5,6}
	fmt.Println(lengthOfLIS1(nums))
	fmt.Println(lengthOfLIS2(nums))
}
