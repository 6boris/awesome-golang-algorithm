package Solution

type RangeFreqQuery struct {
	freq map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{
		freq: make(map[int][]int),
	}
	for i := range arr {
		r.freq[arr[i]] = append(r.freq[arr[i]], i)
	}
	return r
}

func search(n int, ok func(int) bool) int {
	// return indies[mid] >= left
	// return indies[mmid] > right
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if ok(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	indies := this.freq[value]
	a := search(len(indies), func(mid int) bool {
		return indies[mid] >= left
	})
	b := search(len(indies), func(mid int) bool {
		return indies[mid] > right
	})
	return b - a
}

type operation struct {
	left, right, value int
}

func Solution(arr []int, inputs []operation) []int {
	c := Constructor(arr)
	ret := make([]int, len(inputs))
	for i := range inputs {
		ret[i] = c.Query(inputs[i].left, inputs[i].right, inputs[i].value)
	}
	return ret
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
