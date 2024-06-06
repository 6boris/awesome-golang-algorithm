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
		inputs [][]int
		expect [][]int
	}{
		{"TestCase1", [][]int{
			{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
		}, [][]int{
			{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0},
		}},
		{"TestCase2", [][]int{{1, 1}, {1, 0}}, [][]int{{1, 1}, {1, 1}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			Solution(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, c.inputs, c.inputs)
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
