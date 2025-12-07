package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		nums, queries []int
		x             int
		expect        []int
	}{
		{"TestCase1", []int{1, 3, 1, 7}, []int{1, 3, 2, 4}, 1, []int{0, -1, 2, -1}},
		{"TestCase2", []int{1, 2, 3}, []int{10}, 5, []int{-1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.queries, c.x)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.queries, c.x)
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
