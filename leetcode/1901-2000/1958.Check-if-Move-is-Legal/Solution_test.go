package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		board        [][]byte
		rMove, cMove int
		color        byte
		expect       bool
	}{
		{"TestCase1", [][]byte{
			[]byte{'.', '.', '.', 'B', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', 'W', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', 'W', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', 'W', '.', '.', '.', '.'},
			[]byte{'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'},
			[]byte{'.', '.', '.', 'B', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', 'B', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', 'W', '.', '.', '.', '.'},
		}, 4, 3, 'B', true},
		{"TestCase2", [][]byte{
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', 'B', '.', '.', 'W', '.', '.', '.'},
			[]byte{'.', '.', 'W', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', 'W', 'B', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', 'B', 'W', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', 'W', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', 'B'},
		}, 4, 4, 'W', false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.board, c.rMove, c.cMove, c.color)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.board, c.rMove, c.cMove, c.color)
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
