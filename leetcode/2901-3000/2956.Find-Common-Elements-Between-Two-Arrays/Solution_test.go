package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		nums1, nums2 []int
		expect       []int
	}{
		{"TestCase1", []int{2, 3, 2}, []int{1, 2}, []int{2, 1}},
		{"TestCase2", []int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6}, []int{3, 4}},
		{"TestCase3", []int{3, 4, 2, 3}, []int{1, 5}, []int{0, 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums1, c.nums2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums1, c.nums2)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
