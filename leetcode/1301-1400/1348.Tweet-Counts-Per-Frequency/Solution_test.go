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
		expect [][]int
	}{
		{"TestCase1", []opt{
			{"r", "tweet3", "", 0, 0, 0},
			{"r", "tweet3", "", 60, 0, 0},
			{"r", "tweet3", "", 10, 0, 0},
			{"", "tweet3", "minute", 0, 0, 59},
			{"", "tweet3", "minute", 0, 0, 60},
			{"r", "tweet3", "", 120, 0, 0},
			{"", "tweet3", "hour", 0, 0, 210},
		}, [][]int{{2}, {2, 1}, {4}}},
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
