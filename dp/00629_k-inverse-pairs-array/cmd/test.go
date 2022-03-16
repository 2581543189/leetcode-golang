package main

import "fmt"

func kInversePairs1(n int, k int) int {
	if k == 0{
		return 1
	}
	if n == 1 {
		return 0
	}
	if n == 2{
		if k == 1{
			return 1
		}else{
			return 0
		}
	}
	// 下面计算 n > 3 的场景

	if k > n * (n-1) / 2{
		// k 不可能大于 n * (n-1) / 2
		return 0
	}
	// 用于减少计算量的缓存
	cache:=make(map[int]int)
	cache[generateKey(3,1)] = 2
	cache[generateKey(3,2)] = 2

	var dfs func(i int, j int)int
	dfs = func(i int, j int)int{
		if j == 0{
			return 1
		}
		if j == i * (i - 1) / 2{
			return 1
		}
		if j > i * (i-1) / 2{
			// k 不可能大于 n * (n-1) / 2
			return 0
		}
		if val,ok:=cache[generateKey(i,j)];ok{
			return val
		}
		ans:= 0
		kmax := (i - 1)*(i - 2) / 2
		kmax = min1(kmax,j)

		kmin:= 0
		if j > kmax || j > i-1{
			kmin =  max1(j - kmax,j-i+1)
		}

		for k:=kmin; k <= kmax; k++ {
			ans += dfs(i-1,k) % (1e9+7)
		}
		ans = ans % (1e9+7)
		cache[generateKey(i,j)] = ans
		return ans
	}
	return dfs(n,k)
}

func generateKey(n int, k int)int{
	return 2000 * n + k
}

func min1(i,j int)int{
	if i < j{
		return i
	}
	return j
}

func max1(i,j int)int{
	if i > j{
		return i
	}
	return j
}

func main(){
	fmt.Println(kInversePairs1(1000,1000))
}
