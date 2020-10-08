package Solution

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *ListNode
		val    int
		expect *ListNode
	}{
		{"TestCase", &ListNode{
			Val: 1, Next: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 6, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4, Next: &ListNode{
								Val: 5, Next: &ListNode{
									Val: 6, Next: nil,
								},
							},
						},
					},
				}}}, 6, &ListNode{}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := removeElements(c.inputs, c.val)
			for got != nil {
				fmt.Print(got.Val, " ")
				got = got.Next
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
