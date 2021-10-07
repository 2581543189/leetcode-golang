package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	//var s string = "ac"
	//var t string = "bb"
	//fmt.Println(isAnagram(s,t))
	nums := [] int {1,2,3,4,5,6,7}
	rotate(nums,3)
	fmt.Println(json.Marshal(nums))
}

func rotate(nums []int, k int)  {
	n:=len(nums)

	first := nums[0]
	second := nums[0]
	for i := 0; i < n - 1; i++{
		second = nums[(i + 1)  * k % n]
		nums[(i + 1) * k % n] = first
		first = second
	}
	nums[0]=first
}

func isAnagram(s string, t string) bool {
	s1:= []byte(s)
	t1:= []byte(t)
	cache:= make(map[byte] int64)

	for _,ch := range s1{
		cache[ch]++
	}
	for _,ch := range t1{
		cache[ch]--
	}

	for key,_ := range cache{
		if cache[key] != 0{
			return false
		}
	}
	return true
}