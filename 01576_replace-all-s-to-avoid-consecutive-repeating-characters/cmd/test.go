package main

import "fmt"

func main()  {
	fmt.Println(modifyString1("??yw?ipkj?"))
	var a byte
	var b rune
	s := "111"
	for i,ch := range s{
		a  = s[i]
		b = ch
		if a == 'a'{

		}
		if b == 'a'{

		}
	}
	fmt.Println(a)
	fmt.Println(b)
}

func modifyString1(s string) string {
	n:=len(s)
	array:=[]rune(s)
	for i,_ := range s {
		if array[ i ] != rune('?'){
			continue
		}
		a,b,c:=0,0,0
		// 统计 前后的abc
		if i > 0{
			if array[ i - 1] == rune('a'){
				a++
			}
			if array[ i - 1] == rune('b'){
				b++
			}
			if array[ i - 1] == rune('c'){
				c++
			}
		}
		if i < n - 1 && array[ i + 1] != rune('?'){
			if array[ i + 1] == rune('a'){
				a++
			}
			if array[ i + 1] == rune('b'){
				b++
			}
			if array[ i + 1] == rune('c'){
				c++
			}
		}
		if a == 0{
			array[i] = rune('a')
		}else if b == 0{
			array[i] = rune('b')
		}else{
			array[i] = rune('c')
		}

	}

	return string(array)
}