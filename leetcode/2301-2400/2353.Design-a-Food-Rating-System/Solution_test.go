package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name            string
		foods, cuisines []string
		ratings         []int
		opts            []opt
		expect          []string
	}{
		{"TestCase1", []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
			[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7},
			[]opt{
				{name: "h", commonStr: "korean"},
				{name: "h", commonStr: "japanese"},
				{name: "c", commonStr: "sushi", rating: 16},
				{name: "h", commonStr: "japanese"},
				{name: "c", commonStr: "ramen", rating: 16},
				{name: "h", commonStr: "japanese"}},
			[]string{"kimchi", "ramen", "sushi", "ramen"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.foods, c.cuisines, c.ratings, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.foods, c.cuisines, c.ratings, c.opts)
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
