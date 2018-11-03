# [242. Implement Queue using Stacks][title]

## Description

Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.

pop() -- Removes the element from in front of queue.

peek() -- Get the front element.

empty() -- Return whether the queue is empty.

Example:

```
MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
```

## NOTES
- You must use only standard operations of a stack -- which means only `push to top`, `peek/pop from top`, `size`, and is `empty` operations are valid.
- Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
- You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

**Tags:** Queue, Stack

## 题意
> 使用栈来实现队列

主要体现为实现四个函数
- push(x) 在队列的末尾增加元素
- pop() 在队列的头部删除元素
- peek() 获取队列的第一个元素
- empty() 返回队列是否是空队列

### 思路1
定义结构体`MyQueue`, 包含一个切片`array`用来存储队列元素，主要逻辑在于定义中间元素来移动切片，从而达到去栈底元素的目的。
```go
type MyQueue struct {
	array []int
}
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-letcode][me]

[title]: https://leetcode.com/problems/valid-anagram/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
