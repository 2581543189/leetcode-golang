
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/find-the-duplicate-number

给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。


示例：
```
输入：nums = [1,3,4,2,2]
输出：2
```
示例：
```
输入：nums = [3,1,3,4,2]
输出：3
```

示例：
```
输入：nums = [1,1]
输出：1
```

示例：
```
输入：nums = [1,1,2]
输出：1
```

提示：
* 1 <= n <= 10^5
* nums.length == n + 1
* 1 <= nums[i] <= n
* nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次


总结：
* 计算总位数
```golang
    bit_max := 31
    for ((n-1) >> bit_max) == 0 {
        bit_max--
    }
    bit_max++
```
* 判断某一位是否 1
```golang
if val & ( 1 << i) > 0
```
* 某一位设置为 1
```golang
ans |= ( 1 << i)
```
* 将数组看成是链表
```golang
	s,f:= nums[0],nums[nums[0]]
    //数组的坐标是链表节点的val
	//数组的值是next
	所以最后return的时候要return 坐标
```

