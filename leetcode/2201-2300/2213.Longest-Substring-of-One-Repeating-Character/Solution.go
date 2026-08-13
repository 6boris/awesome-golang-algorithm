package Solution

// 实际可用的标准线段树实现结构：
type SegTree2213 struct {
	lMax, rMax, maxLen []int
	lChar, rChar       []byte
	s                  []byte
}

func NewSegTree2213(s string) *SegTree2213 {
	n := len(s)
	st := &SegTree2213{
		lMax:   make([]int, 4*n),
		rMax:   make([]int, 4*n),
		maxLen: make([]int, 4*n),
		lChar:  make([]byte, 4*n),
		rChar:  make([]byte, 4*n),
		s:      []byte(s),
	}
	st.build(1, 0, n-1)
	return st
}

func (st *SegTree2213) pushUp(u, l, r int) {
	left, right := u<<1, u<<1|1
	mid := (l + r) >> 1

	st.lChar[u] = st.lChar[left]
	st.rChar[u] = st.rChar[right]

	st.lMax[u] = st.lMax[left]
	st.rMax[u] = st.rMax[right]
	st.maxLen[u] = max(st.maxLen[left], st.maxLen[right])

	// 左右子树拼合处字符相同时
	if st.rChar[left] == st.lChar[right] {
		cross := st.rMax[left] + st.lMax[right]
		st.maxLen[u] = max(st.maxLen[u], cross)

		// 如果左半部分全部相同
		if st.lMax[left] == mid-l+1 {
			st.lMax[u] = st.lMax[left] + st.lMax[right]
		}
		// 如果右半部分全部相同
		if st.rMax[right] == r-mid {
			st.rMax[u] = st.rMax[right] + st.rMax[left]
		}
	}
}

func (st *SegTree2213) build(u, l, r int) {
	if l == r {
		st.lMax[u], st.rMax[u], st.maxLen[u] = 1, 1, 1
		st.lChar[u] = st.s[l]
		st.rChar[u] = st.s[l]
		return
	}
	mid := (l + r) >> 1
	st.build(u<<1, l, mid)
	st.build(u<<1|1, mid+1, r)
	st.pushUp(u, l, r)
}

func (st *SegTree2213) update(u, l, r int, idx int, val byte) {
	if l == r {
		st.s[idx] = val
		st.lChar[u] = val
		st.rChar[u] = val
		return
	}
	mid := (l + r) >> 1
	if idx <= mid {
		st.update(u<<1, l, mid, idx, val)
	} else {
		st.update(u<<1|1, mid+1, r, idx, val)
	}
	st.pushUp(u, l, r)
}

func Solution(s string, queryCharacters string, queryIndices []int) []int {
	n := len(s)
	st := NewSegTree2213(s)
	k := len(queryIndices)
	ans := make([]int, k)

	for i := 0; i < k; i++ {
		st.update(1, 0, n-1, queryIndices[i], queryCharacters[i])
		ans[i] = st.maxLen[1]
	}
	return ans
}
