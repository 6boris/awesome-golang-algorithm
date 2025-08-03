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
		ops          []op
		expect       []int
	}{
		{"TestCase1", []int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4}, []op{
			{
				name: 'c', a: 7,
			},
			{
				name: 'a', a: 3, b: 2,
			},
			{
				name: 'c', a: 8,
			},
			{
				name: 'c', a: 4,
			},
			{
				name: 'a', a: 0, b: 1,
			},
			{
				name: 'a', a: 1, b: 1,
			},
			{
				name: 'c', a: 7,
			},
		}, []int{8, 2, 1, 11}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums1, c.nums2, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums1, c.nums2, c.ops)
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
