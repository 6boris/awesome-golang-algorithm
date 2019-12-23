package Solution

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *TreeNode
		sum    int
		expect bool
	}{
		{"TestCase",
			&TreeNode{Val: 5,
				Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
				Right: &TreeNode{Val: 8, Left: nil, Right: nil},
			},
			9,
			true,
		},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := hasPathSum(c.inputs, c.sum)
			fmt.Println(got)
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
