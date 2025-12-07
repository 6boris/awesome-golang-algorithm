# [794.Valid Tic-Tac-Toe State][title]

## Description

Given a Tic-Tac-Toe `board` as a string array board, return `true` if and only if it is possible to reach this board position during the course of a valid tic-tac-toe game.

The board is a `3 x 3` array that consists of characters `' '`, `'X'`, and `'O'`. The `' '` character represents an empty square.

Here are the rules of Tic-Tac-Toe:

- Players take turns placing characters into empty squares `' '`.
- The first player always places `'X'` characters, while the second player always places `'O'` characters.
- `'X'` and `'O'` characters are always placed into empty squares, never filled ones.
- The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.
- The game also ends if all squares are non-empty.
- No more moves can be played if the game is over.

**Example 1:**  

![1](./1.jpg)

```
Input: board = ["O  ","   ","   "]
Output: false
Explanation: The first player always plays "X".
```

**Example 2:**  

![2](./2.jpg)

```
Input: board = ["XOX"," X ","   "]
Output: false
Explanation: Players take turns making moves.
```

**Example 3:**  

![3](./3.jpg)

```
Input: board = ["XOX","O O","XOX"]
Output: true
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/valid-tic-tac-toe-state/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
