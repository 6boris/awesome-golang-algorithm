package Solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect [][]int
	}{
		{
			"1 test 1",
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				{-1, 0, 0, 1},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
			},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := fourSum(c.input1, c.input2)
			fmt.Println(got)
			//if !reflect.DeepEqual(ret, c.expect) {
			//	t.Fatalf("expected: %v, but got: %v, with inputs: %v",
			//		c.expect, ret, c.input1)
			//}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect [][]int
	}{
		{
			"1 test 1",
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				{-1, 0, 0, 1},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
			},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := fourSum1(c.input1, c.input2)
			fmt.Println(got)
			//if !reflect.DeepEqual(ret, c.expect) {
			//	t.Fatalf("expected: %v, but got: %v, with inputs: %v",
			//		c.expect, ret, c.input1)
			//}
		})
	}
}
