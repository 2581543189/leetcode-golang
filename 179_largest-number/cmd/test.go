package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber1(nums []int) string {

	if len(nums) == 0{
		return ""
	}

	sort.Slice(nums, func(i, j int) bool {
		x,y := nums[i],nums[j]
		mx,my := 10,10
		for mx <= x {
			mx = mx * 10
		}
		for my <= y {
			my = my * 10
		}

		return  x * my + y > y * mx + x
	})
	ans := []byte{}
	firstNotZero := false
	for _,val := range nums{
		if val != 0{
			firstNotZero = true
		}
		if !firstNotZero && val == 0{
			continue
		}
		ans = append(ans, strconv.Itoa(val)...)
	}
	ret := string(ans)
	if ret == ""{
		ret = "0"
	}
	return string(ans)
}

func main(){

	nums:= [] int{ 0,0 }
	fmt.Println(largestNumber1(nums))

}
