
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。


输入：
```
[7,5,6,4]
```
输出：
```
5
```

提示：
    

* 0 <= 数组长度 <= 50000

总结：

* 树状数组或二叉索引树BIT（英语：Binary Indexed Tree)
* bit 的 下标从 1 开始，数组长度也是 n+1 ,第0项不使用
* 离散化
```golang
    //离散化
    tmp:= make([]int,n)
    copy(tmp,nums)
	sort.Ints(tmp)
	for i:=0;i<n;i++{
		nums[i] = sort.SearchInts(tmp,nums[i]) + 1 //不能有0
	}
```