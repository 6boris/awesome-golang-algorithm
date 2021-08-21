package Solution

const (
	push   = "push"
	pop    = "pop"
	top    = "top"
	getMin = "getMIn"
)

type MinStack struct {
	data, min []int
}

func Constructor4() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		this.data = append(this.data, val)
		this.min = append(this.min, val)
		return
	}

	top := this.min[len(this.min)-1]
	this.data = append(this.data, val)
	if top > val {
		top = val
	}

	this.min = append(this.min, top)
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		l := len(this.data)
		this.data = this.data[:l-1]
		this.min = this.min[:l-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}

	return -1
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

// action, push, pop, top, getMin
func Solution(actions []string, vals []int) []int {
	o := Constructor4()
	expect := make([]int, 0)
	for idx, act := range actions {
		if act == push {
			o.Push(vals[idx])
			continue
		}
		if act == pop {
			o.Pop()
			continue
		}

		if act == top {
			top := o.Top()
			expect = append(expect, top)
			continue
		}

		m := o.GetMin()
		expect = append(expect, m)
	}
	return expect
}
