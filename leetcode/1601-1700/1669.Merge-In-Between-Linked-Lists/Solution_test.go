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
		l1, l2 *ListNode
		a, b   int
		expect *ListNode
	}{
		{"TestCase1", &ListNode{
			Val: 10,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 13,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{Val: 5},
						},
					},
				},
			},
		}, &ListNode{Val: 1000000, Next: &ListNode{Val: 1000001, Next: &ListNode{Val: 1000002}}}, 3, 4, &ListNode{
			Val: 10,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 13,
					Next: &ListNode{
						Val: 1000000,
						Next: &ListNode{
							Val:  1000001,
							Next: &ListNode{Val: 1000002, Next: &ListNode{Val: 5}},
						},
					},
				},
			},
		}},
		{"TestCase2", &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}},
						},
					},
				},
			},
		}, &ListNode{Val: 1000000, Next: &ListNode{Val: 1000001, Next: &ListNode{Val: 1000002, Next: &ListNode{Val: 1000003, Next: &ListNode{Val: 1000004}}}}}, 2, 5, &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1000000,
					Next: &ListNode{
						Val: 1000001,
						Next: &ListNode{
							Val:  1000002,
							Next: &ListNode{Val: 1000003, Next: &ListNode{Val: 1000004, Next: &ListNode{Val: 6}}},
						},
					},
				},
			},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.l1, c.a, c.b, c.l2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.l1, c.l2, c.a, c.b)
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
