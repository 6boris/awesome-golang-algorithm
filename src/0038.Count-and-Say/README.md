# [38. Count and Say][title]

## Description

The count-and-say sequence is the sequence of integers with the first five terms as following:

```
1.     1
2.     11
3.     21
4.     1211
5.     111221
```

`1` is read off as `"one 1"` or `11`.

`11` is read off as `"two 1s"` or `21`.

`21` is read off as `"one 2`, then `one 1"` or `1211`.

Given an integer *n*, generate the *n*<sup>th</sup> term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

**Example 1:**

```
Input: 1
Output: "1"
```

**Example 2:**

```
Input: 4
Output: "1211"
```

**Tags:** String


## 题解
### 思路1
在遍历数组时用 Stack 把数组中的数存起来，如果当前遍历的数比栈顶元素来的大，说明栈顶元素的下一个比它大的数就是当前元素。
```go
func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	say := []byte("1")
	for i := 0; i < n-1; i++ {
		var cnt byte
		var nextSay []byte
		for j := 0; j < len(say); j++ {
			cnt++
			if j == len(say)-1 || say[j] != say[j+1] {
				nextSay = append(nextSay, 48 + cnt, say[j]) // 48 - the ASCII code point of "0"
				cnt = 0
			}
		}
		say = nextSay
	}
	return string(say)
}
```
### 思路2


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/count-and-say/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
