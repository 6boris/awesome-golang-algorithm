package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name   string
		inputs []int
		expect []int
	}{
		{"1 digit is zero", []int{0}, []int{1}},
		{"1 digit is 9", []int{9}, []int{1, 0}},
		{"2 normal digits", []int{2, 3}, []int{2, 4}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := plusOne(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
