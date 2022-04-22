package main

import "fmt"

func main() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
}

func maxRotateFunction(nums []int) int {
	numSum := 0
	for _, v := range nums {
		numSum += v
	}
	f := 0
	for i, num := range nums {
		f += i * num
	}
	ans := f
	for i := len(nums) - 1; i > 0; i-- {
		f += numSum - len(nums)*nums[i]
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
