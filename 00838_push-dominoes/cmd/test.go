package main

import "fmt"

func main(){
	// 并未成功解答
	fmt.Println(pushDominoes1("R..L..R..LR.R.R....."))
}


func pushDominoes1(dominoes string) string {
	n:=len(dominoes)
	if n == 1{
		return dominoes
	}
	ans:=make([]byte,n)


	firstIndex:=-1
	first:=byte(0)
	// 初始化
	for i:=0;i<n;i++{
		if dominoes[i] != byte('.'){
			ans[i] = dominoes[i]
			if firstIndex < 0{
				firstIndex = i
				first = dominoes[i]
			}
		}
	}
	// 处理左边界
	if first == byte('R'){
		for i:=0;i<firstIndex;i++{
			ans[i] = byte('.')
		}
	}else if first ==  byte('L'){
		for i:=0;i<firstIndex;i++{
			ans[i] = byte('L')
		}
	}
	lastIndex:=-1
	last:=byte(0)
	for i:=n-1;i>=0;i--{
		if dominoes[i] != byte('.'){
			if lastIndex < 0{
				lastIndex = i
				last = dominoes[i]
			}
		}
	}
	// 处理右边界
	if last == byte('L'){
		for i:=n-1;i>lastIndex;i--{
			ans[i] = byte('.')
		}
	}else if last == byte('R'){
		for i:=n-1;i>lastIndex;i--{
			ans[i] = byte('R')
		}
	}


	if firstIndex < 0{
		return dominoes
	}
	d:=0
	// 模拟
	for d < n{
		d = 0
		tmp:=make([]byte,n)
		for i:=0;i<n;i++{
			tmp[i] = calOne(i,n,ans)
		}
		for i:=0;i<n;i++{
			if tmp[i] != 0{
				ans[i] = tmp[i]
				d++
			}
		}
	}
	return string(ans)
}

func calOne(i,n int,ans []byte )byte{
	if ans[i] != 0{
		return ans[i]
	}
	if i == 0 {
		if ans[1] == byte('L') || ans[1] == byte('.'){
			return ans[1]

		}
		if ans[1] == byte('R'){
			return byte('.')
		}
		return 0
	}else if i == n - 1{
		if ans[n - 2] == byte('R') || ans[n - 2] == byte('.'){
			return ans[n - 2]
		}
		if ans[n - 2] == byte('L'){
			return byte('.')
		}
		return 0
	}else{
		if ans[ i - 1 ] == 0 && ans[ i + 1 ] == 0{
			return 0
		}
		if ans[ i - 1 ] == byte('L') && ans[ i + 1 ] == 0{
			return 0
		}
		if ans[ i + 1 ] == byte('R') && ans[ i - 1 ] == 0{
			return 0
		}
		if(ans[ i - 1 ] == byte('L') && ans[ i + 1 ] == byte('R'))  || (ans[ i - 1 ] == byte('R') && ans[ i + 1 ] == byte('L')) {
			return byte('.')
		}

		if ans[ i - 1 ] ==byte('R'){
			return ans[ i - 1 ]
		}else if ans[ i + 1 ] ==byte('L'){
			return ans[ i + 1 ]
		}else{
			return byte('.')
		}
	}
}



