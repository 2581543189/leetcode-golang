package main

import "fmt"

func main(){
	fmt.Println(winnerOfGame2("AAAABBBB"))
}

func winnerOfGame2(colors string) bool {
	if len(colors) == 1{
		return false
	}
	aCount,bCount:=0,0
	a,b:=0,0
	if colors[0] == 'B'{
		bCount++
	}else{
		aCount++
	}
	for i:=1;i<len(colors);i++{
		if colors[i] == 'A'{
			aCount++
			if colors[i] != colors[i-1]{
				if bCount > 2{
					b += bCount-2
				}
				bCount=0
			}
		}else{
			bCount++
			if colors[i] != colors[i-1]{
				if aCount > 2{
					a += aCount-2
				}
				aCount=0
			}
		}
	}
	if aCount > 2{
		a += aCount-2
	}
	if bCount > 2{
		b += bCount-2
	}

	if a == 0{
		return false
	}
	return a > b
}