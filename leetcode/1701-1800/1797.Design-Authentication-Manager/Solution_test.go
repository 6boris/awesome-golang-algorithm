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
		inputs []input
		expect []int
	}{
		{"TestCase1", 5, []input{
			{'r', "aaa", 1},
			{'g', "aaa", 2},
			{'c', "", 6},
			{'g', "bbb", 7},
			{'r', "aaa", 8},
			{'r', "bbb", 10},
			{'c', "", 15},
		}, []int{1, 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.inputs)
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
