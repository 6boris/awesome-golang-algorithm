package Solution

import (
	"reflect"
	"strconv"
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

func TestSolution2(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"TestCase", []int{1, 2, 3}, []int{1, 2, 4}},
		{"TestCase", []int{9, 9, 9}, []int{1, 0, 0, 0}},
		{"TestCase", []int{5, 9, 9}, []int{6, 0, 0}},
		{"TestCase", []int{}, []int{1}},
	}
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i+1), func(t *testing.T) {
			got := Solution(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.input)
			}
		})
	}
}
