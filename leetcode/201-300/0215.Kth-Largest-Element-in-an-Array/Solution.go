package Solution

import (
	"container/heap"
)

type iheap []int

// 实现 Sort 接口的Len, Less, Swap方法
func (h iheap) Len() int           { return len(h) }
func (h iheap) Less(i, j int) bool { return h[i] < h[j] }
func (h iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 看一下堆顶元素，不调整
func (h iheap) Peek() int {
	return h[0]
}

// 实现 container/Heap 接口的Pop,Push方法，receiver需要是指针
func (h *iheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 弹出数组最后一个元素，同时数组长度--，堆内部通过替换堆顶元素来进行 heapifyDown 调整
func (h *iheap) Pop() interface{} {
	n := (*h).Len()
	x := (*h)[n-1]  // 数组最后一个元素
	*h = (*h)[:n-1] // size--
	return x
}

func findKthLargest_1(nums []int, k int) int {
	if k <= 0 || len(nums) == 0 {
		return 0
	}

	h := iheap{}
	heap.Init(&h)
	// 先把前k个元素构建小顶堆
	for i := 0; i < k; i++ {
		heap.Push(&h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		// 看一下堆顶元素，如果小于等于当前nums[i]就没必要替换了，减少调整次数
		top := h.Peek()
		if nums[i] > top {
			heap.Pop(&h) // 弹出堆顶
			heap.Push(&h, nums[i])
		}
	}
	// 此时堆顶元素就是第K大
	return h.Peek()
}

func findKthLargest_2(nums []int, k int) int {
	if k <= 0 || len(nums) == 0 {
		return 0
	}

	// 分区操作，返回轴点索引下标
	// 在数组nums的子区间 [left, right] 执行 partition 操作，返回 nums[left]（选取的第一个作为pivot） 排序以后应该在的位置
	partition := func(nums []int, left, right int) int {
		pivot := nums[left]
		j := left
		for i := left + 1; i <= right; i++ {
			// 小于 pivot 的元素都被交换到前面
			if nums[i] < pivot {
				j++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
		// 在之前遍历的过程中，满足 [left+1, j] < pivot，并且 (j, i] >= pivot
		nums[j], nums[left] = nums[left], nums[j]
		// 交换以后 [left, j-1] < pivot, nums[j] = pivot, [j+1, right] >= pivot
		return j
	}

	n := len(nums)
	// 转换一下，第 k 大元素的索引是 n - k
	left, right, target := 0, n-1, n-k
	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}
