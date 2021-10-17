package main

import "fmt"

func calculate1(s string) int {

	stack := []int{}
	num:=0
	lastOp:=' '
	s= s + "+"
	for _,ch := range s {
		if isNum(ch) {
			num = int(ch - '0') + 10 * num
		}else if lastOp == ' ' && isOp(ch){
			stack = append([]int{},num)
			num = 0
			lastOp = ch
		}else if lastOp != ' ' && isOp(ch) {
			operate(&stack,num,lastOp)
			num = 0
			lastOp = ch
		}
	}
	num = 0
	for i:=0 ; i< len(stack); i++{
		num += stack[i]
	}

	return num
}

func isOp(s rune)bool{
	return s == '+' ||s == '-' ||s == '*' ||s == '/'
}
func  isNum(s rune)bool{
	return s - '0' >= 0 &&  s - '0'<=9
}

func operate(p *[]int,num int,lastOp rune){
	n:=len(*p)
	if lastOp == '+'{
		*p = append(*p ,num)
	}else if lastOp == '-'{
		*p = append(*p ,-num)
	}else if lastOp == '*'{
		(*p)[n - 1] = (*p)[len(*p) - 1] * num
	}else if lastOp == '/'{
		(*p)[n - 1] = (*p)[len(*p) - 1] / num
	}
}

func main(){
	fmt.Println(calculate1("1"))
}
