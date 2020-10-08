# [147. Insertion Sort List][title]

## Description

Sort a linked list using insertion sort.

### Algorithm of Insertion Sort:
1. Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
1. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
1. It repeats until no input elements remain.

**Example 1:**
```
Input: 4->2->1->3
Output: 1->2->3->4
```

**Example 2:**
```
Input: -1->5->3->4->0
Output: -1->0->3->4->5
```

**Tags:** Linked List

## 题意
> 对一个单链表进行插入排序。

## 题解

### 思路 1
> 构造一个新的链表，每次将原链表中一个节点插入到新链表中合适的位置

```go
func insertionSortList(head *ListNode) *ListNode {
    if nil == head {
        return nil
    }
    newHead := new(ListNode)
    newHead.Next = &ListNode{head.Val, nil}
    
    head = head.Next
    for head != nil {
        prev, current := newHead, newHead.Next
        for current != nil {
            if head.Val < current.Val {
                prev.Next = &ListNode{head.Val, current}
                break
            }
            prev, current = prev.Next, current.Next
        }
        if nil == current && head.Val >= prev.Val {
            prev.Next = &ListNode{head.Val, nil}
        }
        head = head.Next
    }
    return newHead.Next
}


Sort a linked list using insertion sort.

![Sort](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)

**Tags:** Math, String

## 题意
> 使用插入排序对一个链表排序。

## 题解

### 思路1
> 。。。。

```go

```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/insertion-sort-list/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
