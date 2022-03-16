package main

import "fmt"

type MapSum1 struct {
	data []*MapSum1
	sum int
	val int
	contain bool

}


func Constructor1() MapSum1 {
	return MapSum1{make([]*MapSum1,26),0,0,false}
}

func (this *MapSum1) Get(key string) (bool ,int) {
	cur:=this
	for i:=0;i<len(key) - 1;i++{
		index := key[i] - 'a'
		if cur.data[index] != nil{
			cur = cur.data[index]
			if cur == nil{
				return false,0
			}
		}else{
			return false,0
		}
	}
	index := key[len(key) - 1] - 'a'
	has:= cur.data[index] != nil && cur.data[index].contain
	if !has{
		return false,0
	}else{
		return true,cur.data[index].val
	}
}


func (this *MapSum1) Insert(key string, val int)  {
	has,oldVal:= this.Get(key)
	if has{
		val = val - oldVal
	}
	cur:=this
	for i:=0;i<len(key) - 1;i++{
		index := key[i] - 'a'
		if (cur.data[index] == nil){
			cur.data[index] = &MapSum1{make([]*MapSum1,26),0,0,false}
		}
		cur.data[index].sum += val
		cur = cur.data[index]
	}
	index := key[len(key) - 1] - 'a'
	if (cur.data[index] == nil){
		cur.data[index] = &MapSum1{make([]*MapSum1,26),val,val,true}
	}else{
		cur.data[index].sum += val
		cur.data[index].val += val
		cur.data[index].contain = true
	}

}


func (this *MapSum1) Sum(prefix string) int {
	cur:=this
	for i:=0;i<len(prefix) - 1;i++{
		index := prefix[i] - 'a'
		cur = cur.data[index]
		if cur == nil{
			return 0
		}
	}
	index := prefix[len(prefix) - 1] - 'a'
	if cur.data[index] ==nil{
		return 0
	}
	return cur.data[index].sum
}

func main(){
	sum:=Constructor1()
	sum.Insert("appple",1)
	sum.Insert("ap",2)
	sum.Insert("appple",2)
	fmt.Println(sum.Sum("ap"))

}
