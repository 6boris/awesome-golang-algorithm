package Solution

import (
	"math"
	"strconv"
	"testing"
)

const epsilon float64 = 1e-5

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect float64
	}{
		{"TestCase1", [][]int{{0, 0, 1}, {2, 2, 1}}, 1.0000},
		{"TestCase2", [][]int{{0, 0, 2}, {1, 1, 1}}, 1.16667},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			diff := math.Abs(c.expect - got)
			if diff > epsilon {
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
