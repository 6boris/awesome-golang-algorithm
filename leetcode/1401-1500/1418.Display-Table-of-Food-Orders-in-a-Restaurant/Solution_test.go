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
		inputs [][]string
		expect [][]string
	}{
		{"TestCase1", [][]string{
			{"David", "3", "Ceviche"},
			{"Corina", "10", "Beef Burrito"},
			{"David", "3", "Fried Chicken"},
			{"Carla", "5", "Water"},
			{"Carla", "5", "Ceviche"},
			{"Rous", "3", "Ceviche"},
		}, [][]string{
			{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"},
			{"3", "0", "2", "1", "0"},
			{"5", "0", "1", "0", "1"},
			{"10", "1", "0", "0", "0"},
		}},
		{"TestCase2", [][]string{
			{"James", "12", "Fried Chicken"},
			{"Ratesh", "12", "Fried Chicken"},
			{"Amadeus", "12", "Fried Chicken"},
			{"Adam", "1", "Canadian Waffles"},
			{"Brianna", "1", "Canadian Waffles"},
		}, [][]string{
			{"Table", "Canadian Waffles", "Fried Chicken"},
			{"1", "2", "0"},
			{"12", "0", "3"},
		}},
		{"TestCase3", [][]string{
			{"Laura", "2", "Bean Burrito"},
			{"Jhon", "2", "Beef Burrito"},
			{"Melissa", "2", "Soda"},
		}, [][]string{
			{"Table", "Bean Burrito", "Beef Burrito", "Soda"},
			{"2", "1", "1", "1"},
		}},
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

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
