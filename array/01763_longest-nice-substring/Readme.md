
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/longest-nice-substring

当一个字符串 s 包含的每一种字母的大写和小写形式 同时 出现在 s 中，就称这个字符串 s 是 美好 字符串。比方说，"abABB" 是美好字符串，因为 'A' 和 'a' 同时出现了，且 'B' 和 'b' 也同时出现了。然而，"abA" 不是美好字符串因为 'b' 出现了，而 'B' 没有出现。

给你一个字符串 s ，请你返回 s 最长的 美好子字符串 。如果有多个答案，请你返回 最早 出现的一个。如果不存在美好子字符串，请你返回一个空字符串。


示例：
```
输入：s = "YazaAay"
输出："aAa"
解释："aAa" 是一个美好字符串，因为这个子串中仅含一种字母，其小写形式 'a' 和大写形式 'A' 也同时出现了。
"aAa" 是最长的美好子字符串。
```
示例：
```
输入：s = "Bb"
输出："Bb"
解释："Bb" 是美好字符串，因为 'B' 和 'b' 都出现了。整个字符串也是原字符串的子字符串。
```
示例：
```
输入：s = "c"
输出：""
解释：没有美好子字符串。
```
示例：
```
输入：s = "dDzeE"
输出："dD"
解释："dD" 和 "eE" 都是最长美好子字符串。
由于有多个美好子字符串，返回 "dD" ，因为它出现得最早。
```


提示：

* 1 <= s.length <= 100
* s 只包含大写和小写英文字母。


总结：
* ascii 'A' 要比 'a'  要小。如果全部是字母，可以用 x < 'a' 判断是否大写
* 全部都是字母的时候，统计是否存在可以用两个 int32 位运算
* unicode.ToLower(rune(s[i]))
* string 通过遍历，截成多个字串的方式