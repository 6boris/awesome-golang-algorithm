package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		width, height int
		inputs        []input
		expect        []any
	}{
		{"TestCase1", 6, 3, []input{
			{"step", 2},
			{"step", 2},
			{"pos", 0},
			{"dir", 0},
			{"step", 2},
			{"step", 1},
			{"step", 4},
			{"pos", 0},
			{"dir", 0},
		}, []any{
			[]int{4, 0}, "East", []int{1, 2}, "West",
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.width, c.height, c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.width, c.height, c.inputs)
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
