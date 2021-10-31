package main

import (
	"fmt"
	"math"
)


func longestIncreasingPath1(matrix [][]int) int {
	row,col:= len(matrix),len(matrix[0])
	cache:= make([][]int,row)
	for i:=0;i<row;i++{
		cache[i] = make([]int ,col)
	}

	ans:=math.MinInt32
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			ans = max1(ans,dfs1(i,j,matrix,cache))
		}
	}
	return ans;

}


func dfs1(i,j int,matrix [][]int,cache [][]int)int{
	if cache[i][j] > 0{
		return  cache[i][j]
	}
	tmp := 1
	// 上
	if getTop(i,j,matrix) > matrix[i][j]{
		tmp = max1(tmp,dfs1(i-1,j,matrix,cache)+1)
	}
	// 下
	if getBottom(i,j,matrix) > matrix[i][j]{
		tmp = max1(tmp,dfs1(i+1,j,matrix,cache)+1)
	}
	//左
	if getLeft(i,j,matrix) > matrix[i][j]{
		tmp = max1(tmp,dfs1(i,j-1,matrix,cache)+1)
	}

	//右
	if getRight(i,j,matrix) > matrix[i][j]{
		tmp = max1(tmp,dfs1(i,j+1,matrix,cache)+1)
	}
	cache[i][j] = tmp
	return tmp
}


func getTop(i,j int,matrix [][]int)int{
	if  matrix[i][j] ==  math.MinInt32{
		return 0
	}
	if i == 0 {
		return math.MinInt32
	}
	return matrix[i-1][j]
}
func getBottom(i,j int,matrix [][]int)int{
	if i == len(matrix) - 1{
		return math.MinInt32
	}
	return matrix[i+1][j]

}
func getLeft(i,j int,matrix [][]int)int{
	if j == 0{
		return math.MinInt32
	}
	return matrix[i][j-1]

}
func getRight(i,j int,matrix [][]int)int{
	if j == len(matrix[0]) - 1{
		return math.MinInt32
	}
	return matrix[i][j+1]
}

func max1(i,j int) int{
	if i> j{
		return i
	}
	return j
}


func main(){
	matrix :=[][]int {{9,9,4},{6,6,8},{2,1,1}}
	fmt.Println(longestIncreasingPath(matrix))
}
