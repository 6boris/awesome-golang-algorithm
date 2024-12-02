package Solution

import "strings"

type trie1455 struct {
	end   int
	child [26]*trie1455
}

func (tree *trie1455) insert1455(word string, index int) {
	cur := tree
	for _, b := range word {
		if cur.child[b-'a'] == nil {
			cur.child[b-'a'] = &trie1455{end: index, child: [26]*trie1455{}}
		}
		cur = cur.child[b-'a']
	}
}

func (tree *trie1455) search1455(word string) int {
	cur := tree
	for _, b := range word {
		idx := b - 'a'
		if cur.child[idx] == nil {
			return -1
		}
		cur = cur.child[idx]
	}
	return cur.end
}

func Solution(sentence string, searchWord string) int {
	tree := &trie1455{end: -1}
	for i, w := range strings.Split(sentence, " ") {
		tree.insert1455(w, i+1)
	}
	return tree.search1455(searchWord)
}
