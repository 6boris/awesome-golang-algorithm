package Solution

type Iterator struct {
	data []int
	idx  int
}

func (this *Iterator) hasNext() bool {
	return this.idx < len(this.data)-1
}
func (this *Iterator) next() int {
	ans := this.data[this.idx]
	this.idx++
	return ans
}

type PeekingIterator struct {
	iter  *Iterator
	peekV int
}

func Constructor284(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter, peekV: -1}
}

func (this *PeekingIterator) hasNext() bool {
	if !this.iter.hasNext() {
		return this.peekV != -1
	}
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.peekV != -1 {
		ans := this.peekV
		this.peekV = -1
		return ans
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.peekV == -1 {
		this.peekV = this.iter.next()
	}
	return this.peekV
}

func Solution(iter *Iterator, options []string) []interface{} {
    ans := make([]interface{}, 0)
    o := Constructor284(iter)
    for _, op := range options {
        if op == "next" {
            ans = append(ans, o.next())
            continue
        }
        if op == "peek" {
            ans = append(ans, o.peek())
            continue
        }
        ans = append(ans, o.hasNext())
    }
    return ans
}
