package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"1 test 1", 4, 2},
		{"2 test 2", 8, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := mySqrt(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"1 test 1", 4, 2},
		{"2 test 2", 8, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := mySqrt1(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
