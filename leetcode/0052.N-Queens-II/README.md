# [52. N-Queens II][title]

## Description

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
![](https://assets.leetcode.com/uploads/2018/10/12/8-queens.png)
Given an integer n, return the number of distinct solutions to the n-queens puzzle.
**Example:**

```
Input: 4
Output: 2
Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
[
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
```

**Tags:** DFS, 位运算

## 题意
> 给一个整数n, 计算在n矩阵格子中能够放Queen的个数

## 题解

### 思路1
> 直接使用三个切片，分别存储列，撇，捺的方式，比较占用空间

### 思路2
> 直接使用一个切片，通过该切片元素直接数据换算判断

### 思路3
> 使用位运算
```go
// 关键代码
func dfs(n int, row int, cols int, pie int, na int) {
	if row >= n {
		count++
		return
	}
	bits := (^(cols | pie | na) & ((1 << uint(n)) - 1)) // 查看当前能够放置元素的位置
	for bits > 0 {
		p := bits & -bits // 获取最低位1
		dfs(n, row+1, cols|p, (pie|p)>>1, (na|p)<<1)
		bits = bits & (bits - 1) // 去掉最低位1
	}
}
```

## 参考文档
- [用位运算速解 n 皇后问题](https://zhuanlan.zhihu.com/p/22846106)

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
