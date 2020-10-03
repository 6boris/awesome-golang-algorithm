package Solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 float64
		input2 int
		expect float64
	}{
		{"Test Case 1", 2.00000, 10, 1024.00000},
		{"Test Case 2", 2.10000, 3, 9.26100},
		{"Test Case 3", 2.00000, -2, 0.25000},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := fmt.Sprintf("%.5f", myPow(c.input1, c.input2))
			if !reflect.DeepEqual(got, fmt.Sprintf("%.5f", c.expect)) {
				t.Fatalf("expected: %v, but got: %v, with input1: %v input2: %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 float64
		input2 int
		expect float64
	}{
		{"Test Case 1", 2.00000, 10, 1024.00000},
		{"Test Case 2", 2.10000, 3, 9.26100},
		{"Test Case 3", 2.00000, -2, 0.25000},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := fmt.Sprintf("%.5f", myPow2(c.input1, c.input2))
			if !reflect.DeepEqual(got, fmt.Sprintf("%.5f", c.expect)) {
				t.Fatalf("expected: %v, but got: %v, with input1: %v input2: %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}
