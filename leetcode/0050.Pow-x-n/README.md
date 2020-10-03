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

```go
    X^27
    = X * X ^26
    = X * (X^2)^13
    = X * (X^2) * (X^2)^12
    = X * (X^2) * (X^4)^6
    = X * (X^2) * (X^8)^3
    = X * (X^2) * (X^8) * (X^8)^2
    = X * (X^2) * (X^8) * (X^16)
    3.n:    偶 y = x ^ n/2  ans = y * y
            奇 y = x ^ n/2  ans = y * y * x
```
> 利用递归或者分治的方法求解

```go
func myPow(x float64, n int) float64 {
	//	判断递归结束
	if n == 0 {
		return 1
	}
	//	负数:直接求导
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	//	奇数
	if n%2 == 1 {
		return x * myPow(x, n-1)
	}
	//	偶数
	return myPow(x*x, n/2)
}
```

### 思路2
> 思路2
```go
func myPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	pow := 1.0
	for n != 0 {
		if n&1 != 0 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/powx-n/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
