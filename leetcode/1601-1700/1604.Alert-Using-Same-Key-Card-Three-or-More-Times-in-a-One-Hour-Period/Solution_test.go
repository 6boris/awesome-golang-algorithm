package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name             string
		keyName, keyTime []string
		expect           []string
	}{
		{"TestCase1", []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}, []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}, []string{"daniel"}},
		{"TestCase2", []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"}, []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}, []string{"bob"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.keyName, c.keyTime)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.keyName, c.keyTime)
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
