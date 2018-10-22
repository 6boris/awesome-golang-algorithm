package Solution

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect []int
	}{
		{"1 test 1", [][]int{{2, 7, 11, 15}, {9}}, []int{0, 1}},
		{"2 test 2", [][]int{{3, 2, 4}, {6}}, []int{1, 2}},
		{"3 test 3", [][]int{{2, 7, 11, 15}, {9}}, []int{0, 1}},
		{"4 test 4", [][]int{{7, 6, 5, 3, 2, 1, 4, 9, 10}, {17}}, []int{0, 8}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := TwoSum1(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func TestTwoSum2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect []int
	}{
		{"1 test 1", [][]int{{2, 7, 11, 15}, {9}}, []int{0, 1}},
		{"2 test 2", [][]int{{3, 2, 4}, {6}}, []int{1, 2}},
		{"3 test 3", [][]int{{2, 7, 11, 15}, {9}}, []int{0, 1}},
		{"4 test 4", [][]int{{7, 6, 5, 3, 2, 1, 4, 9, 10}, {17}}, []int{0, 8}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := TwoSum2(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

const N = 100

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum1(nums, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2(nums, 9)
	}
}
