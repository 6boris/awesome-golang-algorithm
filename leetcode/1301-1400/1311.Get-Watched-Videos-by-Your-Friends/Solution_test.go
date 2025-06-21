package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		watchedVideos [][]string
		friends       [][]int
		id, level     int
		expect        []string
	}{
		{"TestCase1", [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}, [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}, 0, 1, []string{"B", "C"}},
		{"TestCase2", [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}, [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}, 0, 2, []string{"D"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.watchedVideos, c.friends, c.id, c.level)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.watchedVideos, c.friends, c.id, c.level)
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
