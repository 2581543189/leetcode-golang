package main

import "fmt"

func main(){
	nums:=[]int{1,1,2,3,3,4,4,8,8}
	fmt.Println(singleNonDuplicate1(nums))
}


func singleNonDuplicate1(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	l,r:= 0, len(nums) - 1
	for r - l > 2{
		mid := l + (r - l) >> 1
		if mid % 2 == 1{
			if nums[mid] == nums[mid + 1]{
				r = mid
			}else if nums[mid] == nums[mid - 1]{
				l = mid + 1
			}else{
				return nums[mid]
			}
		}else{
			if nums[mid] == nums[mid + 1]{
				l = mid
			}else if nums[mid] == nums[mid - 1]{
				r = mid
			}else{
				return nums[mid]
			}
		}

	}
	if nums[l] == nums[l + 1]{
		return nums[r]
	}
	return nums[l]
}