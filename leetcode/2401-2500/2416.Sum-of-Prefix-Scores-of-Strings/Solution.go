package Solution

type trieNode2416 struct {
	count int
	child [26]*trieNode2416
}

func (r *trieNode2416) insert(word string) {
	walker := r
	for i := range word {
		index := word[i] - 'a'
		if walker.child[index] == nil {
			walker.child[index] = &trieNode2416{}
		}
		walker.child[index].count++
		walker = walker.child[index]
	}
}

func (r *trieNode2416) score(word string) int {
	walker := r
	x := 0
	for i := range word {
		index := word[i] - 'a'
		x += walker.child[index].count
		walker = walker.child[index]
	}
	return x
}
func Solution(words []string) []int {
	ans := make([]int, len(words))
	tree := &trieNode2416{}
	for _, w := range words {
		tree.insert(w)
	}
	d := make(map[string]int)
	for i, w := range words {
		if v, ok := d[w]; ok {
			ans[i] = v
			continue
		}
		score := tree.score(w)
		ans[i] = score
		d[w] = score
	}

	return ans
}
