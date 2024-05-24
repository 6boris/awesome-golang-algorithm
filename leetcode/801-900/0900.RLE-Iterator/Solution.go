package Solution

type RLEIterator struct {
	source []int
	index  int
}

func Constructor900(encoding []int) RLEIterator {
	rle := RLEIterator{source: make([]int, 0), index: 0}
	for i := 0; i < len(encoding)-1; i += 2 {
		if encoding[i] == 0 {
			continue
		}
		rle.source = append(rle.source, encoding[i], encoding[i+1])
	}
	return rle
}

func (this *RLEIterator) Next(n int) int {
	for ; this.index < len(this.source)-1 && this.source[this.index] < n; this.index += 2 {
		n -= this.source[this.index]
		this.source[this.index] = 0
	}
	if this.index == len(this.source) {
		return -1
	}
	this.source[this.index] -= n
	return this.source[this.index+1]
}

func Solution(encodings []int, input []int) []int {
	c := Constructor900(encodings)
	ans := make([]int, len(input))
	for i, n := range input {
		ans[i] = c.Next(n)
	}
	return ans
}
