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
		limit  int
		expect []int
	}{
		{"TestCase1", []int{1, 5, 3, 9, 8}, 2, []int{1, 3, 5, 8, 9}},
		{"TestCase2", []int{1, 7, 6, 18, 2, 1}, 3, []int{1, 6, 7, 18, 1, 2}},
		{"TestCase3", []int{1, 7, 28, 19, 10}, 3, []int{1, 7, 28, 19, 10}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.limit)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums, c.limit)
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
