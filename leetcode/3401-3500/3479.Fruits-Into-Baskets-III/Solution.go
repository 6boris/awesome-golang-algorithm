package Solution

import "math"

const (
	INT_MIN = math.MinInt32
)

type SegTree struct {
	segNode []int
	baskets []int
}

func (this *SegTree) build(p, l, r int) {
	if l == r {
		this.segNode[p] = this.baskets[l]
		return
	}
	mid := (l + r) >> 1
	this.build(p<<1, l, mid)
	this.build(p<<1|1, mid+1, r)
	this.segNode[p] = max(this.segNode[p<<1], this.segNode[p<<1|1])
}

func (this *SegTree) query(p, l, r, ql, qr int) int {
	if ql > r || qr < l {
		return INT_MIN
	}
	if ql <= l && r <= qr {
		return this.segNode[p]
	}
	mid := (l + r) >> 1
	return max(this.query(p<<1, l, mid, ql, qr),
		this.query(p<<1|1, mid+1, r, ql, qr))
}

func (this *SegTree) update(p, l, r, pos, val int) {
	if l == r {
		this.segNode[p] = val
		return
	}
	mid := (l + r) >> 1
	if pos <= mid {
		this.update(p<<1, l, mid, pos, val)
	} else {
		this.update(p<<1|1, mid+1, r, pos, val)
	}
	this.segNode[p] = max(this.segNode[p<<1], this.segNode[p<<1|1])
}

func Solution(fruits []int, baskets []int) int {
	m := len(baskets)
	if m == 0 {
		return len(fruits)
	}

	tree := SegTree{
		segNode: make([]int, 4*m+7),
		baskets: baskets,
	}
	tree.build(1, 0, m-1)

	count := 0
	for i := 0; i < len(fruits); i++ {
		l, r, res := 0, m-1, -1
		for l <= r {
			mid := (l + r) >> 1
			if tree.query(1, 0, m-1, 0, mid) >= fruits[i] {
				res = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if res != -1 && tree.baskets[res] >= fruits[i] {
			tree.update(1, 0, m-1, res, INT_MIN)
		} else {
			count++
		}
	}

	return count
}
