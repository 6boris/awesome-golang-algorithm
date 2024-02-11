package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name           string
		inputs         [][]int
		source, target int
		expect         int
	}{
		{"TestCase1", [][]int{
			{1, 2, 7}, {3, 6, 7},
		}, 1, 6, 2},
		{"TestCase2", [][]int{
			{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13},
		}, 15, 12, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.source, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.inputs, c.source, c.target)
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
