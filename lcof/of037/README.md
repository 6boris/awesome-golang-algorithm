---
description: 剑指 Offer 37. 序列化二叉树
---

# OF37.序列化二叉树

## 题目描述

[题目地址](https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof/)

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。



### **示例 1：**

```go
你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
```

## 题解

### 思路1 ： 

**算法流程：**

1. 
**复杂度分析：**

* **时间复杂度**$$O(N)$$**：**\)
* **空间复杂度**$$O(N)$$**：** 

#### 代码

{% tabs %}
{% tab title="Go" %}
```go
func findRepeatNumber2(nums []int) int {
    mapTmp := make(map[int]bool)
    for _, val := range nums {
        if mapTmp[val] {
            return val
        } else {
            mapTmp[val] = true
        }
    }
    return 0
}
```
{% endtab %}
{% endtabs %}

## 总结

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 算法 题解：[awesome-golang-algorithm](https://github.com/kylesliu/awesome-golang-algorithm)

