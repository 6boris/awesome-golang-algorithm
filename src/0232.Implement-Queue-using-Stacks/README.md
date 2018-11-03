# [232. Implement Queue using Stacks][title]

## Description

Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.

pop() -- Removes the element from in front of queue.

peek() -- Get the front element.

empty() -- Return whether the queue is empty.

**Example:**

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
### 思路2
1. 首选封装一个栈的结构

2. 定义队列包含两个栈，分别是输入栈`s1`和输出栈`s2`,入对数据都存入`s1`中，出对数据都从`s2`来。

> 具体实现如下：
```go
// 封装栈结构
type Stack []int
func (s *Stack) Empty() bool {return len(*s) == 0}
func (s *Stack) Push(x int) {*s = append(*s, x)}
func (s *Stack) Peek() int {return (*s)[len(*s)-1]}
func (s *Stack) Pop() int {
    x := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return x
}

type MyQueue struct {
    s1, s2 Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int)  {
    q.s1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
    if q.s2.Empty() {
        for !q.s1.Empty() {
            x := q.s1.Pop()
            q.s2.Push(x)
        }
    }
    return q.s2.Pop()
}


/** Get the front element. */
func (q *MyQueue) Peek() int {
    if q.s2.Empty() {
        for !q.s1.Empty() {
            x := q.s1.Pop()
            q.s2.Push(x)
        }
    }
    return q.s2.Peek()
}


/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
    return q.s1.Empty() && q.s2.Empty()
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-letcode][me]

[title]:https://leetcode.com/problems/implement-queue-using-stacks/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
