package main

import "fmt"

func main(){
	// 这个写法超时了
	fmt.Println(gridIllumination1( 5,[][]int{{0,0},{4,4}},[][]int{{1,1},{1,0}}))
}

func gridIllumination1(n int, lamps [][]int, queries [][]int) []int {
	ans:=make([]int,0)
	for i := range queries {
		ans = append(ans,getResult(n,lamps,queries[i][0],queries[i][1]))
		lamps = off(n,lamps,queries[i][0],queries[i][1])
	}
	return ans
}

// 判断当前是否被照亮
func getResult(n int, lamps [][]int,x int,y int )int{
	// 参数校验
	if x<0 || x >=n || y<0 || y >=n{
		return 0
	}

	for i :=range lamps{
		if x == lamps[i][0] || y == lamps[i][1]{
			// 同一行，同一列 直接返回true
			return 1
		}
		if abs(lamps[i][0] - x) == abs(lamps[i][1] - y){
			// 同一对角线
			return 1
		}
	}
	return 0
}
// 关闭灯
func off(n int, lamps [][]int,x int,y int )[][]int{
	// 参数校验
	if x<0 || x >=n || y<0 || y >=n{
		return lamps
	}
	ans:=make([][]int,0)
	for i :=range lamps{
		if abs(lamps[i][0] - x) <=1 &&  abs(lamps[i][1] - y) <=1 {
			// 同一对角线
			continue
		}
		ans = append(ans,lamps[i])
	}
	return ans;
}

func abs(i int)int{
	if i < 0{
		return -i
	}
	return i
}
