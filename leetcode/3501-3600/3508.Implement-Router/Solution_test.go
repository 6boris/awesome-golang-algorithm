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
		ops    []op
		expect []any
	}{
		{"TestCase1", 3, []op{
			{"add", 1, 4, 90},
			{"add", 2, 5, 90},
			{"add", 1, 4, 90},
			{"add", 3, 5, 95},
			{"add", 4, 5, 105},
			{"for", 0, 0, 0},
			{"add", 5, 2, 110},
			{"get", 5, 100, 110},
		}, []any{true, true, false, true, true, []int{2, 5, 90}, true, 1}},

		{"TestCase12", 2, []op{
			{"add", 7, 4, 90}, {"for", 0, 0, 0}, {"for", 0, 0, 0},
		}, []any{true, []int{7, 4, 90}, []int{}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.ops)
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
