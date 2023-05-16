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
		t1, t2 *QuadNode
		expect *QuadNode
	}{
		{"TestCase1", &QuadNode{
			IsLeaf:      false,
			Val:         true,
			TopLeft:     &QuadNode{IsLeaf: true, Val: true},
			TopRight:    &QuadNode{IsLeaf: true, Val: true},
			BottomLeft:  &QuadNode{IsLeaf: true, Val: false},
			BottomRight: &QuadNode{IsLeaf: true, Val: false},
		}, &QuadNode{
			IsLeaf:  false,
			Val:     true,
			TopLeft: &QuadNode{IsLeaf: true, Val: true},
			TopRight: &QuadNode{
				IsLeaf:      false,
				Val:         true,
				TopLeft:     &QuadNode{IsLeaf: true, Val: false},
				TopRight:    &QuadNode{IsLeaf: true, Val: false},
				BottomLeft:  &QuadNode{IsLeaf: true, Val: true},
				BottomRight: &QuadNode{IsLeaf: true, Val: true},
			},
			BottomLeft:  &QuadNode{IsLeaf: true, Val: true},
			BottomRight: &QuadNode{IsLeaf: true, Val: false},
		}, &QuadNode{
			IsLeaf:      false,
			Val:         false,
			TopLeft:     &QuadNode{IsLeaf: true, Val: true},
			TopRight:    &QuadNode{IsLeaf: true, Val: true},
			BottomLeft:  &QuadNode{IsLeaf: true, Val: true},
			BottomRight: &QuadNode{IsLeaf: true, Val: false},
		}},
		{"TestCase2", &QuadNode{IsLeaf: true, Val: false}, &QuadNode{IsLeaf: true, Val: false}, &QuadNode{IsLeaf: true, Val: false}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.t1, c.t2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.t1, c.t2)
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
