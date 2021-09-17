package Solution

import (
	"reflect"
	"testing"
)

func TestConTainsDuplicate(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []int
		expects bool
	}{
		{"Test case 1", []int{1, 2, 3, 1}, true},
		{"Test case 2", []int{1, 2, 3, 4}, false},
		{"Test case 3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			result := containsDuplicate(testcase.inputs)
			if !reflect.DeepEqual(result, testcase.expects) {
				t.Fatalf("expected: %v, but got: %v, with input: %v ", testcase.expects, result, testcase.inputs)
			}
		})
	}
}
