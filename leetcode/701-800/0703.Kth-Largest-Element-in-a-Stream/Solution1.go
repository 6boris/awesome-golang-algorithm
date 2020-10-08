package Solution

import (
	"fmt"
	"math"
)

type KthLargest1 struct {
	heap []int
	k    int
}

func Constructor1(k int, nums []int) KthLargest1 {
	//	初始化结构体
	kthLargest := KthLargest1{heap: make([]int, k)}
	kthLargest.k = k
	//	需要将原head 总默认值设置最小
	for i, _ := range kthLargest.heap {
		kthLargest.heap[i] = math.MinInt32
	}

	//	初始化数组
	for i := 0; i < k && i < len(nums); i++ {
		kthLargest.heap[i] = nums[i]
	}

	//	对数组进行排序
	kthLargest.heap = InsertSort(kthLargest.heap)

	//	当第一次初始化的数组长度比K大的时候处理后面的数
	for i := k; i < len(nums); i++ {
		kthLargest.Add(nums[i])
	}

	fmt.Println(kthLargest.heap)

	return kthLargest
}

//	想数据流中添加一个数并返回第K大的数
func (this *KthLargest1) Add(val int) int {
	//	如果添加进来的数比末尾的打就直接替换掉末尾的数
	if val > this.heap[this.k-1] {
		this.heap[this.k-1] = val
	}
	//	替换完后需要重新排序
	this.heap = InsertSort(this.heap)
	return this.heap[this.k-1]
}

// 插入排序(Insert Sort)
func InsertSort(arr []int) []int {
	var i, j, tmp int
	for i = 1; i < len(arr); i++ {
		tmp = arr[i]
		for j = i; j > 0 && arr[j-1] < tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}
