# [97. Interleaving String][title]

## Description

Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
**Example 1:**

```
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
```

**Example 2:**

```
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
```

**Tags:** Math, String

## 题意
>给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
 

## 题解

### 思路1
>

```go
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true
	for i := 1; i < len(s1)+1; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j < len(s2)+1; j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			dp[i][j] = (s1[i-1] == s3[i+j-1] && dp[i-1][j]) || (s2[j-1] == s3[i+j-1] && dp[i][j-1])
		}
	}
	return dp[len(s1)][len(s2)]
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/interleaving-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
