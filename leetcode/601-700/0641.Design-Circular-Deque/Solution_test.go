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
		n      int
		opts   []op
		expect []any
	}{
		{"TestCase1", 3, []op{
			{"il", 1}, {"il", 2}, {"if", 3}, {"if", 4}, {"gr", 0}, {"", 0}, {"dl", 0}, {"if", 4}, {"gf", 0},
		}, []any{true, true, true, false, 2, true, true, true, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.opts)
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
