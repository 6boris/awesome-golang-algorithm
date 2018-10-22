# [1. Reverse Linked List][title]

## Description

Reverse a singly linked list.

**Example 1:**

```
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
```

**Example 2:**

```
```

**Follow up:**
A linked list can be reversed either iteratively or recursively. Could you implement both?


**Tags:** Math, String


## 题意
>翻转链表

## 题解

### 思路1
> 正常循环翻转

```go
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
```

### 思路2
> 递归翻转
```go
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = p
	return p
}
```

## 测试
```go
go test -run=Solution_test.go -bench=BenchmarkSolution1 -benchmem -cpuprofile=bench-1.prof
go test -run=Solution_test.go -bench=BenchmarkSolution2 -benchmem -cpuprofile=bench-2.prof

go tool pprof bench-1.prof
web
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/reverse-linked-list/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
