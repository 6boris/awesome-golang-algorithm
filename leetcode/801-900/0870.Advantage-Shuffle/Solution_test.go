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
		{"TestCase1", []int{2, 7, 11, 15}, []int{1, 10, 4, 11}, []int{2, 11, 7, 15}},
		{"TestCase2", []int{12, 24, 8, 32}, []int{13, 25, 32, 11}, []int{24, 32, 8, 12}},
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
