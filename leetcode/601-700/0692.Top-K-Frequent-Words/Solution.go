package Solution

import (
	"container/heap"
)

type topWord struct {
	word  string
	count int
}

type topWordList []topWord

func (tl *topWordList) Len() int {
	return len(*tl)
}

func (tl *topWordList) Swap(i, j int) {
	(*tl)[i], (*tl)[j] = (*tl)[j], (*tl)[i]
}

func (tl *topWordList) Less(i, j int) bool {
	if (*tl)[i].count == (*tl)[j].count {
		return (*tl)[i].word < (*tl)[j].word
	}

	return (*tl)[i].count > (*tl)[j].count
}

func (tl *topWordList) Push(x interface{}) {
	*tl = append(*tl, x.(topWord))
}

func (tl *topWordList) Pop() interface{} {
	old := *tl
	n := len(old)
	item := old[n-1]
	*tl = old[:n-1]
	return item
}

func Solution(words []string, k int) []string {
	ans := make([]string, 0)
	wordBucket := make(map[string]int)
	for _, word := range words {
		wordBucket[word]++
	}
	h := topWordList([]topWord{})
	heap.Init(&h)
	for k, v := range wordBucket {
		heap.Push(&h, topWord{word: k, count: v})
	}

	for ; k > 0; k-- {
		top := heap.Pop(&h)
		ans = append(ans, top.(topWord).word)
	}
	return ans
}
