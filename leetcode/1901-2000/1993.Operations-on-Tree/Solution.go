package Solution

type LockingTree struct {
	tree     []int
	locker   []int
	children map[int][]int
}

func Constructor(parent []int) LockingTree {
	children := make(map[int][]int)
	for c, p := range parent {
		if _, ok := children[p]; !ok {
			children[p] = make([]int, 0)
		}
		children[p] = append(children[p], c)
	}
	return LockingTree{tree: parent, locker: make([]int, len(parent)), children: children}
}

// u, u, ul, l, u
func (this *LockingTree) Lock(num int, user int) bool {
	if this.locker[num] != 0 {
		return false
	}
	this.locker[num] = user
	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.locker[num] == 0 || this.locker[num] != user {
		return false
	}
	this.locker[num] = 0
	return true
}

func (this *LockingTree) ancesors(num int) bool {
	for {
		if num == -1 {
			return false
		}
		if this.locker[num] != 0 {
			return true
		}
		num = this.tree[num]
	}
}

func (this *LockingTree) unlockDescendant(num int) bool {
	q := []int{num}
	found := false
	for len(q) > 0 {
		nq := make([]int, 0)
		for _, i := range q {
			if this.locker[i] != 0 {
				this.locker[i] = 0
				found = true
			}
			for _, c := range this.children[i] {
				nq = append(nq, c)
			}
		}
		q = nq
	}
	return found
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	if this.locker[num] != 0 {
		return false
	}
	if this.ancesors(num) {
		return false
	}
	if !this.unlockDescendant(num) {
		return false
	}
	this.locker[num] = user
	return true
}

type op struct {
	name string
	a, b int
}

func Solution(p []int, opts []op) []bool {
	c := Constructor(p)
	ans := make([]bool, 0)
	for _, o := range opts {
		if o.name == "l" {
			ans = append(ans, c.Lock(o.a, o.b))
			continue
		}
		if o.name == "u" {
			ans = append(ans, c.Unlock(o.a, o.b))
			continue
		}
		ans = append(ans, c.Upgrade(o.a, o.b))
	}
	return ans
}
