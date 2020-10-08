package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect bool
	}{
		{"TestCacse 1", []int{2, 3, 1, 1, 4}, true},
		{"TestCacse 2", []int{3, 2, 1, 0, 4}, false},
		{"TestCacse 2", []int{0, 2, 3}, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := canJump(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect bool
	}{
		{"TestCacse 1", []int{2, 3, 1, 1, 4}, true},
		{"TestCacse 2", []int{3, 2, 1, 0, 4}, false},
		{"TestCacse 2", []int{0, 2, 3}, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := canJump2(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
func TestSolution3(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect bool
	}{
		{"TestCase 1", []int{2, 3, 1, 1, 4}, true},
		{"TestCase 2", []int{3, 2, 1, 0, 4}, false},
		{"TestCase 3", []int{0, 2, 3}, false},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := canJump3(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
