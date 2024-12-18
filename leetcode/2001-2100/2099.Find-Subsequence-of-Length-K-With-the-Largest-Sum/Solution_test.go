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
		k      int
		expect []int
	}{
		{"TestCase1", []int{2, 1, 3, 3}, 2, []int{3, 3}},
		{"TestCase2", []int{-1, -2, 3, 4}, 3, []int{-1, 3, 4}},
		{"TestCase3", []int{3, 4, 3, 3}, 2, []int{3, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums, c.k)
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
