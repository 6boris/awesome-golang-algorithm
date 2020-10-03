# [1. Add Sum][title]

## Description

Given a string *s* consists of upper/lower-case alphabets and empty space characters `' '`, return the length of last word in the string.

If the last word does not exist, return 0.

**Note:** A word is defined as a character sequence consists of non-space characters only.

**Example:**

```
Input: "Hello World"
Output: 5
```

**Tags:** String


## 题解
### 思路1
从后向前找，去除末尾的空格，再向前寻找空格
```go

func lengthOfLastWord(s string) int {
	start := len(s) -1
	//	寻找最后一个单词的尾下标
	for start >= 0 && s[start] == ' ' {
		start--
	}
	//	寻找最后一个单词的初始下标
	end := start
	for  start >= 0 && s[start] != ' '{
		start--
	}

	return end - start
}
```
### 思路2
思路一样，但是用的是自带的strings高级函数完成的
```go
func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	return len(s) - strings.LastIndex(s, " ") - 1
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/length-of-last-word/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
