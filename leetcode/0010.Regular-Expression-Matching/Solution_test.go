package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect bool
	}{
		{"TestCacse 1", []string{"aa", "a"}, false},
		{"TestCacse 2", []string{"aa", "a*"}, true},
		{"TestCacse 3", []string{"ab", ".*"}, true},
		{"TestCacse 4", []string{"aab", "c*a*b"}, true},
		{"TestCacse 5", []string{"mississippi", "mis*is*p*."}, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isMatch(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
