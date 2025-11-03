# [1410.HTML Entity Parser][title]

## Description
**HTML entity parser** is the parser that takes HTML code as input and replace all the entities of the special characters by the characters itself.

The special characters and their entities for HTML are:

- **Quotation Mark**: the entity is `&quot;` and symbol character is `"`.
- **Single Quote Mark**: the entity is `&apos;` and symbol character is `'`.
- **Ampersand**: the entity is `&amp;` and symbol character is `&`.
- **Greater Than Sign**: the entity is `&gt;` and symbol character is `>`.
- **Less Than Sign**: the entity is `&lt;` and symbol character is `<`.
- **Slash**: the entity is `&frasl;` and symbol character is `/`.

Given the input `text` string to the HTML parser, you have to implement the entity parser.

Return the text after replacing the entities by the special characters.

**Example 1:**

```
Input: text = "&amp; is an HTML entity but &ambassador; is not."
Output: "& is an HTML entity but &ambassador; is not."
Explanation: The parser will replace the &amp; entity by &
```

**Example 2:**

```
Input: text = "and I quote: &quot;...&quot;"
Output: "and I quote: \"...\""
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/html-entity-parser
[me]: https://github.com/kylesliu/awesome-golang-algorithm
