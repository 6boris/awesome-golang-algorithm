package Solution

import (
	"container/list"
)

func Solution(arr []int, k int) int {
	l := list.New()
	maxitem := -1
	for _, n := range arr {
		l.PushBack(n)
		if n > maxitem {
			maxitem = n
		}
	}
	winCount := 0
	cur := l.Front()
	for {
		next := cur.Next()
		a := cur.Value.(int)
		b := next.Value.(int)
		if a > b {
			winCount++
			l.MoveToBack(next)
		} else {
			l.MoveToBack(cur)
			cur = next
			winCount = 1
		}
		if winCount == k || cur.Value.(int) == maxitem {
			return cur.Value.(int)
		}
	}
}
