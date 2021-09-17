package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []Interval
		expect []Interval
	}{
		{
			"1 test 1",
			[]Interval{
				{Start: 1, End: 3},
				{Start: 15, End: 18},
				{Start: 2, End: 6},
				{Start: 8, End: 10},
			},
			[]Interval{
				{Start: 1, End: 6},
				{Start: 8, End: 10},
				{Start: 15, End: 18},
			},
		},
		{
			"2 test 2",
			[]Interval{
				{Start: 1, End: 4},
				{Start: 4, End: 5},
			},
			[]Interval{
				{Start: 1, End: 5},
			},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := merge(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
