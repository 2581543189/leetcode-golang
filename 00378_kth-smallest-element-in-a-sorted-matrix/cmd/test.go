package main

import "fmt"

func kthSmallest1(matrix [][]int, k int) int {
	left,right := matrix[0][0] , matrix[len(matrix) - 1][len(matrix[0]) - 1]
	x := 0
	for left < right {
		x = left + (right - left) / 2
		// 计算 <= x 的元素数量
		m,n := len(matrix) - 1,0
		count := 0
		for m >= 0 && n < len(matrix[0]) {
			if matrix[m][n] <= x {
				count += m + 1
				n++
			}else{
				m--
			}
		}
		if count < k {
			left = x + 1
		}else{
			right = x
		}
	}
	return left
}

func main()  {
	matrix := [][] int {{-5,-4},{-5,-4}}
	fmt.Println(kthSmallest1(matrix,2))
}
