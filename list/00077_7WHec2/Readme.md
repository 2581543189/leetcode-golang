
来源：力扣（LeetCode）  
链接：https://leetcode-cn.com/problems/7WHec2/

给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

输入：
```
head = [4,2,1,3]
```
输出：
```
[1,2,3,4]
```

输入：
```
head = [-1,5,3,4,0]
```
输出：
```
[-1,0,3,4,5]
```

输入：
```
head = []
```
输出：
```
[]
```

提示：
    

* 链表中节点的数目在范围 [0, 5 * 104] 内
* -105 <= Node.val <= 105


总结：

* 使用 快慢指针寻找list 中点的时候 fast = head.next
* 既要获取list头又要获取list尾的情况下，创建一个虚拟头，最后删除 使用curr 记录当前值