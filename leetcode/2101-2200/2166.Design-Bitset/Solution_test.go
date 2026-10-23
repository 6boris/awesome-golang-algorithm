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
		size   int
		inputs []operation
		expect []any
	}{
		{"TestCase1", 5, []operation{
			{"fix", 3},
			{"fix", 1},
			{"flip", 0},
			{"all", 0},
			{"unfix", 0},
			{"flip", 0},
			{"one", 0},
			{"unfix", 0},
			{"count", 0},
			{"str", 0},
		}, []any{false, true, 2, "01010"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.size, c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.size, c.inputs)
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
