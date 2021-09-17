package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		nums1  []int
		nums2  []int
		expect int
	}{
		{"TestCase", []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{"TestCase", []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
		{"TestCase", []int{0, 0, 0, 0, 1}, []int{1, 0, 0, 0, 0}, 4},
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

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
