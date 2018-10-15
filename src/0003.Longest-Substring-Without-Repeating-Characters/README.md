# [3. Longest Substring Without Repeating Characters][title]

## Description

## Description

Given a string, find the length of the **longest substring** without repeating characters.

**Examples:**

Given `"abcabcbb"`, the answer is `"abc"`, which the length is 3.

Given `"bbbbb"`, the answer is `"b"`, with the length of 1.

Given `"pwwkew"`, the answer is `"wke"`, with the length of 3. Note that the answer must be a **substring**, `"pwke"` is a *subsequence* and not a substring.

**Tags:** Hash Table, Two Pointers, String

## 题意
>计算不带重复字符的最长子字符串的长度

## 题解

### 思路1
> 开辟一个 hash 数组来存储该字符上次出现的位置，比如 `hash[a] = 3` 就是代表 `a` 字符前一次出现的索引在 3，遍历该字符串，获取到上次出现的最大索引（只能向前，不能退后），与当前的索引做差获取的就是本次所需长度，从中迭代出最大值就是最终答案。

```go
// O(n) time O(1) space Solution
func lengthOfLongestSubstring(s string) int {
	var chPosition [256]int // [0, 0, 0, ...]
	maxLength, substringLen, lastRepeatPos := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if pos := chPosition[s[i]]; pos > 0 {
			// record current substring length
			maxLength = Max(substringLen, maxLength)

			// update characters position
			chPosition[s[i]] = i + 1

			// update last repeat character position
			lastRepeatPos = Max(pos, lastRepeatPos)

			// update the current substring from last repeat character
			substringLen = i + 1 - lastRepeatPos
		} else {
			substringLen += 1
			chPosition[s[i]] = i + 1
		}
	}

	return Max(maxLength, substringLen)
}
```

### 思路2
> 暴力循环
```go
//	暴力求解(会超时)
func lengthOfLongestSubstring2(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if allUnique(s, i, j) {
				ans = Max(ans, j-i)
			}
		}
	}
	return ans
}

func allUnique(s string, start int, end int) bool {
	sMap := make(map[string]int)

	for i := start; i < end; i++ {
		if sMap[string(s[i])] > 0 {
			return false
		}
		sMap[string(s[i])]++
	}

	return true
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

```

## Benchmark
<div align=center>
<img src="https://github.com/kylesliu/awesome-golang-leetcode/blob/master/assets/images/0003-BenchMark.png"></img>
</div>

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode

[Benchmark]: https://github.com/kylesliu/awesome-golang-leetcode/blob/master/assets/images/0003-BenchMark.png
