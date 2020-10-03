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
		{"1 test 1",
			[]Interval{
				Interval{Start: 1, End: 3},
				Interval{Start: 15, End: 18},
				Interval{Start: 2, End: 6},
				Interval{Start: 8, End: 10},
			},
			[]Interval{
				Interval{Start: 1, End: 6},
				Interval{Start: 8, End: 10},
				Interval{Start: 15, End: 18},
			}},
		{"2 test 2",
			[]Interval{
				Interval{Start: 1, End: 4},
				Interval{Start: 4, End: 5},
			},
			[]Interval{
				Interval{Start: 1, End: 5},
			}},
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
