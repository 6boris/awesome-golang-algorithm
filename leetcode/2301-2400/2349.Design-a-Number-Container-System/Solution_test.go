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
		inputs []input
		expect []int
	}{
		{"TestCase1", []input{
			{"find", 0, 10},
			{"change", 2, 10},
			{"change", 1, 10},
			{"change", 3, 10},
			{"change", 5, 10},
			{"find", 0, 10},
			{"change", 1, 20},
			{"find", 0, 10},
		}, []int{-1, 1, 2}},
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
