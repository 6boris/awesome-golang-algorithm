package Solution

import (
	"container/heap"
)

type fr struct {
	name   string
	rating int
	index  *int
}
type frs []fr

func (f *frs) Len() int {
	return len(*f)
}

func (f *frs) Swap(i, j int) {
	*(*f)[i].index = j
	*(*f)[j].index = i
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *frs) Less(i, j int) bool {
	a := (*f)[i]
	b := (*f)[j]
	if a.rating == b.rating {
		return a.name < b.name
	}
	return a.rating > b.rating
}

func (f *frs) Push(x interface{}) {
	*f = append(*f, x.(fr))
}

func (f *frs) Pop() interface{} {
	old := *f
	l := len(old)
	x := old[l-1]
	*f = old[:l-1]
	return x
}

type FoodRatings struct {
	// Record the food's index in the heap via a pointer,
	// which can be modified synchronously
	indies map[string]*int
	// Record the cusisine to which the food belongs
	f2c map[string]string
	// each cuisine maintains a heap
	h map[string]*frs
}

func Constructor2353(foods []string, cuisines []string, ratings []int) FoodRatings {
	r := FoodRatings{indies: make(map[string]*int), f2c: make(map[string]string), h: make(map[string]*frs)}
	for idx, cuisine := range cuisines {
		r.f2c[foods[idx]] = cuisine
		if _, ok := r.h[cuisine]; !ok {
			r.h[cuisine] = &frs{}
		}

		l := r.h[cuisine].Len()
		r.indies[foods[idx]] = &l
		heap.Push(r.h[cuisine], fr{name: foods[idx], rating: ratings[idx], index: &l})
	}
	return r
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	// get the index of the food in the heap
	curIndex := this.indies[food]
	// get the cuisine of  food
	cuisine := this.f2c[food]
	// change rating
	arr := this.h[cuisine]
	(*arr)[*curIndex].rating = newRating
	this.h[cuisine] = arr
	// fix heap
	heap.Fix(this.h[cuisine], *curIndex)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	x := this.h[cuisine]
	return (*x)[0].name
}

type opt struct {
	name      string
	commonStr string
	rating    int
}

func Solution(foods []string, cuisines []string, ratings []int, opts []opt) []string {
	c := Constructor2353(foods, cuisines, ratings)
	ans := make([]string, 0)
	for _, op := range opts {
		if op.name == "h" {
			ans = append(ans, c.HighestRated(op.commonStr))
			continue
		}
		c.ChangeRating(op.commonStr, op.rating)
	}
	return ans
}
