package main

import "fmt"

//读写nil管道会阻塞,写close的管道会panic
//读关闭读管道，有数据会返回true
//slice中括号里可以有两个冒号？
//能否访问slice超过len 的元素？ 不能
//slice 题目2 是啥意思 append会导致扩容，扩容会重新分配内存
func main(){
	SliceExtend()
}

func SliceExtend() {
	s := make([]int, 0, 10)
	s1 := append(s, 1, 2, 3)
	s2 := append(s1, 4)

	fmt.Println(&s1[0] == &s2[0])
}

