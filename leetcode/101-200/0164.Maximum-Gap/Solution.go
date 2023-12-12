package Solution

import (
	"math"
)

func radixSort(nums []int) {
	maximum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}
	digits := 0
	r := maximum
	for r > 0 {
		digits++
		r /= 10
	}
	for i := 1; i <= digits; i++ {
		bucket := [10][]int{}
		p := math.Pow(10, float64(i-1))
		for _, n := range nums {
			bucketIndex := (n / int(p)) % 10
			bucket[bucketIndex] = append(bucket[bucketIndex], n)
		}
		idx := 0
		for j := 0; j < 10; j++ {
			for k := 0; k < len(bucket[j]); k++ {
				nums[idx] = bucket[j][k]
				idx++
			}
		}
	}
}
func Solution(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}
	radixSort(nums)
	ans := 0
	for i := 1; i < l; i++ {
		if r := nums[i] - nums[i-1]; r > ans {
			ans = r
		}
	}
	return ans
}
