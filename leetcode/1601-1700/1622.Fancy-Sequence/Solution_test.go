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
			{"append", 2}, {"add", 3}, {"append", 7},
			{"mul", 2}, {"get", 0}, {"add", 3},
			{"append", 10}, {"mul", 2},
			{"get", 0}, {"get", 1}, {"get", 2},
		}, []int{10, 26, 34, 20}},
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
