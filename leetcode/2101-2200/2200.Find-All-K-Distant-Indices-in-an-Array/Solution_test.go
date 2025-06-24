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
		key, k int
		expect []int
	}{
		{"TestCase1", []int{3, 4, 9, 1, 3, 9, 5}, 9, 1, []int{1, 2, 3, 4, 5, 6}},
		{"TestCase2", []int{2, 2, 2, 2, 2}, 2, 2, []int{0, 1, 2, 3, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.key, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.key, c.k)
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
