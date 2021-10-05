package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 栈
type Stack struct {
	data [] byte
}


func (stack *Stack) push(in byte){
	stack.data = append(stack.data,in)
}

func (stack *Stack) pop() byte{
	if len(stack.data) == 0{
		return 0
	}
	ret := stack.data[len(stack.data) -1]
	stack.data = stack.data[:len(stack.data) -1]
	return ret;
}

func (stack *Stack) print(){
	fmt.Printf("%s\n",string(stack.data))
}

func (stack *Stack) pushStr(in string){
	array := []byte (in)
	for i:=len(array) -1;i>=0;i--{
		stack.push(array[i])
	}
}

// tir-tree

type TirTree struct {
	Node map [byte] *TirTree
	Word bool
}

func (tirTree *TirTree) init(array []string){
	tirTree.Node = make(map [byte] *TirTree)
	for i:=len(array) -1;i>=0;i--{
		tirTree.addStr(array[i])
	}
}

func (tirTree *TirTree) addStr(str string){
	stack := new(Stack)
	stack.pushStr(str)
	tirTree.addStack(stack)
}

func (tirTree *TirTree) addStack(stack *Stack){
	val:=  stack.pop()
	if val == 0{
		tirTree.Word = true
		return
	}
	if(tirTree.Node[val] == nil){
		newNode := new(TirTree)
		newNode.Node = make(map [byte] *TirTree)
		tirTree.Node[val] = newNode
	}
	tirTree.Node[val].addStack(stack)
}

func (tirTree *TirTree) hasString(str string,isPrefix bool) bool{
	stack := new(Stack)
	stack.pushStr(str)
	return tirTree.hasStack(stack,isPrefix)

}

func (tirTree *TirTree) hasStack(stack *Stack,isPrefix bool) bool{
	for{
		val := stack.pop()
		if val == 0{
			if isPrefix {
				return true
			}else{
				return tirTree.Word;
			}
		}
		if tirTree.Node[val] != nil {
			return tirTree.Node[val].hasStack(stack,isPrefix)
		}else{
			return false
		}
	}
}



// 查找方法
func findWords1(board [][]byte, words []string) []string {

	//构造返回值
	result := [] string {}


	words = filterWord(board,words)
	if len(words) == 0{
		return result
	}

	//缓存
	cache := make(map[string]string)

	//构造tirtree
	tirTree := new(TirTree)
	tirTree.init(words)

	//step 1 尝试把所有可能的情况列出来
	for i:=0;i<len(board);i++{
		line := board[i]
		for j:=0;j<len(line);j++{
			//fmt.Printf("开始执行[%d,%d]:%s,所有可能的情况如下：\n",i,j,string(line[j]))
			stack := new(Stack)
			dfs(i,j,board,stack,tirTree,cache,len(words))
		}
	}
	for key,_ := range cache{
		result = append(result,key)
	}

	return result
}

//过滤明显不符合的word
func filterWord(board [][]byte, words []string)[]string{
	result := []string{}

	letterMap := make(map[byte]byte)

	//step 1 构造map
	for i:=0;i<len(board);i++{
		line := board[i]
		for j:=0;j<len(line);j++{
			key:= board[i][j]
			letterMap[key] = 1
		}
	}
	//step 2 过滤
	for _,word := range words{
		array:= []byte(word)
		validate := true
		for _,oneByte:=range array{
			if letterMap[oneByte] == 0 {
				validate = false
			}
		}
		if validate {
			result = append(result,word)
		}
	}
	return result;
}

//动态规划
func dfs(i int,j int ,board [][]byte,stack *Stack,tree *TirTree,cache map[string]string,wordCount int) bool {
	data := board[i][j]
	if string(data) == "#"{
		return false
	}

	if !tree.hasString(string(stack.data),true){
		return false
	}

	stack.push(data)
	board[i][j] = byte('#')

	if tree.hasString(string(stack.data),false){
		cache[string(stack.data)] = ""
		if len(cache) == wordCount{
			return true
		}
	}

	isLast:=true;

	//上
	if(i > 0){
		if dfs(i - 1,j,board,stack,tree,cache,wordCount){
			isLast = false;
		}
	}

	//下
	if(i < len(board) - 1){
		if dfs(i + 1,j,board,stack,tree,cache,wordCount){
			isLast = false;
		}
	}

	//左
	if(j > 0){
		if dfs(i,j-1,board,stack,tree,cache,wordCount){
			isLast = false;
		}
	}

	//右
	if(j < len(board[i]) - 1){
		if dfs(i,j+1,board,stack,tree,cache,wordCount){
			isLast = false;
		}
	}
	if(isLast){
		//stack.print()
	}
	stack.pop()
	board[i][j] = data;
	return true;
}

func main(){

	//var board [][]byte = [][]byte{[]byte("oaan"),[]byte("etae"),[]byte("ihkr"),[]byte("iflv")}
	//var words [] string = [] string { "oath","pea","eat","rain" }

	var board [][]byte = [][]byte{
		[]byte("bababababa"),
		[]byte("ababababab"),
		[]byte("bababababa"),
		[]byte("ababababab"),
		[]byte("bababababa"),
		[]byte("ababababab"),
		[]byte("bababababa"),
		[]byte("ababababab"),
		[]byte("bababababa"),
		[]byte("ababababab"),
	}
	var words [] string = [] string { "ababababaa","ababababab","ababababac","ababababad","ababababae","ababababaf",
		"ababababag","ababababah","ababababai","ababababaj","ababababak","ababababal","ababababam","ababababan",
		"ababababao","ababababap","ababababaq","ababababar","ababababas","ababababat","ababababau","ababababav",
		"ababababaw","ababababax","ababababay","ababababaz","ababababba","ababababbb","ababababbc","ababababbd",
		"ababababbe","ababababbf","ababababbg","ababababbh","ababababbi","ababababbj","ababababbk","ababababbl",
		"ababababbm","ababababbn","ababababbo","ababababbp","ababababbq","ababababbr","ababababbs","ababababbt",
		"ababababbu","ababababbv","ababababbw","ababababbx","ababababby","ababababbz","ababababca","ababababcb",
		"ababababcc","ababababcd","ababababce","ababababcf","ababababcg","ababababch","ababababci","ababababcj",
		"ababababck","ababababcl","ababababcm","ababababcn","ababababco","ababababcp","ababababcq","ababababcr",
		"ababababcs","ababababct","ababababcu","ababababcv","ababababcw","ababababcx","ababababcy","ababababcz",
		"ababababda","ababababdb","ababababdc","ababababdd","ababababde","ababababdf","ababababdg","ababababdh",
		"ababababdi","ababababdj","ababababdk","ababababdl","ababababdm","ababababdn","ababababdo","ababababdp",
		"ababababdq","ababababdr","ababababds","ababababdt","ababababdu","ababababdv","ababababdw","ababababdx",
		"ababababdy","ababababdz","ababababea","ababababeb","ababababec","ababababed","ababababee","ababababef",
		"ababababeg","ababababeh","ababababei","ababababej","ababababek","ababababel","ababababem","ababababen",
		"ababababeo","ababababep","ababababeq","ababababer","ababababes","ababababet","ababababeu","ababababev",
		"ababababew","ababababex","ababababey","ababababez","ababababfa","ababababfb","ababababfc","ababababfd",
		"ababababfe","ababababff","ababababfg","ababababfh","ababababfi","ababababfj","ababababfk","ababababfl",
		"ababababfm","ababababfn","ababababfo","ababababfp","ababababfq","ababababfr","ababababfs","ababababft",
		"ababababfu","ababababfv","ababababfw","ababababfx","ababababfy","ababababfz","ababababga","ababababgb",
		"ababababgc","ababababgd","ababababge","ababababgf","ababababgg","ababababgh","ababababgi","ababababgj",
		"ababababgk","ababababgl","ababababgm","ababababgn","ababababgo","ababababgp","ababababgq","ababababgr",
		"ababababgs","ababababgt","ababababgu","ababababgv","ababababgw","ababababgx","ababababgy","ababababgz",
		"ababababha","ababababhb","ababababhc","ababababhd","ababababhe","ababababhf","ababababhg","ababababhh",
		"ababababhi","ababababhj","ababababhk","ababababhl","ababababhm","ababababhn","ababababho","ababababhp",
		"ababababhq","ababababhr","ababababhs","ababababht","ababababhu","ababababhv","ababababhw","ababababhx",
		"ababababhy","ababababhz","ababababia","ababababib","ababababic","ababababid","ababababie","ababababif",
		"ababababig","ababababih","ababababii","ababababij","ababababik","ababababil","ababababim","ababababin",
		"ababababio","ababababip","ababababiq","ababababir","ababababis","ababababit","ababababiu","ababababiv",
		"ababababiw","ababababix","ababababiy","ababababiz","ababababja","ababababjb","ababababjc","ababababjd",
		"ababababje","ababababjf","ababababjg","ababababjh","ababababji","ababababjj","ababababjk","ababababjl",
		"ababababjm","ababababjn","ababababjo","ababababjp","ababababjq","ababababjr","ababababjs","ababababjt",
		"ababababju","ababababjv","ababababjw","ababababjx","ababababjy","ababababjz","ababababka","ababababkb",
		"ababababkc","ababababkd","ababababke","ababababkf","ababababkg","ababababkh","ababababki","ababababkj",
		"ababababkk","ababababkl","ababababkm","ababababkn","ababababko","ababababkp","ababababkq","ababababkr",
		"ababababks","ababababkt","ababababku","ababababkv","ababababkw","ababababkx","ababababky","ababababkz",
		"ababababla","abababablb","abababablc","ababababld","abababable","abababablf","abababablg","abababablh",
		"ababababli","abababablj","abababablk","ababababll","abababablm","ababababln","abababablo","abababablp",
		"abababablq","abababablr","ababababls","abababablt","abababablu","abababablv","abababablw","abababablx",
		"abababably","abababablz","ababababma","ababababmb","ababababmc","ababababmd","ababababme","ababababmf",
		"ababababmg","ababababmh","ababababmi","ababababmj","ababababmk","ababababml","ababababmm","ababababmn",
		"ababababmo","ababababmp","ababababmq","ababababmr","ababababms","ababababmt","ababababmu","ababababmv",
		"ababababmw","ababababmx","ababababmy","ababababmz","ababababna","ababababnb","ababababnc","ababababnd",
		"ababababne","ababababnf","ababababng","ababababnh","ababababni","ababababnj","ababababnk","ababababnl",
		"ababababnm","ababababnn","ababababno","ababababnp","ababababnq","ababababnr","ababababns","ababababnt",
		"ababababnu","ababababnv","ababababnw","ababababnx","ababababny","ababababnz","ababababoa","ababababob",
		"ababababoc","ababababod","ababababoe","ababababof","ababababog","ababababoh","ababababoi","ababababoj",
		"ababababok","ababababol","ababababom","ababababon","ababababoo","ababababop","ababababoq","ababababor",
		"ababababos","ababababot","ababababou","ababababov","ababababow","ababababox","ababababoy","ababababoz",
		"ababababpa","ababababpb","ababababpc","ababababpd","ababababpe","ababababpf","ababababpg","ababababph",
		"ababababpi","ababababpj","ababababpk","ababababpl","ababababpm","ababababpn","ababababpo","ababababpp",
		"ababababpq","ababababpr","ababababps","ababababpt","ababababpu","ababababpv","ababababpw","ababababpx",
		"ababababpy","ababababpz","ababababqa","ababababqb","ababababqc","ababababqd","ababababqe","ababababqf",
		"ababababqg","ababababqh","ababababqi","ababababqj","ababababqk","ababababql","ababababqm","ababababqn",
		"ababababqo","ababababqp","ababababqq","ababababqr","ababababqs","ababababqt","ababababqu","ababababqv",
		"ababababqw","ababababqx","ababababqy","ababababqz","ababababra","ababababrb","ababababrc","ababababrd",
		"ababababre","ababababrf","ababababrg","ababababrh","ababababri","ababababrj","ababababrk","ababababrl",
		"ababababrm","ababababrn","ababababro","ababababrp","ababababrq","ababababrr","ababababrs","ababababrt",
		"ababababru","ababababrv","ababababrw","ababababrx","ababababry","ababababrz","ababababsa","ababababsb",
		"ababababsc","ababababsd","ababababse","ababababsf","ababababsg","ababababsh","ababababsi","ababababsj",
		"ababababsk","ababababsl","ababababsm","ababababsn","ababababso","ababababsp","ababababsq","ababababsr",
		"ababababss","ababababst","ababababsu","ababababsv","ababababsw","ababababsx","ababababsy","ababababsz",
		"ababababta","ababababtb","ababababtc","ababababtd","ababababte","ababababtf","ababababtg","ababababth",
		"ababababti","ababababtj","ababababtk","ababababtl","ababababtm","ababababtn","ababababto","ababababtp",
		"ababababtq","ababababtr","ababababts","ababababtt","ababababtu","ababababtv","ababababtw","ababababtx",
		"ababababty","ababababtz","ababababua","ababababub","ababababuc","ababababud","ababababue","ababababuf",
		"ababababug","ababababuh","ababababui","ababababuj","ababababuk","ababababul","ababababum","ababababun",
		"ababababuo","ababababup","ababababuq","ababababur","ababababus","ababababut","ababababuu","ababababuv",
		"ababababuw","ababababux","ababababuy","ababababuz","ababababva","ababababvb","ababababvc","ababababvd",
		"ababababve","ababababvf","ababababvg","ababababvh","ababababvi","ababababvj","ababababvk","ababababvl",
		"ababababvm","ababababvn","ababababvo","ababababvp","ababababvq","ababababvr","ababababvs","ababababvt",
		"ababababvu","ababababvv","ababababvw","ababababvx","ababababvy","ababababvz","ababababwa","ababababwb",
		"ababababwc","ababababwd","ababababwe","ababababwf","ababababwg","ababababwh","ababababwi","ababababwj",
		"ababababwk","ababababwl","ababababwm","ababababwn","ababababwo","ababababwp","ababababwq","ababababwr",
		"ababababws","ababababwt","ababababwu","ababababwv","ababababww","ababababwx","ababababwy","ababababwz",
		"ababababxa","ababababxb","ababababxc","ababababxd","ababababxe","ababababxf","ababababxg","ababababxh",
		"ababababxi","ababababxj","ababababxk","ababababxl","ababababxm","ababababxn","ababababxo","ababababxp",
		"ababababxq","ababababxr","ababababxs","ababababxt","ababababxu","ababababxv","ababababxw","ababababxx",
		"ababababxy","ababababxz","ababababya","ababababyb","ababababyc","ababababyd","ababababye","ababababyf",
		"ababababyg","ababababyh","ababababyi","ababababyj","ababababyk","ababababyl","ababababym","ababababyn",
		"ababababyo","ababababyp","ababababyq","ababababyr","ababababys","ababababyt","ababababyu","ababababyv",
		"ababababyw","ababababyx","ababababyy","ababababyz","ababababza","ababababzb","ababababzc","ababababzd",
		"ababababze","ababababzf","ababababzg","ababababzh","ababababzi","ababababzj","ababababzk","ababababzl",
		"ababababzm","ababababzn","ababababzo","ababababzp","ababababzq","ababababzr","ababababzs","ababababzt",
		"ababababzu","ababababzv","ababababzw","ababababzx","ababababzy","ababababzz" }


	finded := findWords1(board,words)
	date,error := json.Marshal(finded)
	if(error != nil){
		log.Fatalf("JSON marshaling failed: %s" , error)
	}
	fmt.Printf("%s\n",date)


}
