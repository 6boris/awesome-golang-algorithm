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
		{"TestCase1", [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {0, 0, 0}},
			[][]int{{19, 19, 19}, {18, 18, 18}, {17, 17, 17}, {16, 16, 16}, {15, 15, 15}, {14, 14, 14}, {13, 13, 13}, {12, 12, 12}, {11, 11, 11}, {10, 10, 10}, {9, 9, 9}, {8, 8, 8}, {7, 7, 7}, {6, 6, 6}, {5, 5, 5}, {4, 4, 4}, {3, 3, 3}, {2, 2, 2}, {1, 1, 1}, {0, 0, 0}}},
		{"TestCase2", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
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
