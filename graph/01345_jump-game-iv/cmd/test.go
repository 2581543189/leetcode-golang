package main

import "fmt"

type element struct{
	val int
	position int
}

func main(){
	fmt.Println(minJumps1([]int{100,-23,-23,404,100,23,23,23,3,404}))
}
func minJumps1(arr []int) int {
	// 只有一个节点，直接返回0
	if len(arr) <= 1{
		return 0
	}
	n:=len(arr)
	// 统计每个值都在哪些位置上
	m:=make(map[int][]int)
	for i,_:=range arr{
		m[arr[i]] = append(m[arr[i]],i)
	}
	queue:=make([]*element,0)
	queue = append(queue,&element{arr[0],0})
	visited:=make(map[int]int)
	visited[0] = 1
	ans:=0
	for len(queue) > 0{
		cq:=make([]*element,len(queue))
		copy(cq,queue)
		queue = make([]*element,0)
		for _,e:=range cq{
			next := make([]int,0)
			hasNext:=false
			hasPre:=false
			for _,nb:=range m[e.val]{
				if visited[nb] > 0{
					continue
				}
				visited[nb] = 1
				if nb == e.position + 1{
					hasNext = true
				}
				if nb == e.position - 1{
					hasPre = true
				}
				next = append(next,nb)
			}
			delete(m, e.val)
			if !hasNext && e.position + 1 < n  && visited[e.position + 1] == 0{
				visited[e.position + 1] = 1
				next = append(next,e.position + 1)
			}
			if !hasPre && e.position - 1 > 0&& visited[e.position - 1] == 0{
				visited[e.position - 1] = 1
				next = append(next,e.position - 1)
			}
			for _,val:=range next{
				if val == n - 1{
					ans++
					return ans
				}
				queue = append(queue,&element{arr[val],val})
			}
		}
		ans++
	}
	return ans - 1

}