package Solution

type seg struct {
	start, end int
}

type MyCalendar struct {
	books []seg
}

func Constructor() MyCalendar {
	return MyCalendar{
		books: make([]seg, 0),
	}
}

func (this *MyCalendar) Book(start, end int) bool {
	s := seg{start, end}
	if len(this.books) == 0 {
		this.books = append(this.books, s)
		return true
	}

	idx := len(this.books) - 1
	breakPoint := false
	for ; idx >= 0; idx-- {
		if s.end >= this.books[idx].end && s.start >= this.books[idx].end {
			breakPoint = true
			break
		}
		if (s.start >= this.books[idx].start && s.start < this.books[idx].end) || (s.end > this.books[idx].start && s.end < this.books[idx].end) || (s.end >= this.books[idx].end && s.start <= this.books[idx].start) {
			return false
		}
		if idx == 0 {
			breakPoint = true
		}
	}
	if breakPoint {
		if idx < 0 {
			this.books = append([]seg{s}, this.books...)
		} else {
			tmp := make([]seg, 0)
			tmp = append(tmp, this.books[:idx+1]...)
			tmp = append(tmp, s)
			tmp = append(tmp, this.books[idx+1:]...)
			this.books = tmp
		}
	}

	return breakPoint
}

func Solution(segs []seg, res []bool) bool {
	if len(segs) != len(res) {
		return false
	}
	c := Constructor()

	for idx := 0; idx < len(segs); idx++ {
		if res[idx] != c.Book(segs[idx].start, segs[idx].end) {
			return false
		}
	}
	return true
}
