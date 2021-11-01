package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0{
		return 0
	}
	//维护单词与索引的映射关系
	dict:=make(map[string]int)
	//图
	graph:=[][]int{}
	//增加索引
	addOrGet:=func(s string) int{
		if id,has:=dict[s];has{
			return id
		}
		id := len(graph)
		dict[s] = id
		graph = append(graph,[]int{})
		return id
	}
	//构造图
	buildGraph:=func(s string){
		id:= addOrGet(s)
		array:=[]byte(s)
		for index,val:=range array{
			array[index] = '*'
			id2:= addOrGet(string(array))
			graph[id] = append(graph[id],id2)
			graph[id2] = append(graph[id2],id)
			array[index] = val
		}
	}
	for _,val:=range wordList{
		buildGraph(val)
	}

	if _,has:=dict[endWord];!has{
		return 0
	}
	endId:=addOrGet(endWord)

	buildGraph(beginWord)
	beginId:=addOrGet(beginWord)

	// 广度优先遍历
	queue:=[]int{beginId}
	lenMap:=make(map[int]int)
	for _,val:=range dict{
		lenMap[val] = -1
	}
	lenMap[beginId] = 0
	for len(queue) > 0{
		target:= queue[0]
		queue= queue[1:]
		currLen,_:= lenMap[target]
		for _,val:=range graph[target]{
			if lenMap[val] < 0{
				if val == endId{
					return (currLen + 1)/2 + 1
				}else{
					lenMap[val] = currLen + 1
					queue = append(queue,val)
				}
			}
		}
	}

	return 0
}

func main(){
	beginWord:="hot"
	endWord:="dog"
	wordList:=[]string{"hot","dog"}
	fmt.Println(ladderLength(beginWord,endWord,wordList))
}
