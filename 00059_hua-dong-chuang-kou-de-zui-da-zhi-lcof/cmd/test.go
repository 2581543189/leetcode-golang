package main

import (
	"fmt"
)


func maxSlidingWindow1(nums []int, k int) []int {
	if k == 0{
		return nums
	}
	n:= len(nums)
	monotonicList := []int {}
	ans := make([]int ,n - k + 1)
	//
	for i:=0 ; i < n ; i++ {
		for len(monotonicList) > 0 && nums[i] > nums[monotonicList[len(monotonicList) - 1]] {
			monotonicList = monotonicList[:len(monotonicList) - 1]
		}
		monotonicList = append(monotonicList,i)
		if i == k-1 {
			ans[0] = nums[monotonicList[0]]
		}
		if i > k - 1 {
			for monotonicList[0] <= i - k {
				monotonicList = monotonicList[1:len(monotonicList)]
			}
			ans[i - k + 1] = nums[monotonicList[0]]
		}
	}

	return  ans
}

func main()  {
	nums := [] int {1,-1}
	fmt.Println(maxSlidingWindow1(nums,1))
}