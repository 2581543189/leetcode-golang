package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	x := m - 1
	y := 0

	for x > 0 && y < n {
		if matrix [x][y] == target {
			return  true
		}

		if matrix [x][y] > target {
			x--
		}else{
			y++
		}
	}
	return false
}

func main()  {

	matrix := [][]int{
		{1,     4,  7,  11, 15},
		{2,     5,  8,  12, 19},
		{3,     6,  9,  16, 22},
		{10,    13, 14, 17, 24},
		{18,    21, 23, 26, 30},
	}
	target := 20
	fmt.Println(searchMatrix(matrix,target))
}