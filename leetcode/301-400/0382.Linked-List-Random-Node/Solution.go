package Solution

import (
	"math/rand"
	"sort"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type Solution382 struct {
	length int
	head   *ListNode
}

func Constructor(head *ListNode) Solution382 {
	s := Solution382{length: 0, head: head}
	for walker := head; walker != nil; walker = walker.Next {
		s.length++
	}
	return s
}

func (this *Solution382) GetRandom() int {
	if this.head == nil {
		return -1
	}
	seed := rand.NewSource(time.Now().UnixNano())
	rr := rand.New(seed)
	index := rr.Intn(this.length)

	walker := this.head
	for ; walker != nil && index > 0; walker = walker.Next {
		index--
	}
	return walker.Val
}

func Solution(count int, head *ListNode) []int {
	s := Constructor(head)
	ans := make([]int, count)
	for i := 0; i < count; i++ {
		ans[i] = s.GetRandom()
	}
	sort.Ints(ans)
	return ans
}
