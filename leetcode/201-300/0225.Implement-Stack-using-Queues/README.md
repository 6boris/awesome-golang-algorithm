# [0225.Implement Stack using Queues][title]

## Description

Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.

pop() -- Removes the element on top of the stack.

top() -- Get the top element.

empty() -- Return whether the stack is empty.

**Example:**

```
MyStack stack = new MyStack();

stack.push(1);
stack.push(2);
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false
```

## NOTES
- You must use only standard operations of a queue -- which means only `push to back`, `peek/pop from front`, `size`, and `is empty` operations are valid.
- Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
- You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).


**Tags:** Queue, Stack

## 题意
>只使用队列的标准操作来实现栈，这些操作仅仅包括`push to back`, `peek/pop from front`, `size`, `is empty`

主要体现为实现四个函数
- push(x) 在栈的末尾增加元素
- pop() 在栈顶删除元素
- peek() 查看栈顶元素
- empty() 查看是否是空栈

### 思路1
定义结构体`MyStack`,包含两个元素，分别是入队列`enqueue`和出队列`dequeue`,所有入栈数据都先存入`enquequ`, 所有的出栈数据都来自`dequeue`
```go
type MyStack struct {
  enqueue []int
  dequeue []int
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-letcode][me]

[title]: https://leetcode.com/problems/implement-stack-using-queues/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
