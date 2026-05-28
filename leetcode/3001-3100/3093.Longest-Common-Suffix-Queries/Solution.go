package Solution

import "sort"

type trieNode3093 struct {
	end   int
	child [26]*trieNode3093
}

func buildTree3093(tree *trieNode3093, word string, index int) {
	root := tree
	for i := len(word) - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if root.child[idx] == nil {
			root.child[idx] = &trieNode3093{end: index}
		}
		root = root.child[idx]
	}
}
func search3093(tree *trieNode3093, word string) int {
	root := tree
	for i := len(word) - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if root.child[idx] == nil {
			break
		}
		root = root.child[idx]
	}
	return root.end
}

func Solution(wordsContainer []string, wordsQuery []string) []int {
	indies := make([]int, len(wordsContainer))
	for i := range indies {
		indies[i] = i
	}
	sort.Slice(indies, func(i, j int) bool {
		wi, wj := wordsContainer[indies[i]], wordsContainer[indies[j]]
		if len(wi) == len(wj) {
			return indies[i] < indies[j]
		}
		return len(wi) < len(wj)
	})
	root := &trieNode3093{end: -1}
	for i := range indies {
		buildTree3093(root, wordsContainer[indies[i]], i)
	}

	ret := make([]int, len(wordsQuery))
	for i := range wordsQuery {
		idx := search3093(root, wordsQuery[i])
		if idx == -1 {
			idx = 0
		}
		ret[i] = indies[idx]
	}

	return ret
}
