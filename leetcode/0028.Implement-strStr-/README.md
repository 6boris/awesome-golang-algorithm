# [28. Implement strStr()][title]

## Description

Implement [strStr()](http://www.cplusplus.com/reference/cstring/strstr/).

Return the index of the first occurrence of needle in haystack, or **-1** if needle is not part of haystack.

**Example 1:**

```
Input: haystack = "hello", needle = "ll"
Output: 2
```

**Example 2:**

```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

**Clarification:**

What should we return when `needle` is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when `needle` is an empty string. This is consistent to C's [strstr()](http://www.cplusplus.com/reference/cstring/strstr/) and Java's [indexOf()](https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)).

Tags:** Two Pointers, String


## 题解
### 思路1
题意是从主串中找到子串的索引，如果找不到则返回-1，当子串长度大于主串，直接返回-1，然后我们只需要遍历比较即可。
```go
func strStr(haystack string, needle string) int {
	//	检查数据
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}

		if j == len(needle) {
			return i
		}
	}
	return -1
}
```

### 思路 2
KMP 算法
```go
func strStr(haystack string, needle string) int {
    prefixTable := calcPrefixTable(needle)
    
    i, j := 0, 0
    for i < len(haystack) && j < len(needle) {
        if -1 == j || haystack[i] == needle[j] {
            i++
            j++
        } else {
            if j == 0 {
                i++
            } else {
            	j = prefixTable[j]
            }
        }
    }
    if j == len(needle) {
        return i - j
    }
    return -1
}

func calcPrefixTable(str string) []int {
    next := make([]int, len(str)+1)
    length := 0
    
    for i := 2; i <= len(str); i++ {
        for length > 0 && str[length] != str[i-1] {
            length = next[length]
        }
        if str[length] == str[i-1] {
            length++
        }
        next[i] = length;
    }
    
    return next
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/implement-strstr/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
