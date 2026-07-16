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
		inputs  string
		queries [][]int
		expect  []int
	}{
		{"TestCase1", "10203004", [][]int{
			{0, 7}, {1, 3}, {4, 6},
		}, []int{12340, 4, 9}},
		{"TestCase2", "1000", [][]int{
			{0, 3}, {1, 1},
		}, []int{1, 0}},
		{"TestCase3", "9876543210", [][]int{{0, 9}}, []int{444444137}},
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
