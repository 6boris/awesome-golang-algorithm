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
		n, start, end int
		edges         [][]int
		succProb      []float64
		expect        float64
	}{
		{"TestCase1", 3, 0, 2, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0.25000},
		{"TestCase2", 3, 0, 2, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0.30000},
		{"TestCase3", 3, 0, 2, [][]int{{0, 1}}, []float64{0.5}, 0.00000},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges, c.succProb, c.start, c.end)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v %v",
					c.expect, got, c.n, c.edges, c.succProb, c.start, c.end)
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
