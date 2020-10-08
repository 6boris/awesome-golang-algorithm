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
		expect int
	}{
		{"TestCacse 1", []int{2, 3, 1, 1, 4}, 2},
		{"TestCacse 1", []int{2}, 0},
		{"TestCacse 1", []int{10, 3, 4}, 1},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := jump(c.inputs)
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
		expect int
	}{
		{"TestCase 1", []int{2, 3, 1, 1, 4}, 2},
		{"TestCase 1", []int{2}, 0},
		{"TestCase 1", []int{10, 3, 4}, 1},
		{"TestCase 4", []int{4, 1, 1, 3, 1, 1, 1}, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := jump2(c.inputs)
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
		expect int
	}{
		{"TestCase 1", []int{2, 3, 1, 1, 4}, 2},
		{"TestCase 2", []int{2}, 0},
		{"TestCase 3", []int{10, 3, 4}, 1},
		{"TestCase 4", []int{4, 1, 1, 3, 1, 1, 1}, 2},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := jump3(c.inputs)
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
