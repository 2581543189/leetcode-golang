package main

import "fmt"

func main(){
	points:= [][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}}
	fmt.Println(maxPoints1(points))
}


func maxPoints1(points [][]int) int {
	n:=len(points)
	if n <= 2{
		return n
	}
	ans:=0
	for i:=0;i<n - 1;i++{
		cache:=make(map[int]int)
		for j:=i+1;j<n;j++{
			if ans >= n - j || ans > n/2 {
				break
			}

			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			// 约分
			gcd:= gcd1(dx,dy)
			dx = dx/gcd
			dy = dy/gcd

			if dx == 0{
				dy = 1
			}
			if dy == 0{
				dx = 1
			}

			if (dx < 0) {
				dx,dy = 0 - dx,0 - dy
			}
			//key := fmt.Sprintf("%d,%d", dx,dy)
			//if _,succ:= cache[key];succ{
			//	cache[key]++
			//}else{
			//	cache[key] = 1
			//}
			cache[dy+dx*20001]++
		}
		for _,val:=range cache{
			ans = max1(ans,val)
		}
	}
	return ans + 1

}

func gcd1(i,j int) int{
	if i==0 || j==0{
		return 1
	}
	if i < j{
		i,j = j,i
	}

	for i % j != 0{
		i,j = j, i % j
	}
	if j<0 {
		return -j
	}
	return j
}


func max1(i,j int)int{
	if i>j{
		return i
	}
	return j
}
