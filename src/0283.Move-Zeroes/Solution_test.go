package Solution

import (
	"reflect"
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
