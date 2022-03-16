package main

import (
	"fmt"
	"time"
)

func findMinStep(board string, hand string) int {
	if len(board) == 0 || len(hand) == 0{
		return -1
	}
	// 用于剪枝
	cache:=make(map[string]bool)
	//手中的球求排序，方便剪枝
	hand=sortHand(hand)
	q:=[][]string{{board,hand}}
	for len(q) > 0{
		cur:=q[0]
		q=q[1:]
		b1:=cur[0]
		h1:=cur[1]

		for ph:=0;ph<len(h1);ph++{
			h2:=h1[:ph]+h1[ph+1:]
			for pb:=0;pb<len(b1);pb++{
				b2:=b1[:pb] + string(h1[ph]) + b1[pb:]
				key:= fmt.Sprintf("%v_%v",b2,h2)
				if cache[key]{
					continue
				}
				cache[key] = true
				// 消除逻辑
				b2=del32(b2)
				if b2 == ""{
					return len(hand) - len(h2)
				}else{
					//消除结果加入队列
					q=append(q,[]string{b2,h2})
				}
			}
		}
	}
	return -1

}

func sortHand(origin string)string{
	counter:=make(map[byte]int)
	for _,s:= range origin{
		if _,ok:=counter[byte(s)];ok{
			counter[byte(s)]++
		}else{
			counter[byte(s)]=1
		}
	}
	ans:=[]byte{}
	for k,v:=range counter{
		for i:=0;i<v;i++{
			ans=append(ans,k)
		}
	}
	return string(ans)
}

// 递归碰撞删除所有长度3及以上的子串
func del32(b2 string) string {
	i,j:=0,0
	for i< len(b2)&& j<len(b2){
		if b2[i] == b2[j]{
			j++
		}else{
			if j - i > 2{
				b2 = b2[:i]+b2[j:]
				j=0
				i=0
			}else{
				i=j
			}
		}
	}
	if j - i > 2{
		b2 = b2[:i]+b2[j:]
	}
	return b2
}

func main(){
	board:="WWRRBBWW"
	hand:="WRBRW"
	fmt.Println(findMinStep(board,hand))
	k1 := time.Now()
	for i:=0;i<1000000;i++{
		del32("WWRRRBBWW")
	}
	k2 := time.Now()
	fmt.Println(k2.UnixNano() - k1.UnixNano())
}