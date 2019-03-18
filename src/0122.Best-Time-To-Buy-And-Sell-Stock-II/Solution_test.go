package Solution

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCase 1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"TestCase 2", []int{1, 2, 3, 4, 5}, 4},
		{"TestCase 3", []int{7, 6, 4, 3, 1}, 0},
	}

	for _, testcase := range cases {
		t.Run(testcase.name, func(t *testing.T) {
			got := maxProfit(testcase.inputs)
			if !reflect.DeepEqual(got, testcase.expect) {
				t.Fatalf("expected: %v, but got %v, with inputs : %v", testcase.expect, got, testcase.inputs)
			}
		})
	}

}
