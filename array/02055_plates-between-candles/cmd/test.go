package main

import "fmt"

func main(){
	// 未解答出来
	s:="**|**|***|"
	queries := [][]int{{2,5},{5,9}}
	fmt.Println(platesBetweenCandles1(s,queries))
}


type PorC int

const(
	P PorC = iota
	C
)

type One struct{
	index int
	porc PorC
}

type QueryWraper struct{
	start int
	end int
	candleCount int
	planteCount int
	coveredPlanteCount int
}

func (this *QueryWraper) deal(index int, porc PorC){
	if  index < this.start || index > this.end{
		return
	}
	if porc == C{
		this.candleCount++
		this.coveredPlanteCount +=  this.planteCount
		this.planteCount = 0
	}
	if porc == P{
		if this.candleCount == 0{
			return
		}
		this.planteCount++
	}
}

func platesBetweenCandles1(s string, queries [][]int) []int {

	// 先将字符串转换为 One
	ns:=make([]*One,0)
	for i:=range s{
		if i == 0{
			if s[i] == '|'{
				ns = append(ns,&One{i,C})
			}else{
				ns = append(ns,&One{i,P})
			}
		}else{
			if s[i] == '|' && ns[len(ns) - 1].porc == C{
				continue
			}
			if s[i] == '*' && ns[len(ns) - 1].porc == P{
				continue
			}
			if s[i] == '|'{
				ns = append(ns,&One{i,C})
			}else{
				ns = append(ns,&One{i,P})
			}
		}
	}

	qws:=make([]*QueryWraper,0)
	for i:=range queries{
		qws = append(qws,&QueryWraper{start:queries[i][0],end:queries[i][1]})
	}
	for i:=range ns{
		if ns[i].porc == P && i < len(ns) - 1{
			for x:= ns[i].index;x< ns[i+1].index;x++{
				for j:=range qws{
					qws[j].deal(x,P)
				}
			}
		}else{
			for j:=range qws{
				qws[j].deal(ns[i].index,C)
			}
		}

	}
	ans:=make([]int,0)
	for j:=range qws{
		ans = append(ans,qws[j].coveredPlanteCount)
	}
	return ans

}