# [67. Add Binary][title]

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
>给你两个二进制串，求其和的二进制串。

## 题解

### 思路1
> 按照小学算数那么来做，用 `carry` 表示进位，从后往前算，依次往前，每算出一位就插入到最前面即可，直到把两个二进制串都遍历完即可。

```go
func addBinary(a, b string) string {
	lenA := len(a)
	lenB := len(b)
	if lenB > lenA {
		a, b = b, a             //swap
		lenA, lenB = lenB, lenA //swap2
	}
	b = strings.Repeat("0", lenA-lenB) + b
	lenB = len(b)

	carry := byte(0)
	ret := make([]byte, lenB+1)

	for i := lenB - 1; i >= 0; i-- {
		numA := a[i] - '0'
		numB := b[i] - '0'
		sum := numA + numB + carry
		ret[i+1] = sum&1 + '0'
		carry = sum >> 1
	}
	if carry == 0 {
		ret = ret[1:]
	} else {
		ret[0] = carry + '0'
	}
	return string(ret)
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/add-binary/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
