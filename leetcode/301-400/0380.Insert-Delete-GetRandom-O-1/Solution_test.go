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
		inputs []action
		expect bool
	}{
		{"TestCase1", []action{
			{
				name:     "insert",
				v:        1,
				opResult: true,
			},
			{
				name:     "remove",
				v:        2,
				opResult: false,
			},
			{
				name:     "insert",
				v:        2,
				opResult: true,
			},
			{
				name: "getRandom",
				exp: map[int]struct{}{
					1: {},
					2: {},
				},
			},
			{
				name:     "remove",
				v:        1,
				opResult: true,
			},
			{
				name:     "insert",
				v:        2,
				opResult: false,
			},
			{
				name: "getRandom",
				exp: map[int]struct{}{
					2: {},
				},
			},
		}, true},
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
