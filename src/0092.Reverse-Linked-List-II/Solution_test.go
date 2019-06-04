package Solution

import (
	"testing"
)

type InputCase struct {
	element []int
	m int
	n int
}

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
<<<<<<< HEAD
		inputs InputCase
		expect []int
	}{
		{
			"TestCase 1",
			InputCase{
				[]int{1,2,3,4,5},
				2,
				4,
			},
			[]int{1,4,3,2,5},
		},
=======
		inputs ListNode
		m      int
		n      int
		expect ListNode
	}{
		{"TestCase 1", ListNode{
			Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}},
				}},
		}, 2, 4, ListNode{}},
>>>>>>> develop
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
<<<<<<< HEAD
			ret := Solution(UnmarshalListBySlice(c.inputs.element), c.inputs.m, c.inputs.n)
			if !isEqual(ret, UnmarshalListBySlice(c.expect)) {
				PrintList(ret)
				PrintList(UnmarshalListBySlice(c.expect))
				t.Fatalf("expected: %v, but got: %v, with inputs: %v", c.expect, ret, c.inputs)
=======
			ret := reverseBetween(&c.inputs, c.m, c.n)
			if !reflect.DeepEqual(ret, c.expect) {
				PrintList(ret)
				//t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				//	c.expect, ret, c.inputs)
>>>>>>> develop
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
