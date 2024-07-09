package Solution

import (
	"strconv"
	"testing"
)

const (
	within = 1e-5
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect float64
	}{
		{"TestCase1", [][]int{
			{1, 2}, {2, 5}, {4, 3},
		}, 5.00000},
		{"TestCase2", [][]int{
			{5, 2}, {5, 4}, {10, 3}, {20, 1},
		}, 3.25000},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			diff := got - c.expect
			if diff < 0 {
				diff = -diff
			}
			if diff > within {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
