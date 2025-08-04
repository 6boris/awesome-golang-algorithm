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
		rows   int
		opts   []opt
		expect []int
	}{
		{"TestCase1", 3, []opt{
			{"get", "=5+7", 0},
			{"set", "A1", 10},
			{"get", "=A1+6", 0},
			{"set", "B2", 15},
			{"get", "=A1+B2", 0},
			{"reset", "A1", 0},
			{"get", "=A1+B2", 0},
		}, []int{12, 16, 25, 15}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.rows, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.rows, c.opts)
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
