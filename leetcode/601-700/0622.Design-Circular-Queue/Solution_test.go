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
		opts   []opt
		expect []any
	}{
		{"TestCase1", 3, []opt{
			{"en", 1}, {"en", 2}, {"en", 3},
			{"en", 4}, {"re", 0}, {"", 0},
			{"de", 0}, {"en", 4}, {"re", 0},
		},
			[]any{true, true, true, false, 3, true, true, true, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.opts)
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
