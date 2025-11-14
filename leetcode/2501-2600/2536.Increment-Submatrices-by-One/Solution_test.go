package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		inputs  int
		queries [][]int
		expect  [][]int
	}{
		{"TestCase1", 3, [][]int{
			{1, 1, 2, 2}, {0, 0, 1, 1},
		}, [][]int{
			{1, 1, 0}, {1, 2, 1}, {0, 1, 1},
		}},
		{"TestCase2", 2, [][]int{{0, 0, 1, 1}}, [][]int{{1, 1}, {1, 1}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.queries)
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
