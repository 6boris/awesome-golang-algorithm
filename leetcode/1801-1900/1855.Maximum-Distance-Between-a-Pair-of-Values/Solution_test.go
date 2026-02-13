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
		{"TestCase1", []int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}, 2},
		{"TestCase2", []int{2, 2, 2}, []int{10, 10, 1}, 1},
		{"TestCase3", []int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}, 2},
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
