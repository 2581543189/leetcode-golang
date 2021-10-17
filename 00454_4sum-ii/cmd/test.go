package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	cache1 := make(map[int]int)
	cache2 := make(map[int]int)

	for i:= 0; i< len(nums1) ; i++ {
		for j:=0 ;j < len(nums2); j++{
			key:= nums1[i] + nums2[j]
			if _,ok := cache1[key] ;ok{
				cache1[key] = 1
			}else{
				cache1[key] = cache1[key] +1
			}
		}
	}

	for i:= 0; i< len(nums3) ; i++ {
		for j:=0 ;j < len(nums4); j++{
			key:= nums3[i] + nums4[j]
			if  _,ok := cache2[key] ;ok{
				cache2[key] = 1
			}else{
				cache2[key] = cache2[key] + 1
			}
		}
	}

	sum := 0
	for key,val := range cache1 {
		if _,ok:=cache2[-key] ;ok {
			sum += val * cache2[-key]
		}
	}
	return sum

}

func main(){
	nums1 := []int{-1,-1}
	nums2 := []int{-1,1}
	nums3 := []int{-1,1}
	nums4 := []int{1,-1}
	fmt.Println(fourSumCount(nums1,nums2,nums3,nums4))

}
