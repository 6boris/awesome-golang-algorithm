package Solution

import "fmt"

type segNode struct {
	l, r, m     int
	left, right *segNode
}

// Just for debug
func (s *segNode) Floor() {
	q := []*segNode{s}
	for len(q) > 0 {
		next := make([]*segNode, 0)
		for _, i := range q {
			fmt.Printf("[%d, %d], max: %d ***$", i.l, i.r, i.m)
			if i.left != nil {
				next = append(next, i.left)
			}
			if i.right != nil {
				next = append(next, i.right)
			}
		}
		fmt.Println()
		q = next
	}
}

// 1, 2, 3, 4,/ 5, 6, 7, 8
func buildSegTree(left, right int, nums []int) *segNode {
	if left > right {
		return nil
	}
	node := &segNode{l: left, r: right}
	if left == right {
		node.m = nums[left]
		return node
	}
	mid := (right-left)/2 + left
	node.left = buildSegTree(left, mid, nums)
	node.right = buildSegTree(mid+1, right, nums)
	node.m = node.left.m
	if node.right.m > node.m {
		node.m = node.right.m
	}
	return node
}
func searchSeg(root *segNode, s, e int, nums []int) int {
	if root.l == s && root.r == e {
		return root.m
	}
	if s == e {
		return nums[s]
	}
	mid := (root.r-root.l)/2 + root.l

	if e < mid {
		return searchSeg(root.left, s, e, nums)
	}
	if s > mid {
		return searchSeg(root.right, s, e, nums)
	}

	if s < mid && e > mid {
		a := searchSeg(root.left, s, mid, nums)
		b := searchSeg(root.right, mid+1, e, nums)
		if a > b {
			return a
		}
		return b
	}

	if e == mid {
		a := searchSeg(root.left, e, e, nums)
		b := searchSeg(root.left, s, mid-1, nums)
		if a > b {
			return a
		}
		return b
	}
	a := searchSeg(root.right, s, s, nums)
	b := searchSeg(root.right, mid+1, e, nums)
	if a > b {
		return a
	}
	return b
}

func Solution(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	l := len(nums)
	ans := make([]int, l-k+1)
	segRoot := buildSegTree(0, l-1, nums)
	for i := 0; i <= l-k; i++ {
		e := i + k - 1
		ans[i] = searchSeg(segRoot, i, e, nums)
	}
	return ans
}
