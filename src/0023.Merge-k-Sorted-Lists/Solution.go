package Solution

import "math/rand"

//	1.暴力转换
func mergeKLists1(lists []*ListNode) *ListNode {
	nums := make([]int, 0, 10)

	//	1.链表转化为数组
	for _, v := range lists {
		nums = append(nums, MarshalListNodeToSlice(v)...)
	}
	//	对数组排序
	nums = QuickSort(nums)

	//	将数组转化为链表
	return MarshalSliceToListNode(nums)
}

// 插入排序(Insert Sort)
func InertSort(arr []int) []int {
	var i, j, tmp int
	for i = 1; i < len(arr); i++ {
		tmp = arr[i]
		for j = i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}

// 快速排序(Qucik Sort)
func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = QuickSort(low_part)
	high_part = QuickSort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}

//	将ListNode 序列化为一个数组
func MarshalListNodeToSlice(node *ListNode) []int {
	ans := make([]int, 0, 10)
	for node != nil {
		ans = append(ans, node.Val)
		node = node.Next
	}
	return ans
}

//	将数组转换为ListNode
func MarshalSliceToListNode(nums []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}
