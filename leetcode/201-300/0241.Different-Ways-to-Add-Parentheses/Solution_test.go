package Solution

import (
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect []int
	}{
		{"TestCase1", "2-1-1", []int{0, 2}},
		{"TestCase2", "2*3-4*5", []int{-34, -14, -10, -10, 10}},
		{"TestCase3", "10+5", []int{15}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := diffWaysToCompute(c.inputs)
			if len(got) != len(c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
			}
			m := make(map[int]int)
			for v := range got {
				m[v]++
			}
			for v := range c.expect {
				if _, ok := m[v]; !ok {
					t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
				}
			}

			got = diffWaysToCompute2(c.inputs)
			sort.Ints(got)
			if len(got) != len(c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
			}
			m = make(map[int]int)
			for v := range got {
				m[v]++
			}
			for v := range c.expect {
				if _, ok := m[v]; !ok {
					t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, got, c.inputs)
				}
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
