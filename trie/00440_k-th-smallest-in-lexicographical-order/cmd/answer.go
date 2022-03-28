package main

import "fmt"

func main(){
	fmt.Println(findKthNumber(122,16))
}

func count(pre ,end int)int{
	ans:=0
	s,e:=pre,pre
	for s <= end{
		ans += min(e,end) - s + 1
		s *=10
		e = e * 10 + 9
	}
	return ans
}

func findKthNumber(n, k int) int {
	cur:=1
	k-- // 根结点
	for k > 0{
		c:= count(cur,n)
		if c > k{
			// 前缀不变，继续寻找
			cur *= 10
			k--
		}else{
			// 前缀+1
			cur++
			k-=c
		}
	}
	return cur
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}