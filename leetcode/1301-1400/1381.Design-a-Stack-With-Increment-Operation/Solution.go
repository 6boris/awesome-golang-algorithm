package Solution

type CustomStack struct {
	index int
	stack []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, maxSize),
		index: -1,
	}
}

func (this *CustomStack) Push(x int) {
	if this.index == len(this.stack)-1 {
		return
	}

	this.index++
	this.stack[this.index] = x
}

func (this *CustomStack) Pop() int {
	if this.index == -1 {
		return -1
	}

	top := this.stack[this.index]
	this.index--
	return top
}

func (this *CustomStack) Increment(k int, val int) {
	if this.index == -1 {
		return
	}

	// 1, 2, 3, 4,
	// 3
	for i := 0; i <= this.index && i < k; i++ {
		this.stack[i] += val
	}
}

type op struct {
	name string
	a, b int
}

func Solution(maxSize int, opts []op) []int {
	c := Constructor(maxSize)
	ans := make([]int, 0)
	for _, o := range opts {
		if o.name == "push" {
			c.Push(o.a)
			continue
		}
		if o.name == "i" {
			c.Increment(o.a, o.b)
			continue
		}

		ans = append(ans, c.Pop())
	}
	return ans
}
