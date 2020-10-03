package Solution

import (
	"reflect"
	"testing"
)

func TestCheckPerfectNumber(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Test case 1", 28, true},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			got := checkPerfectNumber(testcase.input)
			if !reflect.DeepEqual(got, testcase.expected) {
				t.Fatalf("expected: %v, but got: %v, with input: %v",
					testcase.expected, got, testcase.input)
			}
		})
	}
}
