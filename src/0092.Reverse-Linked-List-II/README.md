# [92. Reverse Linked List II][title]

## Description

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

**Example 1:**

```
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
```

**Tags:** Math, String

## 题意
>请将链表中第 nn 个节点和第 mm 个节点之间的部分翻转。
 链表最多只能遍历一遍。

## 题解

### 思路1
> 和problem 206类似

```go
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := new(ListNode)
	head, dummy.Next = dummy, head

	for i := 0; i < m-1; i++ {
		head = head.Next
	}
	var curr, prev *ListNode = head.Next, nil
	for i := 0; i < n-m+1; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next.Next = curr
	head.Next = prev

	return dummy.Next
}

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
