# [1. Add Sum][title]

## Description

Given an input string, reverse the string word by word.

**Example 1:**

```
Input: "the sky is blue"
Output: "blue is sky the"
```

**Example 2:**

```
Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
```

**Example 2:**

```
Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
```
**Tags:** Math, String



## 题意
> 给定一个字符串，请将字符串中单词的顺序倒过来。
> 注意：

- 单词是指连续非空格字符；
- 单词之间可能包含多余空格，请在结果中删除多余空格，使得单词之间仅包含一个空格；
- 请删除字符串首尾多余的空格；

## 题解

### 思路1
> 。。。。

```go
func reverseWords(s string) string {
	//	用自带的函数拆分数组
	words := strings.Fields(s)

	i := 0
	j := len(words) - 1

	for i < j {
		//	左右22交换
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words, " ")
}

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
