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
		n      int
		pos    []int
		s      string
		expect []int
	}{
		{"TestCase1", 3, []int{0, 1}, "RRDDLU", []int{1, 5, 4, 3, 1, 0}},
		{"TestCase2", 2, []int{1, 1}, "LURD", []int{4, 1, 0, 0}},
		{"TestCase3", 1, []int{0, 0}, "LRUD", []int{0, 0, 0, 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.pos, c.s)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.pos, c.s)
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
