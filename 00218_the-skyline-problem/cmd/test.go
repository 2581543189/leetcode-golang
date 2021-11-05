package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type  hp1 struct {
	buildings [][]int
}

func (this *hp1) Push(val interface{}){
	this.buildings = append(this.buildings,val.([]int))
}
func (this *hp1) Pop()interface{}{
	val := this.buildings[len(this.buildings) - 1]
	this.buildings = this.buildings[0:len(this.buildings) - 1]
	return val
}
func (this hp1) Swap(i,j int){
	this.buildings[i],this.buildings[j] = this.buildings[j],this.buildings[i]
}

func (this hp1) Less(i,j int)bool{
	return this.buildings[i][2] > this.buildings[j][2]
}

func (this hp1) Len()int{
	return len(this.buildings)
}


func getSkyline1(buildings [][]int) [][]int {
	n := len(buildings)
	if n == 0 {
		return [][]int{}
	}
	// 对墙的坐标进行排序
	walls := make([]int ,0 , 2 * n)
	for _,building := range buildings{
		walls = append(walls,building[0],building[1])
	}
	sort.Ints(walls)

	//使用队列计算结果
	ans:= [][]int{}
	cache:= &hp1{}
	index:= 0
	highest:=-1
	currentHighest:=-1

	for _,wall :=range walls{
		// 处理新增
		for index < n && buildings[index][0] == wall{
			heap.Push(cache,[]int{buildings[index][0],buildings[index][1],buildings[index][2]})
			index++
		}
		//删除无效
		for cache.Len() > 0 && cache.buildings[0][1] <= wall{
			heap.Pop(cache)
		}
		if cache.Len() == 0{
			currentHighest = 0
		}else{
			currentHighest = cache.buildings[0][2]
		}
		//写入最高值
		if currentHighest != highest{
			highest = currentHighest
			ans = append(ans,[]int{wall,highest})
		}
	}

	return ans

}

func main(){
	buildings :=[][]int {{2,9,10},{3,7,15},{5,12,12},{15,20,10},{19,24,8}}
	skyLine := getSkyline1(buildings)
	fmt.Println(skyLine)

	//验证循环中删除数据
	for len(buildings) > 0 && buildings[0][0] < 19{
		buildings = buildings[1:len(buildings)]
	}
	fmt.Println(buildings)
}
