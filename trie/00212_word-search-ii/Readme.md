
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/word-search-ii

给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。


输入：
```
board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
```
输出：
```
["eat","oath"]
```

输入：
```
board = [["a","b"],["c","d"]], words = ["abcb"]
```
输出：
```
[]
```

提示：
    

* m == board.length
* n == board[i].length
* 1 <= m, n <= 12
* board[i][j] 是一个小写英文字母
* 1 <= words.length <= 3 * 10^4
* 1 <= words[i].length <= 10
* words[i] 由小写英文字母组成
* words 中的所有字符串互不相同


总结：
* 单词Trie,来源于retrieval(检索)中取中间四个字符构成(读音同try)。
* 矩阵循环的方式
```go
for i, row := range board {
    for j := range row {
        dfs(t, i, j)
    }
}
```
* 访问某个节点上下左右的方式
```go
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
for _, d := range dirs {
    nx, ny := x+d.x, y+d.y
    if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
        dfs(node, nx, ny)
    }
}
```
* 将dfs定义在方法内部，减少参数传递
* Trie 的定义 和 string 的遍历
```go
type Trie struct {
	// 数组指定长度可以初始化，如果换成map不会进行初始化
	children [26]*Trie
	word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}
```
* 为了记录当前递归的轨迹，增加了stack，实际上，直接把Trie 的node 当作 dfs 的入参即可
* 为什么 Trie 的定义 的时候，数组是 *，方法名称前边括号里必须加 *？
```go
1. invalid recursive type X的原因是:当创建类型 X 时,要计算这个类型的大小,而编译器发现类型 X 的 next 字段也是 X 类型,一个同样的还没有确定大小的字段,于是我们会陷入一个无穷递归中.但是如果是一个指针类,编译器就能知道在指定平台上指针类型的确定大小.
2. struct 传递的是副本
```
* 数组定义 var x = [] type{} 是大括号不是中括号
* 方法体内定义递归调用
```go
    //先声明
	var dfs func(node *TrieTree, x, y int)
    //再定义
    dfs = func(tria *TrieTree ,i ,j int) {
    	//这样不会报错找不到变量
        dfs(tria,i,j)
    }
```