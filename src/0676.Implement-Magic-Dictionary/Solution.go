package Solution

type MagicDictionary struct {
	Children [26]*MagicDictionary
	End      bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (root *MagicDictionary) Insert(word string) {
	start := root
	for i, w := range word {
		if node := start.Children[w-'a']; node != nil {
			start = node
		} else {
			for _, nw := range word[i:] {
				start.Children[nw-'a'] = &MagicDictionary{}
				start = start.Children[nw-'a']
			}
			break
		}
	}
	start.End = true

}

/** Build a dictionary through a list of words */
func (root *MagicDictionary) BuildDict(dict []string) {
	for _, w := range dict {
		root.Insert(w)
	}
}

//the second parameter is mismatched character, the third one is word length, and the fourth is the word access index
func (root *MagicDictionary) SearchHelper(word string, miss int, wl int, ind int) bool {
	//if mismatch is more than 1
	if miss > 1 {
		return false
	}
	//if end of the word, then must have exact one mismatch
	if wl == ind {
		return root.End && miss == 1
	}
	//current access character of the word
	ch := int(word[ind] - 'a')
	for i, nc := range root.Children {
		if nc != nil {
			nm := miss
			if i != ch {
				nm += 1
			}
			if nc.SearchHelper(word, nm, wl, ind+1) {
				return true
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (root *MagicDictionary) Search(word string) bool {
	return root.SearchHelper(word, 0, len(word), 0)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
