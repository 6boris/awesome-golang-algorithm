package Solution

type LazyTag struct {
	toAdd int
}

func (l *LazyTag) Add(other *LazyTag) *LazyTag {
	l.toAdd += other.toAdd
	return l
}

func (l *LazyTag) HasTag() bool {
	return l.toAdd != 0
}

func (l *LazyTag) Clear() {
	l.toAdd = 0
}

type SegmentTreeNode struct {
	minValue int
	maxValue int
	lazyTag  *LazyTag
}

func NewSegmentTreeNode() *SegmentTreeNode {
	return &SegmentTreeNode{
		minValue: 0,
		maxValue: 0,
		lazyTag:  &LazyTag{},
	}
}

type SegmentTree struct {
	n    int
	tree []*SegmentTreeNode
}

func NewSegmentTree(data []int) *SegmentTree {
	n := len(data)
	tree := make([]*SegmentTreeNode, n*4+1)
	for i := range tree {
		tree[i] = NewSegmentTreeNode()
	}
	seg := &SegmentTree{n: n, tree: tree}
	seg.build(data, 1, n, 1)
	return seg
}

func (seg *SegmentTree) Add(l, r, val int) {
	tag := &LazyTag{toAdd: val}
	seg.update(l, r, tag, 1, seg.n, 1)
}

func (seg *SegmentTree) FindLast(start, val int) int {
	if start > seg.n {
		return -1
	}
	return seg.find(start, seg.n, val, 1, seg.n, 1)
}

func (seg *SegmentTree) applyTag(i int, tag *LazyTag) {
	seg.tree[i].minValue += tag.toAdd
	seg.tree[i].maxValue += tag.toAdd
	seg.tree[i].lazyTag.Add(tag)
}

func (seg *SegmentTree) pushdown(i int) {
	if seg.tree[i].lazyTag.HasTag() {
		tag := &LazyTag{toAdd: seg.tree[i].lazyTag.toAdd}
		seg.applyTag(i<<1, tag)
		seg.applyTag((i<<1)|1, tag)
		seg.tree[i].lazyTag.Clear()
	}
}

func (seg *SegmentTree) pushup(i int) {
	left := seg.tree[i<<1]
	right := seg.tree[(i<<1)|1]
	seg.tree[i].minValue = min(left.minValue, right.minValue)
	seg.tree[i].maxValue = max(left.maxValue, right.maxValue)
}

func (seg *SegmentTree) build(data []int, l, r, i int) {
	if l == r {
		seg.tree[i].minValue = data[l-1]
		seg.tree[i].maxValue = data[l-1]
		return
	}

	mid := l + ((r - l) >> 1)
	seg.build(data, l, mid, i<<1)
	seg.build(data, mid+1, r, (i<<1)|1)
	seg.pushup(i)
}

func (seg *SegmentTree) update(targetL, targetR int, tag *LazyTag, l, r, i int) {
	if targetL <= l && r <= targetR {
		seg.applyTag(i, tag)
		return
	}

	seg.pushdown(i)
	mid := l + ((r - l) >> 1)
	if targetL <= mid {
		seg.update(targetL, targetR, tag, l, mid, i<<1)
	}
	if targetR > mid {
		seg.update(targetL, targetR, tag, mid+1, r, (i<<1)|1)
	}
	seg.pushup(i)
}

func (seg *SegmentTree) find(targetL, targetR, val, l, r, i int) int {
	if seg.tree[i].minValue > val || seg.tree[i].maxValue < val {
		return -1
	}

	if l == r {
		return l
	}

	seg.pushdown(i)
	mid := l + ((r - l) >> 1)

	if targetR >= mid+1 {
		res := seg.find(targetL, targetR, val, mid+1, r, (i<<1)|1)
		if res != -1 {
			return res
		}
	}

	if l <= targetR && mid >= targetL {
		return seg.find(targetL, targetR, val, l, mid, i<<1)
	}

	return -1
}

func Solution(nums []int) int {
	occurrences := make(map[int][]int)

	sgn := func(x int) int {
		if x%2 == 0 {
			return 1
		}
		return -1
	}

	length := 0
	prefixSum := make([]int, len(nums))
	prefixSum[0] = sgn(nums[0])
	occurrences[nums[0]] = append(occurrences[nums[0]], 1)

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1]
		occ := occurrences[nums[i]]
		if len(occ) == 0 {
			prefixSum[i] += sgn(nums[i])
		}
		occurrences[nums[i]] = append(occ, i+1)
	}

	seg := NewSegmentTree(prefixSum)
	for i := 0; i < len(nums); i++ {
		length = max(length, seg.FindLast(i+length, 0)-i)
		nextPos := len(nums) + 1
		occurrences[nums[i]] = occurrences[nums[i]][1:]
		if len(occurrences[nums[i]]) > 0 {
			nextPos = occurrences[nums[i]][0]
		}

		seg.Add(i+1, nextPos-1, -sgn(nums[i]))
	}

	return length
}
