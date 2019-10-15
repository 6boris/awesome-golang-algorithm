# [1. Add Sum][title]

## Description
一个只包含 A-Z 的消息可以用如下方式编码成数字：
```bash
'A' -> 1
'B' -> 2
...
'Z' -> 26
```
给定一个只包含数字的非空字符串，返回共有多少种解码方案。
**Example 1:**

```
Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
```

**Example 2:**

```
Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
```

**Tags:** Math, String

## 题意
>给你两个二进制串，求其和的二进制串。

## 题解

### 思路1
> (动态规划) O(n)
>
这道题目可以用动态规划来做。
状态表示：f[i]f[i] 表示前 ii 个数字共有多少种解码方式。
初始化：0个数字解码的方案数1，即 f[0]=1f[0]=1。
状态转移：f[i]f[i] 可以表示成如下两部分的和：

- 如果第 ii 个数字不是0，则 ii 个数字可以单独解码成一个字母，此时的方案数等于用前 i−1i−1 个数字解码的方案数，即 f[i−1]f[i−1]；
- 如果第 i−1i−1个数字和第 ii 个数字组成的两位数在 1010 到 2626 之间，则可以将这两位数字解码成一个字符，此时的方案数等于用前 i−2i−2 个数字解码的方案数，即 f[i−2]f[i−2]；
时间复杂度分析：状态数是 nn 个，状态转移的时间复杂度是 O(1)O(1)，所以总时间复杂度是 O(n)O(n)。

```go
func numDecodings(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] < '0' || s[i-1] > '9' {
			return 0
		}
		f[i] = 0
		if s[i-1] != '0' {
			f[i] = f[i-1]
		}
		if i > 1 {
			t := (s[i-2]-'0')*10 + s[i-1] - '0'
			if t >= 10 && t <= 26 {
				f[i] += f[i-2]
			}
		}
	}
	return f[n]
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
