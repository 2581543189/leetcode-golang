package main

import (
	"fmt"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxProduct(nums []int) int {
	n:= len(nums)
	cacheMax := make([] int,n)
	cacheMin := make([] int,n)
	for i:=0;i<n;i++ {
		cacheMax[i] = INT_MAX
		cacheMin[i] = INT_MIN
	}

	var dfs func(i int,max bool) int
	dfs = func(i int,max bool) int{
		if i >= n - 1 {
			return nums[n-1]
		}
		ret := nums[i]
		var tmp int
		if(max){
			// 最大值逻辑
			if cacheMax[i] < INT_MAX{
				return cacheMax[i]
			}
			if nums[i] > 0{
				tmp = nums[i] * dfs(i+1,true)
			}else if nums[i] < 0{
				tmp = nums[i] * dfs(i+1,false)
			}
			if ret < tmp {
				ret = tmp
			}
			cacheMax[i] = ret
		}else{
			// 最小值逻辑
			if cacheMin[i] > INT_MIN{
				return cacheMin[i]
			}
			if nums[i] > 0{
				tmp = nums[i] * dfs(i+1,false)
			}else if nums[i] < 0{
				tmp = nums[i] * dfs(i+1,true)
			}
			if ret > tmp {
				ret = tmp
			}
			cacheMin[i] = ret
		}

		return ret
	}
	max := INT_MIN
	for i:=0;i<n;i++{
		tmp:= dfs(i,true)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main() {

	nums := [] int {-1,-2,-9,-6}
	fmt.Println(maxProduct(nums))
}

