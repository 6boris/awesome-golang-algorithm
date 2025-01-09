package Solution

type trieNode2185 struct {
	c     int
	child [26]*trieNode2185
}

func insert(tree *trieNode2185, word string) {
	cur := tree
	for _, b := range word {
		idx := b - 'a'
		if cur.child[idx] == nil {
			cur.child[idx] = &trieNode2185{}
		}
		cur.child[idx].c++
		cur = cur.child[idx]
	}
}

func search(tree *trieNode2185, target string) int {
	cur := tree
	for _, b := range target {
		if cur.child[b-'a'] == nil {
			return 0
		}
		cur = cur.child[b-'a']
	}
	return cur.c
}

func Solution(words []string, pref string) int {
	tree := &trieNode2185{}
	for _, w := range words {
		insert(tree, w)
	}
	return search(tree, pref)
}
