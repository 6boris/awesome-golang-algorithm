package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"Test case 1", "HELLO", "hello"},
		{"Test case 2", "here", "here"},
		{"Test case 1", "LOVELY", "lovely"},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			got := toLowerCase(testcase.input)
			if !reflect.DeepEqual(got, testcase.expect) {
				t.Fatalf("expected: %v, but got: %v, with input: %v ", testcase.expect, got, testcase.input)
			}
		})
	}
}
