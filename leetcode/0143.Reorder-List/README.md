# [143. Reorder List][title]

## Description
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

**Example 1:**
```
Given 1->2->3->4, reorder it to 1->4->2->3.
```

**Example 2:**
```
Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
```

**Tags:** Linked List

## 题意
> 给定一个单链表，对链表进行重新排序。

## 题解
> 把链表分成两部分，对后半段链表进行逆序，然后再将两部分合并。

### 思路1
1. 找到链表的中点，将链表拆分成前后两部分。
1. 对后半段链表进行逆序操作。
1. 将前半段链表和逆序后的后半段链表合并。

```go
func reorderList(head *ListNode)  {
    if nil == head || nil == head.Next {
        return
    }
    mid := findMid(head)
    first, second := head, reverseList(mid.Next)
    mid.Next = nil
    
    newHead := new(ListNode)
    current := newHead
    for first != nil || second != nil {
        if first != nil {
            current.Next = &ListNode{first.Val, nil}
            current = current.Next
            first = first.Next
        }
        if second != nil {
            current.Next = &ListNode{second.Val, nil}
            current = current.Next
            second = second.Next
        }
    }
    head.Next = newHead.Next.Next
}

func findMid(head *ListNode) *ListNode {
    slow, fast := head, head.Next
    for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func reverseList(head *ListNode) *ListNode {
    newHead := head
    for head.Next != nil {
        next := head.Next
        head.Next = next.Next
        next.Next = newHead
        newHead = next
    }
    return newHead
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/reorder-list/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
