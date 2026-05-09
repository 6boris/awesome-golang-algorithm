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
		expect int64
	}{
		{"TestCase1", [][]int{{0, 0, 0, 0, 0}, {0, 0, 3, 0, 0}, {0, 1, 0, 0, 0}, {5, 0, 0, 3, 0}, {0, 0, 0, 0, 2}}, 11},
		{"TestCase2", [][]int{{10, 9, 0, 0, 15}, {7, 1, 0, 8, 0}, {5, 20, 0, 11, 0}, {0, 0, 0, 1, 2}, {8, 12, 1, 10, 3}}, 94},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
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
