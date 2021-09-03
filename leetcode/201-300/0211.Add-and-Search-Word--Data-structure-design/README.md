# [211.Add and Search Word - Data structure design][title]

## Description

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

  *  `WordDictionary()` Initializes the object.
  *  `void addWord(word)` Adds `word` to the data structure, it can be matched later.
  *  `bool search(word)` Returns `true` if there is any string in the data structure that matches `word` or `false` otherwise. `word` may contain dots `'.'` where dots can be matched with any letter.

## NOTES
- `1 <= word.length <= 500`
- `word` in `addWord` consists lower-case English letters.
- `word` in `search` consist of  `'.'` or lower-case English letters.
- At most 50000 calls will be made to `addWord` and `search`.

**Tags:** String, DFS, Trie, Design

**Example 1:**

```
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]
```

## 题意
> ...

## 题解

### 思路1
> ...
Add and Search Word - Data structure design
```go
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/add-and-search-word-data-structure-design/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
