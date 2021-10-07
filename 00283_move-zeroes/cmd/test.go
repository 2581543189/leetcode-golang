package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main(){
	nums := []int{1,0,2}
	moveZeroes(nums)
	fmt.Printf("%v",nums)
}
