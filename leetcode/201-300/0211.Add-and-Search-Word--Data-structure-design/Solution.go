package Solution

type WordDictionary struct {
    trie map[int][]string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
    return WordDictionary{make(map[int][]string)};
}

func (this *WordDictionary) AddWord(word string)  {
    k := len(word);
    if v, ok := this.trie[k]; ok {
        this.trie[k] = append(v, word);
    } else {
        this.trie[k] = []string{word};
    }
}

func (this *WordDictionary) Search(word string) bool {
    k := len(word);
    words, ok := this.trie[k];
    if !ok {
        return false;
    }
    for _, cword := range words {
        found := true;
        for i := 0; i < len(cword); i++ {
            if word[i] != '.' && word[i] != cword[i] {
                found = false;
            }
        }
        if found {
            return true;
        }
    }
    return false;
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */