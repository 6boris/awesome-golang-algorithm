package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []int
		expects []int
	}{
		{"test case 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			result := moveZeroes(testcase.inputs)
			if !reflect.DeepEqual(result, testcase.expects) {
				t.Fatalf("expected: %v, but got: %v, with input: %v ", testcase.expects, result, testcase.inputs)
			}

		})
	}
}

func TestMoveZeroes2(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []int
		expects []int
	}{
		{"TestCase", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"TestCase", []int{0}, []int{0}},
		{"TestCase", []int{1}, []int{1}},
	}

	for i, testcase := range cases {
		t.Run(testcase.name+" "+strconv.Itoa(i), func(t *testing.T) {
			result := moveZeroes2(testcase.inputs)
			if !reflect.DeepEqual(result, testcase.expects) {
				t.Fatalf("expected: %v, but got: %v, with input: %v ", testcase.expects, result, testcase.inputs)
			}

		})
	}
}
