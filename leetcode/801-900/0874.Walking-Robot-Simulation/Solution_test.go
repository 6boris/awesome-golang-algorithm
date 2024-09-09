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
		commands  []int
		obstacles [][]int
		expect    int
	}{
		{"TestCase1", []int{4, -1, 3}, [][]int{}, 25},
		{"TestCase2", []int{4, -1, 4, -2, 4}, [][]int{{2, 4}}, 65},
		{"TestCase3", []int{6, -1, -1, 6}, [][]int{}, 36},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.commands, c.obstacles)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.commands, c.obstacles)
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
