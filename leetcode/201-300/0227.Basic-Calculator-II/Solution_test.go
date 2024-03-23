package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"TestCase1", "3+2*2", 7},
		{"TestCase2", "3/2 ", 1},
		{"TestCase3", "3+5 / 2 ", 5},
		{"TestCase4", "1*2-3/4+5*6-7*8+9/10", -24},
		{"TestCase5", "282-1*2*13-30-2*2*2/2-95/5*2+55+804+3024", 4067},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
