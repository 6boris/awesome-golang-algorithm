package Solution

import (
	"container/heap"
	"math"
	"sort"
)

type Pair struct {
	wageRatio float64
	quality   int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	return p[i].wageRatio < p[j].wageRatio
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Solution(quality []int, wage []int, k int) float64 {
	n := len(quality)
	totalCost := float64(1<<63 - 1)
	currentTotalQuality := 0.0
	wageToQualityRatio := make(PairList, n)

	for i := 0; i < n; i++ {
		wageToQualityRatio[i] = Pair{float64(wage[i]) / float64(quality[i]), quality[i]}
	}

	sort.Sort(wageToQualityRatio)

	highestQualityWorkers := &MaxHeap{}
	heap.Init(highestQualityWorkers)

	for i := 0; i < n; i++ {
		heap.Push(highestQualityWorkers, wageToQualityRatio[i].quality)
		currentTotalQuality += float64(wageToQualityRatio[i].quality)

		if highestQualityWorkers.Len() > k {
			currentTotalQuality -= float64(heap.Pop(highestQualityWorkers).(int))
		}

		if highestQualityWorkers.Len() == k {
			totalCost = math.Min(totalCost, currentTotalQuality*wageToQualityRatio[i].wageRatio)
		}
	}

	return totalCost
}
