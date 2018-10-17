# [50. Pow(x, n)][title]

## Description

Implement [pow(*x*, *n*)](http://www.cplusplus.com/reference/valarray/pow/), which calculates *x* raised to the power *n* (xn).

**Example 1:**

```
Input: 2.00000, 10
Output: 1024.00000
```

**Example 2:**

```
Input: 2.10000, 3
Output: 9.26100
```

**Example 3:**

```
Input: 2.00000, -2
Output: 0.25000
Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
```

**Note:**

- -100.0 < *x* < 100.0
- *n* is a 32-bit signed integer, within the range [−2<sup>31</sup>, 2<sup>31</sup> − 1]

**Tags:** Math, Binary Search

## 题意
>计算 x^n

## 题解

### 思路1
> 如果直接计算肯定会超时，那么我们可以想到可以使用二分法来降低时间复杂度。

```go
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / pow(x, -n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	ans := pow(x, n>>1)
	if n&1 == 0 {
		return ans * ans
	}
	return ans * ans * x
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/powx-n/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
