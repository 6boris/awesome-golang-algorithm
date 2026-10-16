package Solution

const MOD = 1000000007

type SegmentTree struct {
	countTree []int64
	sqTree    []int64
	lazy      []int64
}

func NewSegmentTree(n int) *SegmentTree {
	return &SegmentTree{
		countTree: make([]int64, 4*n),
		sqTree:    make([]int64, 4*n),
		lazy:      make([]int64, 4*n),
	}
}

func (st *SegmentTree) pushDown(node, l, r int) {
	if st.lazy[node] != 0 {
		x := st.lazy[node]
		mid := (l + r) / 2

		// 左子树更新
		leftLen := int64(mid - l + 1)
		st.sqTree[2*node] = (st.sqTree[2*node] + 2*x*st.countTree[2*node] + x*x*leftLen) % MOD
		st.countTree[2*node] = (st.countTree[2*node] + x*leftLen) % MOD
		st.lazy[2*node] += x

		// 右子树更新
		rightLen := int64(r - mid)
		st.sqTree[2*node+1] = (st.sqTree[2*node+1] + 2*x*st.countTree[2*node+1] + x*x*rightLen) % MOD
		st.countTree[2*node+1] = (st.countTree[2*node+1] + x*rightLen) % MOD
		st.lazy[2*node+1] += x

		st.lazy[node] = 0
	}
}

func (st *SegmentTree) update(node, l, r, ql, qr int, val int64) {
	if ql <= l && r <= qr {
		length := int64(r - l + 1)
		st.sqTree[node] = (st.sqTree[node] + 2*val*st.countTree[node] + val*val*length) % MOD
		st.countTree[node] = (st.countTree[node] + val*length) % MOD
		st.lazy[node] += val
		return
	}

	st.pushDown(node, l, r)
	mid := (l + r) / 2
	if ql <= mid {
		st.update(2*node, l, mid, ql, qr, val)
	}
	if qr > mid {
		st.update(2*node+1, mid+1, r, ql, qr, val)
	}

	st.countTree[node] = (st.countTree[2*node] + st.countTree[2*node+1]) % MOD
	st.sqTree[node] = (st.sqTree[2*node] + st.sqTree[2*node+1]) % MOD
}

func Solution(nums []int) int {
	n := len(nums)
	st := NewSegmentTree(n)
	lastPos := make(map[int]int)
	var ans int64 = 0

	for i, x := range nums {
		prev := -1
		if p, ok := lastPos[x]; ok {
			prev = p
		}
		// 将区间 [prev + 1, i] 内的所有起点对应子数组的 distinct count +1
		st.update(1, 0, n-1, prev+1, i, 1)
		lastPos[x] = i

		// 根节点维护的就是以当前 i 结尾的所有子数组的平方和总和
		ans = (ans + st.sqTree[1]) % MOD
	}

	return int(ans)
}
