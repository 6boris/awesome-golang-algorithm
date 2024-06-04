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
		inputs []op
		expect []int
	}{
		{"TestCase1", []op{
			{"a", 3}, {"a", 0},
			{"a", 2},
			{"a", 5},
			{"a", 4},
			{"", 2}, {"", 3},
			{"", 4}, {"a", 8}, {"", 2},
		}, []int{20, 40, 0, 32}},
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
