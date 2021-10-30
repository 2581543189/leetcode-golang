package main

import (
	"fmt"
)

// 二分查找
func findDuplicate1(nums []int) int {
	l,r := 0,len(nums) - 1
	for l < r{
		mid := l + (r - l) >> 1
		count:= 0
		for _,val := range nums{
			if val <= mid{
				count ++
			}
		}
		if count <= mid {
			l = mid + 1
		}else {
			r = mid
		}
	}
	return  l
}

// 位运算
func findDuplicate2(nums []int) int {
	// 获取n 的位数？
	n := len(nums) - 1
	bit_max:=31
	for n >> bit_max == 0{
		bit_max--
	}
	bit_max++
	ans:=0
	for i:=0;i<bit_max;i++{
		x:=0
		y:=0
		for index,val:=range nums{
			x += val & ( 1 << i)
			y += index & ( 1 << i)
		}
		if x > y{
			ans += ( 1 << i)
		}
	}

	return ans
}

//快慢指针
func findDuplicate3(nums []int) int {

	s,f:= nums[0],nums[nums[0]]
	for f != s{
		s = nums[s]
		f = nums[nums[f]]
	}
	s=0
	for f != s{
		f = nums[f]
		s = nums[s]
	}

	return f

}

func main(){
	nums:=[]int{1,3,4,2,2}
	fmt.Println(findDuplicate1(nums))
	fmt.Println(findDuplicate2(nums))
	fmt.Println(findDuplicate3(nums))
}
