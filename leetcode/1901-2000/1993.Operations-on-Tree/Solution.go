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
func (t *LockingTree) Lock(num, user int) bool {
	if t.locker[num] != 0 {
		return false
	}
	t.locker[num] = user
	return true
}

func (t *LockingTree) Unlock(num, user int) bool {
	if t.locker[num] == 0 || t.locker[num] != user {
		return false
	}
	t.locker[num] = 0
	return true
}

func (t *LockingTree) ancesors(num int) bool {
	for {
		if num == -1 {
			return false
		}
		if t.locker[num] != 0 {
			return true
		}
		num = t.tree[num]
	}
}

func (t *LockingTree) unlockDescendant(num int) bool {
	q := []int{num}
	found := false
	for len(q) > 0 {
		nq := make([]int, 0)
		for _, i := range q {
			if t.locker[i] != 0 {
				t.locker[i] = 0
				found = true
			}
			nq = append(nq, t.children[i]...)
		}
		q = nq
	}
	return found
}

func (t *LockingTree) Upgrade(num, user int) bool {
	if t.locker[num] != 0 {
		return false
	}
	if t.ancesors(num) {
		return false
	}
	if !t.unlockDescendant(num) {
		return false
	}
	t.locker[num] = user
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
