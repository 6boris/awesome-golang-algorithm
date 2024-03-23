package Solution

type LUPrefix struct {
	prefix []bool
	index  int
}

func Constructor2424(n int) LUPrefix {
	return LUPrefix{prefix: make([]bool, n+1), index: 0}
}

func (this *LUPrefix) Upload(video int) {
	this.prefix[video] = true
	next := this.index + 1
	for ; next < len(this.prefix) && this.prefix[next]; next++ {
	}
	this.index = next - 1
}

func (this *LUPrefix) Longest() int {
	return this.index
}

func Solution(n int, opts []int) []int {
	c := Constructor2424(n)
	r := make([]int, 0)
	for _, op := range opts {
		if op == -1 {
			r = append(r, c.Longest())
			continue
		}
		c.Upload(op)
	}
	return r
}
