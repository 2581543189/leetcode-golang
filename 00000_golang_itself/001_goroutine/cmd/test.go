package main

import (
	"fmt"
)

func main (){
	letter,number,finish := make(chan bool),make(chan bool),make(chan bool)

	go func() {
		i:=0
		for{
			_,ok:= <-letter
			if ok{
				if i < 26{
					fmt.Print(string('A'+i))
					i++
					fmt.Print(string('A'+i))
					i++
					number<-true
				}else{
					close(number)
				}
			}else{
				close(finish)
				return
			}
		}
	}()

	go func() {
		i:=1
		for{
			_,ok:= <-number
			if ok{
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter<-true
			}else{
				close(letter)
				return
			}
		}
	}()
	number<-true
	<- finish
}