package Solution

import "sort"

type segNode2070 struct {
	l, r, m     int
	left, right *segNode2070
}

func buildSegTree2070(left, right int, nums []int, m map[int]int) *segNode2070 {
	if left == right {
		return &segNode2070{
			l: left, r: right, m: m[nums[left]],
		}
	}

	mid := (left + right) / 2
	leftChild := buildSegTree2070(left, mid, nums, m)
	rightChild := buildSegTree2070(mid+1, right, nums, m)

	node := &segNode2070{
		l: left, r: right, m: max(leftChild.m, rightChild.m),
		left: leftChild, right: rightChild,
	}
	return node
}
func queryMax2070(node *segNode2070, left, right int) int {
	if node.l == left && node.r == right {
		return node.m
	}

	mid := (node.l + node.r) / 2

	if right <= mid {
		return queryMax2070(node.left, left, right)
	} else if left > mid {
		return queryMax2070(node.right, left, right)
	} else {
		leftMax := queryMax2070(node.left, left, mid)
		rightMax := queryMax2070(node.right, mid+1, right)
		return max(leftMax, rightMax)
	}
}

func Solution(items [][]int, queries []int) []int {
	ans := make([]int, len(queries))
	p2b := make(map[int]int)
	keys := make([]int, 0)
	for _, item := range items {
		p, b := item[0], item[1]
		v, ok := p2b[p]
		if !ok {
			keys = append(keys, p)
		}
		if v < b {
			p2b[p] = b
		}
	}
	l := len(keys)
	sort.Ints(keys)
	tree := buildSegTree2070(0, l-1, keys, p2b)
	for i, q := range queries {
		idx := sort.Search(l, func(i int) bool {
			return keys[i] > q
		})
		if idx == 0 {
			ans[i] = 0
			continue
		}
		ans[i] = queryMax2070(tree, 0, idx-1)
	}
	return ans
}
