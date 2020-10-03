# [24. Swap Nodes in Pairs][title]

## Description

Given a linked list, swap every two adjacent nodes and return its head.

**Example:**

```
Given 1->2->3->4, you should return the list as 2->1->4->3.
```

**Note:**

- Your algorithm should use only constant extra space.
- You may **not** modify the values in the list's nodes, only nodes itself may be changed.

**Tags:** Linked List

## 题意
>题意是让你交换链表中相邻的两个节点，最终返回交换后链表的头

## 题解

### 思路1
> 我们可以用递归来算出子集合的结果，递归的终点就是指针指到链表末少于两个元素时，如果不是终点，那么我们就对其两节点进行交换，这里我们需要一个临时节点来作为交换桥梁，就不多说了。

```go
//	递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head.Next
	head.Next = swapPairs(node.Next)
	node.Next = head
	return node
}
```

### 思路2
> 另一种实现方式就是用循环来实现了，两两交换节点，也需要一个临时节点来作为交换桥梁，直到当前指针指到链表末少于两个元素时停止，代码很简单，如下所示。
```go
// 循环
func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{Val: 0, Next: nil}
	cur := preHead
	preHead.Next = head
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next.Next
		cur.Next.Next = tmp.Next
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = cur.Next.Next
	}

	return preHead.Next
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/swap-nodes-in-pairs/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
