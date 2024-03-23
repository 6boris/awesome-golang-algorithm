package Solution

type unionFind2709 struct {
	father []int
}

func (u *unionFind2709) find(x int) int {
	if u.father[x] != x {
		u.father[x] = u.find(u.father[x])
	}
	return u.father[x]
}
func (u *unionFind2709) union(x, y int) {
	fx := u.find(x)
	fy := u.find(y)
	if fx < fy {
		u.father[fy] = fx
	} else {
		u.father[fx] = fy
	}
}

func Solution(nums []int) bool {
	l := len(nums)
	if l == 1 {
		return true
	}
	alloc := 100000
	primeFactor := make([]int, alloc+1)
	for i := 2; i <= alloc; i++ {
		if primeFactor[i] == 0 {
			for next := i; next <= alloc; next += i {
				primeFactor[next] = i
			}
		}
	}

	u := unionFind2709{father: make([]int, 2*alloc+1)}
	for i := 0; i < 2*alloc+1; i++ {
		u.father[i] = i
	}
	for _, n := range nums {
		if n == 1 {
			return false
		}
		cur := n
		for cur > 1 {
			fac := primeFactor[cur]
			root := fac + alloc
			u.union(root, n)
			for cur%fac == 0 {
				cur /= fac
			}
		}
	}
	root := u.find(nums[0])
	for i := 1; i < l; i++ {
		if u.find(nums[i]) != root {
			return false
		}
	}
	return true
}
