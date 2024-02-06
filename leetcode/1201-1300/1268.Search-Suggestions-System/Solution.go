package Solution

import "fmt"

type trieNode1268 struct {
	end   bool
	child [26]*trieNode1268
}

func insertString(root *trieNode1268, str string) {
	w := root
	for i, b := range str {
		child := w.child[b-'a']
		if child == nil {
			w.child[b-'a'] = &trieNode1268{end: false, child: [26]*trieNode1268{}}
		}
		if i == len(str)-1 {
			w.child[b-'a'].end = true
		}
		w = w.child[b-'a']
	}
}

func prefixNode(root *trieNode1268, prefix string) *trieNode1268 {
	w := root
	for _, b := range prefix {
		child := w.child[b-'a']
		if child == nil {
			return nil
		}
		w = child
	}
	return w
}
func searchWords(root *trieNode1268, base string, count *int, res *[]string) {
	if *count == 0 {
		return
	}
	if root.end {
		*res = append(*res, base)
		*count--
	}
	for i := 0; i < 26; i++ {
		if root.child[i] != nil {
			c := fmt.Sprintf("%s%c", base, i+'a')
			searchWords(root.child[i], c, count, res)
		}
	}
}
func Solution(products []string, searchWord string) [][]string {
	l := len(searchWord)
	root := &trieNode1268{end: false, child: [26]*trieNode1268{}}
	for _, p := range products {
		insertString(root, p)
	}

	result := make([][]string, l)
	for i := 0; i < l; i++ {
		base := searchWord[:i+1]
		tailNode := prefixNode(root, base)
		if tailNode == nil {
			continue
		}
		count := 3
		searchWords(tailNode, base, &count, &result[i])
	}
	return result
}
