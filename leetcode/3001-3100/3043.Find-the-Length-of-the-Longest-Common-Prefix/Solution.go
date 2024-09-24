package Solution

import (
	"strconv"
)

type trieNode3043 struct {
	child [10]*trieNode3043
}

func (r *trieNode3043) insert(intStr string) {
	walker := r
	for i := range intStr {
		index := intStr[i] - '0'
		if walker.child[index] == nil {
			walker.child[index] = &trieNode3043{child: [10]*trieNode3043{}}
		}
		walker = walker.child[index]
	}
}

func (r *trieNode3043) prefixMatch(intStr string) int {
	walker := r
	i := 0
	for ; i < len(intStr); i++ {
		index := intStr[i] - '0'
		if walker.child[index] == nil {
			break
		}
		walker = walker.child[index]
	}
	return i
}

func Solution(arr1 []int, arr2 []int) int {
	ans := 0
	tree := &trieNode3043{}
	for _, intVal := range arr1 {
		v := strconv.Itoa(intVal)
		tree.insert(v)
	}
	for _, intVal := range arr2 {
		v := strconv.Itoa(intVal)
		ans = max(ans, tree.prefixMatch(v))
	}
	return ans
}
