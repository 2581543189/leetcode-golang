package main

import "fmt"

func main(){
	fmt.Println(countHighestScoreNodes1([]int{-1,3,3,5,7,6,0,0}))
}
// cache 中保存的是每一个节点的子节点数量
var cache []int = make([]int,0)

func countHighestScoreNodes1(parents []int) int {
	n:=len(parents)
	if n <= 2{
		return n
	}
	// 构造图
	cache = make([]int,n)
	graph:=make([][]int,n)
	for i:=range parents{
		if parents[i] != -1{
			graph[parents[i]] = append(graph[parents[i]],i)
		}
	}

	// 通过深度优先遍历 统计每个节点的子节点数量
	dfs(graph,0)

	// 统计每个节点分数
	arr:=make([]int,n)
	for i:=range graph{
		// 叶子节点 快速返回
		if len(graph[i]) == 0{
			arr[i] = n - 1
			continue
		}
		// 乘子节点数量
		arr[i]=1
		sum:=0
		for j:= range graph[i]{
			sum +=cache[graph[i][j]] + 1
			arr[i] *= cache[graph[i][j]] + 1
		}
		// 乘父节点数量
		if n - sum - 1 > 0{
			arr[i] *= n - sum - 1
		}
	}

	// 计算结果
	max:=0
	ans:=0
	for i:=range arr{
		if arr[i] > max{
			max = arr[i]
			ans = 1
		}else if  arr[i] == max{
			ans++
		}
	}
	return ans
}
//深度优先遍历
func dfs(graph [][]int,x int)int{
	if len(graph[x]) == 0{
		return 0
	}
	if cache[x] > 0{
		return cache[x]
	}
	ans:=0
	for i:=range graph[x]{
		node:= graph[x][i]
		if cache[node] > 0{
			ans += cache[node]
		}else{
			ans += dfs(graph,node) + 1
		}
	}
	cache[x] = ans
	return ans
}
