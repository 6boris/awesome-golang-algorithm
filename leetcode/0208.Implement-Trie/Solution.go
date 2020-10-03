package Solution

type Trie struct {
	IsTerminated bool
	Children     map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		IsTerminated: false,
		Children:     make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{IsTerminated: false, Children: make(map[rune]*Trie)}
			parent.Children[ch] = newChild
			parent = newChild
		}
	}
	parent.IsTerminated = true
}

func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.IsTerminated
}

func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, ch := range prefix {
		if child, ok := parent.Children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
