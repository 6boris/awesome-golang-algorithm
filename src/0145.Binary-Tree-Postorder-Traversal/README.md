# [144. Binary Tree Preorder Traversal][title]

## Description

Given a binary tree, return the postorder traversal of its nodes' values.

**Example 1:**

```
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
```
Follow up: Recursive solution is trivial, could you do it iteratively?

**Tags:** Tree Stack

## 题意
>

## 题解

### 思路1
> 使用两个栈，实现数据顺序处理，代码直接引用`leetcode`

### 思路2
> 使用top计数和切片存储数据，定义临时的数据接口
```go
type seqStack struct {
	data []*Node
	tag []int // 后续遍历准备
	top int // 数组下标
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/linked-list-cycle-ii/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
