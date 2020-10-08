package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	t.Run("test-1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3

		merge(nums1, m, nums2, n)

		fmt.Println(nums1)

	})
}
