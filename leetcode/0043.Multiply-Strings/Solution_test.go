package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect string
	}{
		{"0 test 0", []string{"0", "0"}, "0"},
		{"1 test 1", []string{"2", "3"}, "6"},
		{"2 test 2", []string{"123", "456"}, "56088"},
		{"3 test 3", []string{"498828660196", "840477629533"}, "419254329864656431168468"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := multiply(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	b.Run("test1", func(b *testing.B) {
		b.ResetTimer()
		got := multiply("0", "0")
		want := "0"
		if got != want {
			b.Error("test fail")
		}
	})
	b.Run("test2", func(b *testing.B) {
		b.ResetTimer()
		got := multiply("2", "3")
		want := "6"
		if got != want {
			b.Error("test fail")
		}
	})
	b.Run("test3", func(b *testing.B) {
		b.ResetTimer()
		got := multiply("123", "456")
		want := "56088"
		if got != want {
			b.Error("test fail")
		}
	})
	b.Run("test4", func(b *testing.B) {
		b.ResetTimer()
		got := multiply("498828660196", "840477629533")
		want := "419254329864656431168468"
		if got != want {
			b.Error("test fail")
		}
	})
}
