package Solution

import (
	"math/rand"
	"time"
)

type Shuffle struct {
	origin []int
	r      *rand.Rand
}

func Constructor(nums []int) Shuffle {
	return Shuffle{
		origin: nums,
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Shuffle) Reset() []int {
	this.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return this.origin
}

func (this *Shuffle) Shuffle() []int {
	n := len(this.origin)
	cur := make([]int, len(this.origin))
	copy(cur, this.origin)
	for i := n - 1; i > 0; i-- {
		j := this.r.Intn(i + 1)
		cur[i], cur[j] = cur[j], cur[i]
	}
	return cur
}
func Solution(nums []int, op []string) [][]int {
	c := Constructor(nums)
	var ret [][]int
	for _, name := range op {
		if name == "shuffle" {
			ret = append(ret, c.Shuffle())
			continue
		}
		ret = append(ret, c.Reset())
	}
	return ret
}
