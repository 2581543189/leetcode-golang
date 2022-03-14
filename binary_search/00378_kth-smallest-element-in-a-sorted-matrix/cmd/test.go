package main

import "fmt"

func kthSmallest1(matrix [][]int, k int) int {
	n:=len(matrix)
	l,r := matrix[0][0], matrix[n - 1][n - 1]
	for l < r {
		mid := l + (r - l) >> 1

		// 统计有多少个元素小于mid
		x,y:= n-1,0
		count :=0
		for x >=0 && y< n{

			// 为什么 == 要 y++?
			//当前统计的是，小于等于mid 的数字有多少个
			//如果往上走，右边的数字有可能还是 mid，就被漏掉了
			// 如果往右走，上边的数字有可能还是 mid，但是没有漏掉
			if matrix[x][y] <= mid{
				count += (x + 1)
				y++
			}else{
				x--
			}
		}
		// 为什么 >= k?
		// 小于等于mid 的数字有k 个的时候，需要继续在左边计算
		if count >= k {
			r = mid
		}else{
			l = mid + 1
		}
	}
	return l
}


func main()  {
	matrix := [][] int {{-5,-4},{-5,-4}}
	fmt.Println(kthSmallest1(matrix,2))
}
