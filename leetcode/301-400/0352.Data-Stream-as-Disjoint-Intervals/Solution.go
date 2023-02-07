package Solution

type SummaryRanges struct {
	// stack?
	next []int
	max  int
}

func Constructor352() SummaryRanges {
	sr := SummaryRanges{
		next: make([]int, 10001),
	}
	return sr
}

// 1,3,7
func (this *SummaryRanges) AddNum(value int) {
	this.next[value] = 1
	if value > this.max {
		this.max = value
	}
}

// 3,_4_ 5, 7
// 3, 5
func (this *SummaryRanges) GetIntervals() [][]int {
	ans := make([][]int, 0)
	i := 0
	for i <= this.max {
		if this.next[i] == 0 {
			i++
			continue
		}
		start := i
		n := i
		for ; n <= this.max && this.next[n] == 1; n++ {

		}
		ans = append(ans, []int{start, n - 1})
		i = n
	}
	return ans
}

func Solution(ops []string, nums []int) [][][]int {
	sr := Constructor352()
	ans := make([][][]int, 0)
	for i, op := range ops {
		if op == "addNum" {
			sr.AddNum(nums[i])
			continue
		}

		tmp := sr.GetIntervals()
		ans = append(ans, tmp)
	}
	return ans
}
