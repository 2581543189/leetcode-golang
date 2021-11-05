package main

import (
	"fmt"
	"math"
)

type Cache struct {
	Data map[int]int
}

func (cache *Cache) get(i,j int)int{
	if i == 1{
		return j
	}else if j == 0{
		return 0
	}
	key := i + 20000 * j
	if val,ok := cache.Data[key];ok{
		return val
	}else{
		return -1
	}
}

func (cache *Cache) put(i,j,val int){
	key := i + 20000 * j
	cache.Data[key] = val
}


func superEggDrop(k int, n int) int {
	// 加入我们先不考虑二分
	cache := &Cache{make(map[int]int)}

	var dfs func(x,y int) int
	dfs = func(x,y int) int{
		ans := cache.get(x,y)
		if ans >= 0{
			return ans
		}
		if x > y{
			ans = dfs(x-1,y)
			cache.put(x,y,ans)
			return ans
		}
		ans=math.MaxInt32
		//for i:=1; i<=y; i++{
		//	ans = min(ans,max(dfs(x - 1,i - 1),dfs(x,y - i)) + 1)
		//}
		//这里改为二分
		l,r := 1,y
		for l < r{
			mid := l + (r - l) >> 1
			broken:= dfs(x - 1,mid - 1) + 1
			notBroken := dfs(x,y - mid) + 1
			if broken > notBroken{
				r = mid
			}else{
				l = mid + 1
			}
			ans = min(ans,max(broken,notBroken))
		}

		cache.put(x,y,ans)
		return ans
	}
	return dfs(k,n)
}

func max(i,j int)int{
	if i > j {
		return i
	}
	return j
}

func min(i,j int)int{
	if i < j {
		return i
	}
	return j
}

func main()  {
	fmt.Println(superEggDrop(2,6))
}

