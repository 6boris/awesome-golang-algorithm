package Solution

import "container/heap"

func Solution(nums []int, k int) int64 {
	n := len(nums)
	var ans int64 = 0
	var operations int64 = 0
	// Maintain the window maximum with a priority queue
	h := &MaxHeap{}
	heap.Init(h)
	// Record the set of elements covered by this index
	cnt := make([][]int, n)
	for i := range cnt {
		cnt[i] = []int{}
	}
	// Index of the first greater-or-equal element to the right
	pos := make([]int, n)
	for i := range pos {
		pos[i] = n
	}
	// Decreasing monotonic stack
	st := []int{}

	left := 0
	for right := 0; right < n; right++ {
		// Maintain the monotonic stack
		for len(st) > 0 && nums[right] >= nums[st[len(st)-1]] {
			pos[st[len(st)-1]] = right
			st = st[:len(st)-1]
		}
		// The stack top is the nearest greater element to the left of nums[right]
		if len(st) > 0 {
			cnt[st[len(st)-1]] = append(cnt[st[len(st)-1]], right)
		}
		st = append(st, right)

		// The priority queue front is the window maximum
		heap.Push(h, [2]int{nums[right], right})
		// Adding nums[right] increases the operation count
		if diff := (*h)[0][0] - nums[right]; diff > 0 {
			operations += int64(diff)
		}

		// The minimum operation count exceeds k; shrink the window
		for operations > int64(k) {
			// Removing nums[left] from the window decreases the operation count
			for _, i := range cnt[left] {
				// Subtract the contribution
				operations -= int64(nums[left]-nums[i]) * int64(min(pos[i], right+1)-i)
			}
			left++
		}

		// Remove elements with indices less than left from the priority queue
		for h.Len() > 0 && (*h)[0][1] < left {
			heap.Pop(h)
		}

		// There are right - left + 1 valid contiguous subarrays ending at right
		ans += int64(right - left + 1)
	}

	return ans
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
