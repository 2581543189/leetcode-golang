package main

import "fmt"

func countPrimes1(n int) int {
	cache :=make([]bool,n+1)
	prime :=[]int{}

	for i:=0;i<n;i++{
		cache[i] = true
	}

	for i:=2;i<n;i++{
		if cache[i] {
			prime = append(prime,i)
			for _,val:=range prime{
				if i * val < n{
					cache[i * val] = false
				}else{
					break
				}
			}
		}else{
			for _,val:=range prime{
				if i * val < n {
					cache[i * val] = false
					if  i % val == 0{
						break
					}
				}else{
					break
				}
			}
		}
	}
	return len(prime)
}

func main(){
	fmt.Println(countPrimes1(2000))
}
