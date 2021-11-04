package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(fractionToDecimal1(4,333))
}

func fractionToDecimal1(numerator int, denominator int) string {
	array:=[]byte{}
	if numerator % denominator == 0{
		return strconv.Itoa(numerator / denominator)
	}

	if (numerator < 0) != (denominator < 0){
		array = append(array,'-')
	}
	numerator = abs1(numerator)
	denominator = abs1(denominator)
	str:=strconv.Itoa(numerator / denominator)
	array = append(array,([]byte(str))... )
	array = append(array,'.')
	cache:=make(map[int]int)
	mod := numerator % denominator
	for mod !=0 && cache[mod] == 0{
		cache[mod] = len(array)
		numerator = 10 * mod
		mod = numerator % denominator
		array = append(array,'0' + byte(numerator / denominator))
	}

	if mod != 0 {
		array = append(array[:cache[mod]],append([]byte{'('},array[cache[mod]:]...)...)
		array = append(array,')')
	}

	return string(array)
}

func abs1(i int)int{
	if i < 0{
		return -i
	}
	return i
}