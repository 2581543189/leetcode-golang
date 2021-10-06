package main

import (
	"encoding/json"
	"fmt"
	"log"
)


// trie-tree
type TrieTree struct {
	childs  [26] *TrieTree
	words string
}


func (tirTree *TrieTree) addStr(str string){
	node := tirTree
	for _,ch := range str{
		index := ch - 'a'
		if node.childs[index] == nil {
			node.childs[index] = &TrieTree{}
		}
		node = node.childs[index];
	}
	node.words = str
}

var fourDirections =  []struct{x,y int}{{0,-1},{0,1},{-1,0},{1,0}}

// 查找方法
func findWords1(board [][]byte, words []string) []string {

	//构造返回值
	result := [] string {}

	//过滤
	words = filterWord(board,words)
	if(len(words) == 0){
		return result
	}

	trieTree := &TrieTree{}

	for _,word := range  words {
		trieTree.addStr(word)
	}

	//缓存
	cache := map[string]bool{}

	//定义dfs 方法
	var dfs func(node *TrieTree, x, y int)
	dfs = func(tria *TrieTree ,i ,j int) {
		ch := board[i][j]
		if ch == '#' {
			return
		}


		if tria.childs[ch - 'a'] == nil {
			// 不在字典中不在继续
			return
		}
		board[i][j] = '#'
		tria = tria.childs[ch - 'a']
		if tria.words != ""  {
			cache[tria.words] = true
		}

		//继续判断上下左右
		for _,direction := range fourDirections{
			x:= j + direction.x
			y:= i + direction.y
			if(y >= 0 && y < len(board) && x >= 0 && x < len(board[0]) && board[y][x] != '#'){
				dfs(tria,y,x)
			}
		}
		board[i][j] = ch
	}

	//开始执行
	for i,line := range board {
		for j,_:= range line {
			dfs(trieTree,i,j)
		}
	}


	//words = filterWord(board,words)
	//if len(words) == 0{
	//	return result
	//}

	for key,_ := range cache{
		result = append(result,key)
	}

	return result
}

//过滤明显不符合的word
func filterWord(board [][]byte, words []string)[]string{
	result := []string{}

	var letters = [26]bool{}

	//step 1 构造map
	for i,line := range board{
		for j,_:= range line{
			key:= board[i][j]
			letters[key - 'a'] = true
		}
	}
	//step 2 过滤
	for _,word := range words{
		validate := true
		for _,ch:= range word{
			if !letters[ch - 'a'] {
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
//func dfs(i int,j int ,board [][]byte,stack *Stack,tree *TirTree,cache map[string]string,wordCount int) bool {
//	data := board[i][j]
//	if string(data) == "#"{
//		return false
//	}
//
//	if !tree.hasString(string(stack.data),true){
//		return false
//	}
//
//	stack.push(data)
//	board[i][j] = byte('#')
//
//	if tree.hasString(string(stack.data),false){
//		cache[string(stack.data)] = ""
//		if len(cache) == wordCount{
//			return true
//		}
//	}
//
//	isLast:=true;
//
//	//上
//	if(i > 0){
//		if dfs(i - 1,j,board,stack,tree,cache,wordCount){
//			isLast = false;
//		}
//	}
//
//	//下
//	if(i < len(board) - 1){
//		if dfs(i + 1,j,board,stack,tree,cache,wordCount){
//			isLast = false;
//		}
//	}
//
//	//左
//	if(j > 0){
//		if dfs(i,j-1,board,stack,tree,cache,wordCount){
//			isLast = false;
//		}
//	}
//
//	//右
//	if(j < len(board[i]) - 1){
//		if dfs(i,j+1,board,stack,tree,cache,wordCount){
//			isLast = false;
//		}
//	}
//	if(isLast){
//		//stack.print()
//	}
//	stack.pop()
//	board[i][j] = data;
//	return true;
//}

func main(){

	var board [][]byte = [][]byte{[]byte("oaan"),[]byte("etae"),[]byte("ihkr"),[]byte("iflv")}
	var words [] string = [] string { "oath","pea","eat","rain" }

	//var board = [][]byte{
	//	[]byte("bababababa"),
	//	[]byte("ababababab"),
	//	[]byte("bababababa"),
	//	[]byte("ababababab"),
	//	[]byte("bababababa"),
	//	[]byte("ababababab"),
	//	[]byte("bababababa"),
	//	[]byte("ababababab"),
	//	[]byte("bababababa"),
	//	[]byte("ababababab"),
	//}
	//var words [] string = [] string { "ababababaa","ababababab","ababababac","ababababad","ababababae","ababababaf",
	//	"ababababag","ababababah","ababababai","ababababaj","ababababak","ababababal","ababababam","ababababan",
	//	"ababababao","ababababap","ababababaq","ababababar","ababababas","ababababat","ababababau","ababababav",
	//	"ababababaw","ababababax","ababababay","ababababaz","ababababba","ababababbb","ababababbc","ababababbd",
	//	"ababababbe","ababababbf","ababababbg","ababababbh","ababababbi","ababababbj","ababababbk","ababababbl",
	//	"ababababbm","ababababbn","ababababbo","ababababbp","ababababbq","ababababbr","ababababbs","ababababbt",
	//	"ababababbu","ababababbv","ababababbw","ababababbx","ababababby","ababababbz","ababababca","ababababcb",
	//	"ababababcc","ababababcd","ababababce","ababababcf","ababababcg","ababababch","ababababci","ababababcj",
	//	"ababababck","ababababcl","ababababcm","ababababcn","ababababco","ababababcp","ababababcq","ababababcr",
	//	"ababababcs","ababababct","ababababcu","ababababcv","ababababcw","ababababcx","ababababcy","ababababcz",
	//	"ababababda","ababababdb","ababababdc","ababababdd","ababababde","ababababdf","ababababdg","ababababdh",
	//	"ababababdi","ababababdj","ababababdk","ababababdl","ababababdm","ababababdn","ababababdo","ababababdp",
	//	"ababababdq","ababababdr","ababababds","ababababdt","ababababdu","ababababdv","ababababdw","ababababdx",
	//	"ababababdy","ababababdz","ababababea","ababababeb","ababababec","ababababed","ababababee","ababababef",
	//	"ababababeg","ababababeh","ababababei","ababababej","ababababek","ababababel","ababababem","ababababen",
	//	"ababababeo","ababababep","ababababeq","ababababer","ababababes","ababababet","ababababeu","ababababev",
	//	"ababababew","ababababex","ababababey","ababababez","ababababfa","ababababfb","ababababfc","ababababfd",
	//	"ababababfe","ababababff","ababababfg","ababababfh","ababababfi","ababababfj","ababababfk","ababababfl",
	//	"ababababfm","ababababfn","ababababfo","ababababfp","ababababfq","ababababfr","ababababfs","ababababft",
	//	"ababababfu","ababababfv","ababababfw","ababababfx","ababababfy","ababababfz","ababababga","ababababgb",
	//	"ababababgc","ababababgd","ababababge","ababababgf","ababababgg","ababababgh","ababababgi","ababababgj",
	//	"ababababgk","ababababgl","ababababgm","ababababgn","ababababgo","ababababgp","ababababgq","ababababgr",
	//	"ababababgs","ababababgt","ababababgu","ababababgv","ababababgw","ababababgx","ababababgy","ababababgz",
	//	"ababababha","ababababhb","ababababhc","ababababhd","ababababhe","ababababhf","ababababhg","ababababhh",
	//	"ababababhi","ababababhj","ababababhk","ababababhl","ababababhm","ababababhn","ababababho","ababababhp",
	//	"ababababhq","ababababhr","ababababhs","ababababht","ababababhu","ababababhv","ababababhw","ababababhx",
	//	"ababababhy","ababababhz","ababababia","ababababib","ababababic","ababababid","ababababie","ababababif",
	//	"ababababig","ababababih","ababababii","ababababij","ababababik","ababababil","ababababim","ababababin",
	//	"ababababio","ababababip","ababababiq","ababababir","ababababis","ababababit","ababababiu","ababababiv",
	//	"ababababiw","ababababix","ababababiy","ababababiz","ababababja","ababababjb","ababababjc","ababababjd",
	//	"ababababje","ababababjf","ababababjg","ababababjh","ababababji","ababababjj","ababababjk","ababababjl",
	//	"ababababjm","ababababjn","ababababjo","ababababjp","ababababjq","ababababjr","ababababjs","ababababjt",
	//	"ababababju","ababababjv","ababababjw","ababababjx","ababababjy","ababababjz","ababababka","ababababkb",
	//	"ababababkc","ababababkd","ababababke","ababababkf","ababababkg","ababababkh","ababababki","ababababkj",
	//	"ababababkk","ababababkl","ababababkm","ababababkn","ababababko","ababababkp","ababababkq","ababababkr",
	//	"ababababks","ababababkt","ababababku","ababababkv","ababababkw","ababababkx","ababababky","ababababkz",
	//	"ababababla","abababablb","abababablc","ababababld","abababable","abababablf","abababablg","abababablh",
	//	"ababababli","abababablj","abababablk","ababababll","abababablm","ababababln","abababablo","abababablp",
	//	"abababablq","abababablr","ababababls","abababablt","abababablu","abababablv","abababablw","abababablx",
	//	"abababably","abababablz","ababababma","ababababmb","ababababmc","ababababmd","ababababme","ababababmf",
	//	"ababababmg","ababababmh","ababababmi","ababababmj","ababababmk","ababababml","ababababmm","ababababmn",
	//	"ababababmo","ababababmp","ababababmq","ababababmr","ababababms","ababababmt","ababababmu","ababababmv",
	//	"ababababmw","ababababmx","ababababmy","ababababmz","ababababna","ababababnb","ababababnc","ababababnd",
	//	"ababababne","ababababnf","ababababng","ababababnh","ababababni","ababababnj","ababababnk","ababababnl",
	//	"ababababnm","ababababnn","ababababno","ababababnp","ababababnq","ababababnr","ababababns","ababababnt",
	//	"ababababnu","ababababnv","ababababnw","ababababnx","ababababny","ababababnz","ababababoa","ababababob",
	//	"ababababoc","ababababod","ababababoe","ababababof","ababababog","ababababoh","ababababoi","ababababoj",
	//	"ababababok","ababababol","ababababom","ababababon","ababababoo","ababababop","ababababoq","ababababor",
	//	"ababababos","ababababot","ababababou","ababababov","ababababow","ababababox","ababababoy","ababababoz",
	//	"ababababpa","ababababpb","ababababpc","ababababpd","ababababpe","ababababpf","ababababpg","ababababph",
	//	"ababababpi","ababababpj","ababababpk","ababababpl","ababababpm","ababababpn","ababababpo","ababababpp",
	//	"ababababpq","ababababpr","ababababps","ababababpt","ababababpu","ababababpv","ababababpw","ababababpx",
	//	"ababababpy","ababababpz","ababababqa","ababababqb","ababababqc","ababababqd","ababababqe","ababababqf",
	//	"ababababqg","ababababqh","ababababqi","ababababqj","ababababqk","ababababql","ababababqm","ababababqn",
	//	"ababababqo","ababababqp","ababababqq","ababababqr","ababababqs","ababababqt","ababababqu","ababababqv",
	//	"ababababqw","ababababqx","ababababqy","ababababqz","ababababra","ababababrb","ababababrc","ababababrd",
	//	"ababababre","ababababrf","ababababrg","ababababrh","ababababri","ababababrj","ababababrk","ababababrl",
	//	"ababababrm","ababababrn","ababababro","ababababrp","ababababrq","ababababrr","ababababrs","ababababrt",
	//	"ababababru","ababababrv","ababababrw","ababababrx","ababababry","ababababrz","ababababsa","ababababsb",
	//	"ababababsc","ababababsd","ababababse","ababababsf","ababababsg","ababababsh","ababababsi","ababababsj",
	//	"ababababsk","ababababsl","ababababsm","ababababsn","ababababso","ababababsp","ababababsq","ababababsr",
	//	"ababababss","ababababst","ababababsu","ababababsv","ababababsw","ababababsx","ababababsy","ababababsz",
	//	"ababababta","ababababtb","ababababtc","ababababtd","ababababte","ababababtf","ababababtg","ababababth",
	//	"ababababti","ababababtj","ababababtk","ababababtl","ababababtm","ababababtn","ababababto","ababababtp",
	//	"ababababtq","ababababtr","ababababts","ababababtt","ababababtu","ababababtv","ababababtw","ababababtx",
	//	"ababababty","ababababtz","ababababua","ababababub","ababababuc","ababababud","ababababue","ababababuf",
	//	"ababababug","ababababuh","ababababui","ababababuj","ababababuk","ababababul","ababababum","ababababun",
	//	"ababababuo","ababababup","ababababuq","ababababur","ababababus","ababababut","ababababuu","ababababuv",
	//	"ababababuw","ababababux","ababababuy","ababababuz","ababababva","ababababvb","ababababvc","ababababvd",
	//	"ababababve","ababababvf","ababababvg","ababababvh","ababababvi","ababababvj","ababababvk","ababababvl",
	//	"ababababvm","ababababvn","ababababvo","ababababvp","ababababvq","ababababvr","ababababvs","ababababvt",
	//	"ababababvu","ababababvv","ababababvw","ababababvx","ababababvy","ababababvz","ababababwa","ababababwb",
	//	"ababababwc","ababababwd","ababababwe","ababababwf","ababababwg","ababababwh","ababababwi","ababababwj",
	//	"ababababwk","ababababwl","ababababwm","ababababwn","ababababwo","ababababwp","ababababwq","ababababwr",
	//	"ababababws","ababababwt","ababababwu","ababababwv","ababababww","ababababwx","ababababwy","ababababwz",
	//	"ababababxa","ababababxb","ababababxc","ababababxd","ababababxe","ababababxf","ababababxg","ababababxh",
	//	"ababababxi","ababababxj","ababababxk","ababababxl","ababababxm","ababababxn","ababababxo","ababababxp",
	//	"ababababxq","ababababxr","ababababxs","ababababxt","ababababxu","ababababxv","ababababxw","ababababxx",
	//	"ababababxy","ababababxz","ababababya","ababababyb","ababababyc","ababababyd","ababababye","ababababyf",
	//	"ababababyg","ababababyh","ababababyi","ababababyj","ababababyk","ababababyl","ababababym","ababababyn",
	//	"ababababyo","ababababyp","ababababyq","ababababyr","ababababys","ababababyt","ababababyu","ababababyv",
	//	"ababababyw","ababababyx","ababababyy","ababababyz","ababababza","ababababzb","ababababzc","ababababzd",
	//	"ababababze","ababababzf","ababababzg","ababababzh","ababababzi","ababababzj","ababababzk","ababababzl",
	//	"ababababzm","ababababzn","ababababzo","ababababzp","ababababzq","ababababzr","ababababzs","ababababzt",
	//	"ababababzu","ababababzv","ababababzw","ababababzx","ababababzy","ababababzz" }


	finded := findWords1(board,words)
	date,error := json.Marshal(finded)
	if(error != nil){
		log.Fatalf("JSON marshaling failed: %s" , error)
	}
	fmt.Printf("%s\n",date)


}

