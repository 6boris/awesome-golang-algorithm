package Solution

import (
	"math/rand"
	"reflect"
	"testing"
)

type Case struct {
	name   string
	inputs *ListNode
	expect bool
}

//func TestSolution1(t *testing.T) {
//	//	测试用例
//	cases := []struct {
//		name   string
//		inputs *ListNode
//		expect bool
//	}{
//		{"TestCacse 1", getRandList(true, 1000), true},
//		{"TestCacse 2", getRandList(false, 1000), false},
//		{"TestCacse 3", getRandList(true, 100000), true},
//		//{"TestCacse 4", getRandList(false, 100000), false},
//	}
//
//	//	开始测试
//	for _, c := range cases {
//		t.Run(c.name, func(t *testing.T) {
//			ret := hasCycle1(c.inputs)
//			if !reflect.DeepEqual(ret, c.expect) {
//				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
//					c.expect, ret, c.inputs)
//			}
//		})
//	}
//
//}

//func TestSolution2(t *testing.T) {
//	//	测试用例
//	cases := []struct {
//		name   string
//		inputs *ListNode
//		expect bool
//	}{
//		{"TestCacse 1", getRandList(false, 10), false},
//		//{"TestCacse 2", getRandList(false, 1000), false},
//		//{"TestCacse 3", getRandList(true, 100000), true},
//		//{"TestCacse 4", getRandList(false, 100000), false},
//	}
//
//	//	开始测试
//	for _, c := range cases {
//		t.Run(c.name, func(t *testing.T) {
//			ret := hasCycle2(c.inputs)
//			if !reflect.DeepEqual(ret, c.expect) {
//				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
//					c.expect, ret, c.inputs)
//			}
//		})
//	}
//
//}

func BenchmarkSoution1(b *testing.B) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *ListNode
		expect bool
	}{
		{"TestCacse 1", getRandList(true, 10), true},
		//{"TestCacse 2", getRandList(false, 1000), false},

		//{"TestCacse 3", getRandList(true, 100000), true},
		//{"TestCacse 4", getRandList(false, 100000), false},
	}

	//	开始测试
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			ret := hasCycle1(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				b.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

// 给list添加一个环
func getRandList(hasCycle bool, len int) *ListNode {
	//	生成没有环的链表
	if !hasCycle {
		return UnmarshalListByRand(100, len)
	}

	// 生成有环的链表，让第index个及节点指向头结点
	index := rand.Intn(len)
	head := UnmarshalListByRand(100, len)
	tmp := head

	for i := 0; i < index; i++ {
		tmp = tmp.Next
	}

	tmp.Next = head
	return head
}

//	更具数组反序列化链表
func UnmarshalListBySlice(nums []int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}

// 随机初始化链表
func UnmarshalListByRand(max_num int, len int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head

	for i := 0; i < len; i++ {
		tmp.Next = &ListNode{Val: rand.Intn(max_num), Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}
