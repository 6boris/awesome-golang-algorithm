package Solution

type trieNode677 struct {
	val   int
	child [26]*trieNode677
}

func (this *trieNode677) insert(key string, val int) {
	root := this
	for _, b := range key {
		idx := b - 'a'
		if root.child[idx] == nil {
			root.child[idx] = &trieNode677{child: [26]*trieNode677{}, val: 0}
		}
		root.child[idx].val += val
		root = root.child[idx]
	}
}

func (this *trieNode677) sum(prefix string) int {
	root := this
	ans := 0
	for _, b := range prefix {
		idx := b - 'a'
		if root.child[idx] == nil {
			return 0
		}
		ans = root.child[idx].val
		root = root.child[idx]
	}
	return ans

}

type MapSum struct {
	tree  *trieNode677
	cache map[string]int
}

func Constructor677() MapSum {
	return MapSum{tree: &trieNode677{child: [26]*trieNode677{}}, cache: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	source := val
	v, ok := this.cache[key]
	if ok {
		val -= v
	}
	this.tree.insert(key, val)
	this.cache[key] = source
}

func (this *MapSum) Sum(prefix string) int {
	return this.tree.sum(prefix)
}

type opt struct {
	name string
	key  string
	val  int
}

func Solution(opts []opt) []int {
	c := Constructor677()
	ans := make([]int, 0)
	for _, op := range opts {
		if op.name == "i" {
			c.Insert(op.key, op.val)
			continue
		}
		ans = append(ans, c.Sum(op.key))
	}
	return ans
}
