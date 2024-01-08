package Solution

import "strings"

type trieNode648 struct {
	child [26]*trieNode648
	end   bool
}

func (t *trieNode648) insert(word string) {
	root := t
	for i, b := range word {
		idx := b - 'a'
		if root.child[idx] == nil {
			root.child[idx] = &trieNode648{child: [26]*trieNode648{}, end: false}
		}
		if i == len(word)-1 {
			root.child[idx].end = true
		}
		root = root.child[idx]
	}
}

func (t *trieNode648) search(word string) int {
	idx := 0
	root := t
	for _, b := range word {
		ix := b - 'a'
		idx++
		if root.child[ix] == nil {
			return -1
		}
		if root.child[ix].end {
			return idx
		}
		root = root.child[ix]
	}
	return -1
}

func Solution(dictionary []string, sentence string) string {

	t := &trieNode648{child: [26]*trieNode648{}, end: false}
	for _, root := range dictionary {
		t.insert(root)
	}
	sb := strings.Builder{}

	words := strings.Split(sentence, " ")

	for idx, word := range words {
		i := t.search(word)
		if i == -1 {
			sb.WriteString(word)
		} else {
			sb.WriteString(word[:i])
		}
		if idx != len(words)-1 {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}
