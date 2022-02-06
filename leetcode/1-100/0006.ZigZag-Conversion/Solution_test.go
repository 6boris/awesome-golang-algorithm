package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 string
		input2 int
		expect string
	}{
		{"1 test 1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"2 test 2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"3 test 3", "A", 1, "A"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := convert(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v-%v",
					c.expect, ret, c.input1, c.input2)
			}
			ret = convert2(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v-%v",
					c.expect, ret, c.input1, c.input2)
			}
			ret = convert3(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v-%v",
					c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
