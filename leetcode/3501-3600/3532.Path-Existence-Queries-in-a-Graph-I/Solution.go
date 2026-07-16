package Solution

type unionFind3532 struct {
	father []int
}

func (u *unionFind3532) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}

func (u *unionFind3532) union(x, y int) {
	fx, fy := u.find(x), u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func Solution(n int, nums []int, maxDiff int, queries [][]int) []bool {
	u := &unionFind3532{
		father: make([]int, n),
	}
	for i := range n {
		u.father[i] = i
	}

	size := len(nums)

	for i := 1; i < size; i++ {
		diff := nums[i] - nums[i-1]
		if diff <= maxDiff {
			u.union(i, i-1)
		}
	}
	ret := make([]bool, len(queries))
	for i, q := range queries {
		a, b := u.find(q[0]), u.find(q[1])
		ret[i] = a == b
	}
	return ret
}
