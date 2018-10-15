package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"1 test 1", "42", 42},
		{"2 test 2", "   -42", -42},
		{"3 test 3", "4193 with words", 4193},
		{"4 test 4", "words and 987", 0},
		{"5 test 5", "-91283472332", -2147483648},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := myAtoi(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
