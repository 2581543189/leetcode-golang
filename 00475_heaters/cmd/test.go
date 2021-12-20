package main

import (
	"fmt"
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	ans := 0
	// 暖气排序
	sort.Ints(heaters)

	for _, ho := range houses {
		min := binarySearch(ho, heaters)
		if min > ans {
			ans = min
		}
	}
	return ans

}

// 使用二分法查找最近的暖气
func binarySearch(house int, heaters []int) int {
	j := sort.SearchInts(heaters, house+1)
	minDis := math.MaxInt32
	if j < len(heaters) {
		minDis = heaters[j] - house
	}
	i := j - 1
	if i >= 0 && house-heaters[i] < minDis {
		minDis = house - heaters[i]
	}
	return minDis
}

func main(){
	houses := []int{1,2,3}
	heaters := []int{2}
	fmt.Println(findRadius(houses,heaters))

}