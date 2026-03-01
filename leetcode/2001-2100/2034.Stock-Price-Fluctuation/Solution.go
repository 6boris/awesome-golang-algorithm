package Solution

import "container/heap"

type stock2034 struct {
	price, timestamp, index int
}

type heap2034 struct {
	data []*stock2034
	cmp  func(a, b *stock2034) bool
}

func (h *heap2034) Len() int {
	return len(h.data)
}

func (h *heap2034) Less(i, j int) bool {
	return h.cmp(h.data[i], h.data[j])
}

func (h *heap2034) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].index = i
	h.data[j].index = j
}

func (h *heap2034) Push(x any) {
	item := x.(*stock2034)
	item.index = len(h.data)
	h.data = append(h.data, item)
}

func (h *heap2034) Pop() any {
	old := h.data
	size := len(old)
	x := old[size-1]
	h.data = old[:size-1]
	return x
}

type StockPrice struct {
	lessHeap    *heap2034
	greaterHeap *heap2034

	lessCache    map[int]*stock2034
	greaterCache map[int]*stock2034

	maxTimestamp int
	currentPrice int
}

func Constructor() StockPrice {
	return StockPrice{
		lessHeap: &heap2034{data: make([]*stock2034, 0), cmp: func(a, b *stock2034) bool {
			return a.price < b.price
		}},
		greaterHeap: &heap2034{
			data: make([]*stock2034, 0),
			cmp: func(a, b *stock2034) bool {
				return a.price > b.price
			},
		},
		lessCache:    make(map[int]*stock2034),
		greaterCache: make(map[int]*stock2034),
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if timestamp >= this.maxTimestamp {
		this.maxTimestamp = timestamp
		this.currentPrice = price
	}
	lessValue, ok := this.lessCache[timestamp]
	if ok {
		if lessValue.price == price {
			return
		}

		lessValue.price = price
		greaterValue := this.greaterCache[timestamp]
		greaterValue.price = price
		heap.Fix(this.lessHeap, lessValue.index)
		heap.Fix(this.greaterHeap, greaterValue.index)
		return
	}

	lessItem := &stock2034{price: price, timestamp: timestamp}
	greaterItem := &stock2034{price: price, timestamp: timestamp}
	this.lessCache[timestamp] = lessItem
	this.greaterCache[timestamp] = greaterItem
	heap.Push(this.lessHeap, lessItem)
	heap.Push(this.greaterHeap, greaterItem)
}

func (this *StockPrice) Current() int {
	return this.currentPrice

}

func (this *StockPrice) Maximum() int {
	return this.greaterHeap.data[0].price
}

func (this *StockPrice) Minimum() int {
	return this.lessHeap.data[0].price
}

type input struct {
	name string

	timestamp, price int
}

func Solution(opts []input) []int {
	ret := make([]int, 0)
	c := Constructor()
	for i := range opts {
		if opts[i].name == "update" {
			c.Update(opts[i].timestamp, opts[i].price)
			continue
		}
		if opts[i].name == "current" {
			ret = append(ret, c.Current())
			continue
		}
		if opts[i].name == "max" {
			ret = append(ret, c.Maximum())
			continue
		}
		ret = append(ret, c.Minimum())
	}
	return ret
}
