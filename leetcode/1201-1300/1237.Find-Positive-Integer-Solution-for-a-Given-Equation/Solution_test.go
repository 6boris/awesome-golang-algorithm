package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func f1(a, b int) int {
	return a + b
}
func f2(a, b int) int {
	return a * b
}
func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs func(int, int) int
		z      int
		expect [][]int
	}{
		{"TestCase1", f1, 5, [][]int{{4, 1}, {3, 2}, {2, 3}, {1, 4}}},
		{"TestCase2", f2, 480, [][]int{
			{480, 1}, {240, 2}, {160, 3}, {120, 4}, {96, 5}, {80, 6}, {60, 8}, {48, 10},
			{40, 12}, {32, 15}, {30, 16}, {24, 20}, {20, 24}, {16, 30}, {15, 32},
			{12, 40}, {10, 48}, {8, 60}, {6, 80}, {5, 96}, {4, 120}, {3, 160}, {2, 240}, {1, 480},
		}},
		{"TestCase3", f2, 5, [][]int{{5, 1}, {1, 5}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.z)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.z)
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
