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
		inputs []operation
		expect []int
	}{
		{"TestCase1", []operation{{name: "put", k: 1, v: 1}, {name: "put", k: 2, v: 2}, {name: "get", k: 1}, {name: "get", k: 3}, {name: "put", k: 2, v: 1}, {name: "get", k: 2}, {name: "remove", k: 2}, {name: "get", k: 2}}, []int{1, -1, 1, -1}},
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
