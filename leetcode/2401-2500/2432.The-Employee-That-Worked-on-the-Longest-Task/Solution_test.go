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
		inputs int
		logs   [][]int
		expect int
	}{
		{"TestCase1", 10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}, 1},
		{"TestCase2", 26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}, 3},
		{"TestCase3", 2, [][]int{{0, 10}, {1, 20}}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.logs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.logs)
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
