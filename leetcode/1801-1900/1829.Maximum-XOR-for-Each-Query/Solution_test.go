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
		nums   []int
		mb     int
		expect []int
	}{
		{"TestCase1", []int{0, 1, 1, 3}, 2, []int{0, 3, 2, 3}},
		{"TestCase2", []int{2, 3, 4, 7}, 3, []int{5, 2, 6, 5}},
		{"TestCase3", []int{0, 1, 2, 2, 5, 7}, 3, []int{4, 3, 6, 4, 6, 7}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.mb)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums, c.mb)
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
