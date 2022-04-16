package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestPalindrome(2))
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	max := int(math.Pow10(n)) - 1
	for i := max; i > 0; i-- {
		tmp := i
		// 构造回文
		for j := i; j > 0; j /= 10 {
			tmp = tmp*10 + j%10
		}

		for k := max; k*k >= tmp; k-- {
			if tmp%k == 0 {
				return tmp % 1337
			}
		}
	}
	return 0
}
