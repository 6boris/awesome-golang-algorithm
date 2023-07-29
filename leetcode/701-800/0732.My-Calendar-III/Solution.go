package Solution

import "sort"

type MyCalendarThree struct {
	store map[int]int
	keys  []int
}

// 这种感觉
// 从一个端点开始，在没有遇到结束端点的时候，我们多次遇到其他的开始端点，那就是重合
// 例如开始写入 1， 20, 然后跟进2，15
func Constructor732() MyCalendarThree {
	return MyCalendarThree{
		store: make(map[int]int),
		keys:  make([]int, 0),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	_, ok1 := this.store[startTime]
	_, ok2 := this.store[endTime]
	this.store[startTime]++
	this.store[endTime]--
	if !ok1 {
		this.keys = append(this.keys, startTime)
	}
	if !ok2 {
		this.keys = append(this.keys, endTime)
	}
	sort.Ints(this.keys)
	cur, ans := 0, 0
	for _, key := range this.keys {
		cur += this.store[key]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

func Solution(events [][]int) []int {
	c := Constructor732()
	ans := make([]int, len(events))
	for idx, event := range events {
		ans[idx] = c.Book(event[0], event[1])
	}
	return ans
}
