# [232. Implement Queue using Stacks][title]

## Description

Implement a trie with `insert`, `search`, and `startsWith` methods.

**Example:**

```
Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
```

## NOTES
- You may assume that all inputs are consist of lowercase letters `a-z`.
- All inputs are guaranteed to be non-empty strings.

**Tags:** Trie Tree

## 题意
> 封装`Trie`树结构
### 思路1
使用`map`来保存每一层的数据
```go
type Trie struct {
	IsTerminated bool // 是否是叶子节点
	Children     map[rune]*Trie // 某一层的Trie树结构
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-letcode][me]

[title]: https://leetcode.com/problems/valid-anagram/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
