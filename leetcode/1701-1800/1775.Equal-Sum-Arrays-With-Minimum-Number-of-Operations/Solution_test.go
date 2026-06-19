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
		expect       int
	}{
		{"TestCase1", []int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}, 3},
		{"TestCase2", []int{1, 1, 1, 1, 1, 1, 1}, []int{6}, -1},
		{"TestCase3", []int{6, 6}, []int{1}, 3},
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
