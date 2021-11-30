package main

import (
	"fmt"
	"math/rand"
)

//读写nil管道会阻塞,写close的管道会panic
//读关闭读管道，有数据会返回true
//slice中括号里可以有两个冒号？
//能否访问slice超过len 的元素？ 不能
//slice 题目2 是啥意思 append会导致扩容，扩容会重新分配内存
//nil管道在case中会被忽略 没有default会永久阻塞
// fmt.Errorf 生成的 error 调用 Error 返回格式
// 子goroutine panic 整个程序也会崩溃


func main() {
	for i:=0;i<100;i++{
		fmt.Println(rand.Intn(5))
	}
}