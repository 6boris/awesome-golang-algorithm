package Solution

import "container/heap"

type IntHeap []int

func (ih IntHeap) Len() int           { return len(ih) }
func (ih IntHeap) Less(i, j int) bool { return ih[i] < ih[j] }
func (ih IntHeap) Swap(i, j int)      { ih[i], ih[j] = ih[j], ih[i] }
func (ih *IntHeap) Push(num interface{}) {
	*ih = append(*ih, num.(int))
}

func (ih *IntHeap) Pop() interface{} {
	old := *ih
	n := ih.Len()
	num := old[n-1]
	*ih = old[:n-1]
	return num
}

type MedianFinder struct {
	Small IntHeap
	Large IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{[]int{}, []int{}}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.Small.Len() == mf.Large.Len() {
		heap.Push(&mf.Large, -num)
		val := heap.Pop(&mf.Large)
		heap.Push(&mf.Small, -(val.(int)))
	} else {
		heap.Push(&mf.Small, num)
		val := heap.Pop(&mf.Small)
		heap.Push(&mf.Large, -(val.(int)))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.Small.Len() == mf.Large.Len() {
		return float64(mf.Small[0]-mf.Large[0]) / 2
	}
	return float64(mf.Small[0])
}
