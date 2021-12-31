package main

import "fmt"

func main(){
	fmt.Println(checkPerfectNumber1(28))
}

func checkPerfectNumber1(num int) bool {
	if num <= 2{
		return false
	}
	sum:=1
	left,right:=2,num
	for left < right{
		if num % left == 0{
			sum+=left
			sum+=(num / left)
			right = num / left
			if sum > num{
				return false
			}
		}
		left++
	}
	return sum == num
}
