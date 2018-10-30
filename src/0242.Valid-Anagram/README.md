# [242. Valid Anagram][title]

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
>求异位词【字母相同，只是循序不同】

## 题解

### 思路1
> 先排序再比较  时间复杂度O(NlogN)

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s, t = InsertSort(s), InsertSort(t)

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

// 插入排序(Insert Sort)
func InsertSort(arr string) string {
	strs := []byte(arr)

	i, j := 0, 0
	for i = 1; i < len(strs); i++ {
		tmp := strs[i]
		for j = i; j > 0 && strs[j-1] > tmp; j-- {
			strs[j] = strs[j-1]
		}
		strs[j] = tmp
	}

	return string(strs)
}
```

### 思路2
> 使用Map 统计  时间复杂度O(N)
```go
//	用Map
func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap1 := make(map[uint8]int)
	sMap2 := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		sMap1[s[i]] += 1
		sMap2[t[i]] += 1
	}
	if !reflect.DeepEqual(sMap1, sMap2) {
		return false
	}
	return true
}
```
### 思路3
> 一种奇怪的方法，现在还没明白
```go
func isAnagram(s string, t string) bool {
	var xor, squaresum1, squaresum2 rune
	for _, char := range s {
		xor ^= char
		squaresum1 += char * char
	}
	for _, char := range t {
		xor ^= char
		squaresum2 += char * char
	}
	return xor == 0 && squaresum1 == squaresum2
}
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/valid-anagram/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
