package Solution

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

func Solution(nums []int, k int, x int) []int64 {
	helper := NewHelper(x)
	ans := []int64{}

	for i := 0; i < len(nums); i++ {
		helper.Insert(nums[i])
		if i >= k {
			helper.Remove(nums[i-k])
		}
		if i >= k-1 {
			ans = append(ans, helper.Get())
		}
	}

	return ans
}

type Helper struct {
	x      int
	result int64
	large  *redblacktree.Tree[pair, struct{}]
	small  *redblacktree.Tree[pair, struct{}]
	occ    map[int]int
}

type pair struct {
	freq int
	num  int
}

func pairComparator(a, b pair) int {
	if a.freq != b.freq {
		return a.freq - b.freq
	}
	return a.num - b.num
}

func NewHelper(x int) *Helper {
	return &Helper{
		x:      x,
		result: 0,
		large:  redblacktree.NewWith[pair, struct{}](pairComparator),
		small:  redblacktree.NewWith[pair, struct{}](pairComparator),
		occ:    make(map[int]int),
	}
}

func (h *Helper) Insert(num int) {
	if h.occ[num] > 0 {
		h.internalRemove(pair{freq: h.occ[num], num: num})
	}
	h.occ[num]++
	h.internalInsert(pair{freq: h.occ[num], num: num})
}

func (h *Helper) Remove(num int) {
	h.internalRemove(pair{freq: h.occ[num], num: num})
	h.occ[num]--
	if h.occ[num] > 0 {
		h.internalInsert(pair{freq: h.occ[num], num: num})
	}
}

func (h *Helper) Get() int64 {
	return h.result
}

func (h *Helper) internalInsert(p pair) {
	if h.large.Size() < h.x {
		h.result += int64(p.freq) * int64(p.num)
		h.large.Put(p, struct{}{})
	} else {
		minLarge := h.large.Left().Key
		if pairComparator(p, minLarge) > 0 {
			h.result += int64(p.freq) * int64(p.num)
			h.large.Put(p, struct{}{})
			toRemove := h.large.Left().Key
			h.result -= int64(toRemove.freq) * int64(toRemove.num)
			h.large.Remove(toRemove)
			h.small.Put(toRemove, struct{}{})
		} else {
			h.small.Put(p, struct{}{})
		}
	}
}

func (h *Helper) internalRemove(p pair) {
	if _, found := h.large.Get(p); found {
		h.result -= int64(p.freq) * int64(p.num)
		h.large.Remove(p)

		if h.small.Size() > 0 {
			maxSmall := h.small.Right().Key
			h.result += int64(maxSmall.freq) * int64(maxSmall.num)
			h.small.Remove(maxSmall)
			h.large.Put(maxSmall, struct{}{})
		}
	} else if _, found := h.small.Get(p); found {
		h.small.Remove(p)
	}
}
