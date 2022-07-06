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
		inputs *MultiChildNode
		expect []int
	}{
		{"TestCase1", &MultiChildNode{
			Val: 1,
			Children: []*MultiChildNode{
				{
					Val: 3,
					Children: []*MultiChildNode{
						{
							Val: 5,
						},
						{
							Val: 6,
						},
					},
				},
				{
					Val: 2,
				},
				{
					Val: 4,
				},
			},
		}, []int{1, 3, 5, 6, 2, 4}},
		{"TestCase2", &MultiChildNode{
			Val: 1,
			Children: []*MultiChildNode{
				{
					Val: 2,
				},
				{
					Val: 3,
					Children: []*MultiChildNode{
						{
							Val: 6,
						},
						{
							Val: 7,
							Children: []*MultiChildNode{
								{
									Val: 11,
									Children: []*MultiChildNode{
										{
											Val: 14,
										},
									},
								},
							},
						},
					},
				},
				{
					Val: 4,
					Children: []*MultiChildNode{
						{
							Val: 8,
							Children: []*MultiChildNode{
								{
									Val: 12,
								},
							},
						},
					},
				},
				{
					Val: 5,
					Children: []*MultiChildNode{
						{
							Val: 9,
							Children: []*MultiChildNode{
								{
									Val: 13,
								},
							},
						},
						{
							Val: 10,
						},
					},
				},
			},
		}, []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}},
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

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
