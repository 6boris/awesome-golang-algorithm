package Solution

type DataStream struct {
	value, k, cnt int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		value: value,
		k:     k,
		cnt:   0,
	}
}

func (this *DataStream) Consec(num int) bool {
	if num == this.value {
		this.cnt++
		return this.cnt >= this.k
	}
	this.cnt = 0
	return false
}

func Solution(value, k int, consec []int) []bool {
	ret := make([]bool, len(consec))
	c := Constructor(value, k)
	for i := range consec {
		ret[i] = c.Consec(consec[i])
	}
	return ret
}
