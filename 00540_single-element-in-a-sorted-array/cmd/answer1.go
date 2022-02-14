package main

import "fmt"

func main(){
	nums:=[]int{1,1,2,3,3,4,4,8,8}
	fmt.Println(singleNonDuplicate2(nums))
}


func singleNonDuplicate2(nums []int) int {
	l,r:= 0,len(nums) - 1
	for l < r{
		mid:= l + (r - l ) >> 1
		/**
		 * 奇数 会 - 1
		 * 偶数 会 + 1
		 */
		if nums[mid] == nums[mid ^ 1]{
			l = mid + 1
		}else{
			r = mid
		}
	}
	return  nums[l]
}
