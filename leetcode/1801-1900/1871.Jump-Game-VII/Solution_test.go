package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name             string
		inputs           string
		minJump, maxJump int
		expect           bool
	}{
		{"TestCase1", "011010", 2, 3, true},
		{"TestCase2", "01101110", 2, 3, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.minJump, c.maxJump)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v1",
					c.expect, got, c.inputs, c.minJump, c.maxJump)
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
