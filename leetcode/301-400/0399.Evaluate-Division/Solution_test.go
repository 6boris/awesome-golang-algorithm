package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		equations [][]string
		value     []float64
		queries   [][]string
		expect    []float64
	}{
		{"TestCase", [][]string{
			{"a", "b"},
			{"b", "c"},
		}, []float64{2.0, 3.0}, [][]string{
			{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"},
		}, []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.equations, c.value, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.equations, c.value, c.queries)
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
