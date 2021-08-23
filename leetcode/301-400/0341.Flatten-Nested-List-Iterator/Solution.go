package Solution

type NestedInteger struct {
	Val   int
	isInt bool
	list  []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.isInt
}

func (this NestedInteger) GetInteger() int {
	return this.Val
}

func (this NestedInteger) SetInteger(value int) {

}
func (this NestedInteger) Add(ele NestedInteger) {

}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.list
}

type NestedIterator struct {
	Raw []*NestedInteger
	idx int
}

func Constructor3(nestedList []*NestedInteger) *NestedIterator {
	raw := make([]*NestedInteger, 0)
	deepTraverse(nestedList, &raw)
	return &NestedIterator{
		Raw: raw,
		idx: 0,
	}
}

func deepTraverse(nestedList []*NestedInteger, raw *[]*NestedInteger) {
	for idx, ele := range nestedList {
		if ele.IsInteger() {
			*raw = append(*raw, nestedList[idx])
			continue
		}
		deepTraverse(ele.GetList(), raw)
	}
}

func (this *NestedIterator) Next() int {
	r := this.Raw[this.idx].GetInteger()
	this.idx++
	return r
}

func (this *NestedIterator) HasNext() bool {
	return this.idx < len(this.Raw)
}

func Solution(nestedList []*NestedInteger) []int {
	o := Constructor3(nestedList)
	expect := make([]int, 0)
	for o.HasNext() {
		expect = append(expect, o.Next())
	}
	return expect
}
