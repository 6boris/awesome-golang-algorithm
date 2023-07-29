package Solution

type FreqStack struct {
	stack, freq []int
	item        map[int]int
}

func Constructor895() FreqStack {
	return FreqStack{
		stack: make([]int, 0),
		freq:  make([]int, 0),
		item:  make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	// 5, 7, 5, 4,
	// 5, 7, 5, 5,
	this.item[val]++
	l := len(this.freq)
	this.stack = append(this.stack, val)
	if l == 0 {
		this.freq = append(this.freq, val)
		return
	}
	latest := this.freq[l-1]
	count := this.item[latest]
	if count <= this.item[val] {
		latest = val
	}
	this.freq = append(this.freq, latest)
}

func (this *FreqStack) Pop() int {
	l := len(this.freq) - 1
	which := this.freq[l]

	this.item[which]--
	store := make([]int, 0)
	for ; l >= 0; l-- {
		if this.stack[l] == which {
			break
		}
		store = append(store, this.stack[l])
		this.item[this.stack[l]]--
	}

	this.stack = this.stack[:l]
	this.freq = this.freq[:l]
	for i := len(store) - 1; i >= 0; i-- {
		this.Push(store[i])
	}

	return which
}

type opt struct {
	name string
	v    int
}

func Solution(ops []opt) []int {
	c := Constructor895()
	ans := make([]int, 0)
	for _, op := range ops {
		if op.name == "push" {
			c.Push(op.v)
			continue
		}
		ans = append(ans, c.Pop())
	}
	return ans
}
