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
		opts   []opt
		expect []int
	}{
		{"TestCase1", 5, []opt{{name: "r"}, {name: "r"}, {name: "u", val: 2}, {name: "r"}, {name: "r"}, {name: "r"}, {name: "r"}, {name: "u", val: 5}}, []int{1, 2, 2, 3, 4, 5}},
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
