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
		{"TestCase1", [][]int{{2, 5, 4}, {1, 5, 1}}, 4},
		{"TestCase2", [][]int{{3, 3, 1}, {8, 5, 2}}, 4},
		{"TestCase3", [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}, 7},
		{"TestCase4", [][]int{{2, 4, 6, 8, 10, 12}, {11, 9, 7, 5, 3, 1}}, 27},
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
