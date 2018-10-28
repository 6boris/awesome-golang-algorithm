# [703. Kth Largest Element in a Stream][title]

## Description

Given two binary strings, return their sum (also a binary string).

The input strings are both **non-empty** and contains only characters `1` or `0`.

**Example 1:**

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2:**

```
Input: a = "1010", b = "1011"
Output: "10101"
```

**Tags:** Math, String

## 题意
>数据流中的第K大元素

## 题解

### 思路1
> 维护一个有K个数的数组每次操作的时候都会去修改这个数组
-  因为各种原因导致我这个代码写的比较复杂不太友好

```go
type KthLargest1 struct {
	heap []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	//	初始化结构体
	kthLargest := KthLargest{heap: make([]int, k)}
	kthLargest.k = k
	//	需要将原head 总默认值设置最小
	for i, _ := range kthLargest.heap {
		kthLargest.heap[i] = math.MinInt32
	}

	//	初始化数组
	for i := 0; i < k && i < len(nums); i++ {
		kthLargest.heap[i] = nums[i]
	}

	//	对数组进行排序
	kthLargest.heap = InsertSort(kthLargest.heap)

	//	当第一次初始化的数组长度比K大的时候处理后面的数
	for i := k; i < len(nums); i++ {
		kthLargest.Add(nums[i])
	}


	return kthLargest
}

//	想数据流中添加一个数并返回第K大的数
func (this *KthLargest) Add(val int) int {
	//	如果添加进来的数比末尾的打就直接替换掉末尾的数
	if val > this.heap[this.k-1] {
		this.heap[this.k-1] = val
	}
	//	替换完后需要重新排序
	this.heap = InsertSort(this.heap)
	return this.heap[this.k-1]
}

// 插入排序(Insert Sort)
func InsertSort(arr []int) []int {
	var i, j, tmp int
	for i = 1; i < len(arr); i++ {
		tmp = arr[i]
		for j = i; j > 0 && arr[j-1] < tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}
```

### 思路2
> 使用最小堆
- 自己比较懒直接找到别人的，用例GoLang自带的最小堆，实现的基本的接口

```go

// copied from golang doc
// mininum setup of integer min heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// real thing starts here
type KthLargest struct {
	Nums *IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	// push all elements to the heap
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	// remove the smaller elements from the heap such that only the k largest elements are in the heap
	for len(*h) > k {
		heap.Pop(h)
	}
	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else if val > (*this.Nums)[0] {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}
	return (*this.Nums)[0]
}

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
