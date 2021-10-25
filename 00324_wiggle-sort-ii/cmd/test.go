package main

import "fmt"

func wiggleSort(nums []int)  {
	if len(nums) <= 1 {
		return
	}

	/**
	寻找中位数要算在前半部分
	1 [2] 3 4
	1 2 [3] 4 5
	 */
	n := len(nums)
	target:=0
	left:=0
	right:= n-1
	mid:= (n + 1) / 2 - 1
	for target != mid {
		ol := left
		or:= right
		targetVal := nums[target]
		for left < right{
			for left < right && nums[right] >= targetVal {
				right--
			}
			nums[right],nums[left] = nums[left],nums[right]

			for left < right && nums[left] <= targetVal {
				left++
			}
			nums[right],nums[left] = nums[left],nums[right]
		}
		target = left
		if target == mid{
			break
		}else if target < mid{
			left = left + 1
			right = or
		}else{
			right = left - 1
			left = ol
		}
		target = left
	}
	fmt.Println(target)

	//虚地址 + 三分排序
	var V func (i int) int
	V = func(i int) int {
		return (2 * i + 1) % (n|1)
	}
	mid = nums[target]
	target = 0
	left = 0
	right = n-1

	for left <= right && target < n{
		if nums[V(target)] > mid {
			nums[V(left)],nums[V(target)] = nums[V(target)],nums[V(left)]
			left++
			target++
		}else if nums[V(target)] < mid{
			nums[V(target)],nums[V(right)] = nums[V(right)],nums[V(target)]
			right--
		}else{
			target++
		}
	}

}


func main(){
	nums:= []int {1,3,2,2,3,1}
	wiggleSort(nums)
	fmt.Println(nums)

}
