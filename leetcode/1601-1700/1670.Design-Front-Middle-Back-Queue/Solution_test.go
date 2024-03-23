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
		inputs []opt
		expect []int
	}{
		{"TestCase1", []opt{
			{"pushFront", 1},
			{"pushBack", 2},
			{"pushMiddle", 3},
			{"pushMiddle", 4},
			{"popFront", 0},
			{"popMiddle", 0},
			{"popMiddle", 0},
			{"popBack", 0},
			{"popFront", 0},
		}, []int{1, 3, 4, 2, -1}},
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
