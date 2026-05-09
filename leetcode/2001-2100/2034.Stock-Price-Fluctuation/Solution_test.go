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
		inputs []input
		expect []int
	}{
		{"TestCase1", []input{
			{"update", 1, 10},
			{"update", 2, 5},
			{"current", 0, 0},
			{"max", 0, 0},
			{"update", 1, 3},
			{"max", 0, 0},
			{"update", 4, 2},
			{"min", 0, 0},
		}, []int{5, 10, 5, 2}},
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
