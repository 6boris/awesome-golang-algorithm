package Solution

import (
	"container/heap"
)

const mod2818 = 1_000_000_007

type Pair2818 struct {
	value int
	index int
}

// PriorityQueue implements a priority queue for pairs based on value
type PriorityQueue []Pair2818

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value > pq[j].value // Max-heap based on value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Pair2818))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// Helper function to compute the number of distinct prime factors
func distinctPrimeFactors(x int) int {
	count := 0
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			count++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		count++ // If x is a prime number larger than sqrt(original x)
	}
	return count
}

func Solution(nums []int, k int) int {
	n := len(nums)
	primeScores := make([]int, n)

	for i := 0; i < n; i++ {
		primeScores[i] = distinctPrimeFactors(nums[i])
	}

	nextDominant := make([]int, n)
	prevDominant := make([]int, n)
	for i := 0; i < n; i++ {
		nextDominant[i] = n
		prevDominant[i] = -1
	}

	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && primeScores[stack[len(stack)-1]] < primeScores[i] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextDominant[topIndex] = i
		}

		if len(stack) > 0 {
			prevDominant[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	numOfSubarrays := make([]int64, n)
	for i := 0; i < n; i++ {
		numOfSubarrays[i] = int64(nextDominant[i]-i) * int64(i-prevDominant[i])
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	// Push each number and its index onto the priority queue
	for i := 0; i < n; i++ {
		heap.Push(pq, Pair2818{value: nums[i], index: i})
	}

	score := int64(1)

	// Process elements while there are operations left
	for k > 0 {
		// Get the element with the maximum value from the queue
		top := heap.Pop(pq).(Pair2818)
		num := top.value
		index := top.index

		// Calculate operations to apply on the current element
		operations := int64(k)
		if operations > numOfSubarrays[index] {
			operations = numOfSubarrays[index]
		}

		// Update the score
		score = (score * power(num, operations)) % mod2818

		// Reduce the remaining operations count
		k -= int(operations)
	}

	return int(score)
}

// Helper function to compute the power of a number modulo mod2818
func power(base int, exponent int64) int64 {
	res := int64(1)

	// Calculate the exponentiation using binary exponentiation
	for exponent > 0 {
		if exponent%2 == 1 {
			res = (res * int64(base)) % mod2818
		}
		base = (base * base) % mod2818
		exponent /= 2
	}

	return res
}
