package Solution

import "container/heap"

//	根据golang只带的container类型自定义一个堆
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

//	基本结构体
type KthLargest struct {
	Nums *IntHeap
	K    int
}

//	根据所给的数据初始化heap
func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	// 将所有元素推送到heap
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	// 弹出不必要的小元素，只要留K个最大的
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
	h.
	return (*this.Nums)[0]
}
