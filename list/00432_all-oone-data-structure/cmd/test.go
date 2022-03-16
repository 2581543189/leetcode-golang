package main

import "fmt"

func main(){
	a:= Constructor1()
	a.Inc("hello")
	a.Inc("goodbye")
	a.Inc("hello")
	a.Inc("hello")
	fmt.Println(a.GetMaxKey())
	a.Inc("leet")
	a.Inc("code")
	a.Inc("leet")
	a.Dec("hello")
	a.Inc("leet")
	a.Inc("code")
	a.Inc("code")
	fmt.Println(a.GetMaxKey())

}


type Node struct{
	pre *Node
	next *Node
	count int
	val string
}

type AllOne1 struct {
	head *Node
	tail *Node
	count int
	dict map[string] *Node
}


func Constructor1() AllOne1 {
	head:= &Node{}
	tail:=&Node{}
	head.next = tail
	tail.pre = head
	return AllOne1{
		head,
		tail,
		0,
		make(map[string]*Node),
	}
}


func (this *AllOne1) Inc(key string)  {
	dict := this.dict
	if val,ok:=dict[key];ok{
		newVal:=val.count + 1
		tmp:=val
		for val.next.count <= newVal && val.next!=this.tail{
			val = val.next
		}
		newNode :=&Node{val,val.next,newVal,key}
		if tmp != val{
			tmp.pre.next = tmp.next
			tmp.next.pre = tmp.pre
			val.next.pre = newNode
			val.next = newNode
		}else{
			newNode.pre = val.pre
			val.next.pre = newNode
			val.pre.next = newNode
		}

		dict[key] = newNode
	}else{
		newNode :=&Node{this.head,this.head.next,1,key}
		this.head.next.pre = newNode
		this.head.next = newNode
		dict[key] = newNode
	}
	this.count++

}


func (this *AllOne1) Dec(key string)  {
	dict := this.dict
	oldNode:=dict[key]
	if oldNode.count == 1{
		oldNode.pre.next = oldNode.next
		oldNode.next.pre = oldNode.pre
		delete(dict,key)
	}else{
		newVal:=oldNode.count - 1
		tmp:=oldNode
		for oldNode.pre.count >= newVal && oldNode.pre!=this.head{
			oldNode = oldNode.pre
		}
		newNode :=&Node{oldNode.pre,oldNode,newVal,key}
		if tmp!=oldNode{
			tmp.next.pre = tmp.pre
			tmp.pre.next = tmp.next
			oldNode.pre.next = newNode
			oldNode.pre = newNode

		}else{
			newNode.next = oldNode.next
			tmp.pre.next = newNode
			tmp.next.pre = newNode
		}

		dict[key] = newNode
	}
	this.count--
}


func (this *AllOne1) GetMaxKey() string {
	if this.count > 0{
		return this.tail.pre.val
	}else{
		return ""
	}
}


func (this *AllOne1) GetMinKey() string {
	if this.count > 0{
		return this.head.next.val
	}else{
		return ""
	}
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */