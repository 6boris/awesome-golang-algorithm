# [1. Add Sum][title]

## Description

Given two binary strings, return their sum (also a binary string).

The input strings are both **non-empty** and contains only characters `1` or `0`.

**Example 1:**

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2:**

```
Input: a = "1010", b = "1011"
Output: "10101"
```

**Tags:** Math, String

## 题意
>写一个高效算法，在矩阵中查找一个数是否存在。矩阵有如下特点：
- 矩阵中每行的数，从左到右单调递增；
- 每行行首的数大于上一行行尾的数；

## 题解

### 思路1
> 我们可以想象把整个矩阵，按行展开成一个一维数组，那么这个一维数组单调递增，然后直接二分即可。
  二分时可以通过整除和取模运算得到二维数组的坐标。
  时间复杂度分析：二分的时间复杂度是 O(logn2)=O(logn)O(logn2)=O(logn)。


```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	l, r := 0, n*m-1

	for l < r {
		mid := (l + r) / 2
		if matrix[mid/m][mid%m] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return matrix[r/m][r%m] == target
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
