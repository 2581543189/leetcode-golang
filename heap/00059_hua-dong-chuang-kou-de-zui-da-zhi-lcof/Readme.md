
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/

给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值


输入：
```
nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
```
输出：
```
[3,3,5,5,6,7]
```

提示：
    

* 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

总结：

* type hp struct{ sort.IntSlice } 可以省略 swap 和 len
* heap 没有peek 方法，因为可以直接从数组里获取第 0 项