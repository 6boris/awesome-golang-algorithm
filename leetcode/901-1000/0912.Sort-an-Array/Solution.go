package Solution

import "math/rand"

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	//head和i用来标记某个点，head标记等待替换的值的下标，i是用作遍历一次数组的下标，类似for i:=0;i<tail;//i++{}
	head, tail, i := 0, len(nums)-1, 0
	//随机获取一个下标，比如len(nums)=5,如果获取的下标为4，nums[4]大于nums[0]-nums[3]的值，则无排序效果
	pivot := rand.Intn(len(nums) - 1)
	nums[pivot], nums[tail] = nums[tail], nums[pivot]
	for i < tail {
		if nums[i] <= nums[tail] {
			nums[head], nums[i] = nums[i], nums[head]
			i++
			head++
		} else {
			i++
		}
	}
	if nums[head] > nums[tail] {
		nums[head], nums[tail] = nums[tail], nums[head]
	}
	//由于nums[head]可以确定大于左边，小于右边，因此无需排序
	quickSort(nums[:head])
	quickSort(nums[head+1:])
	return nums
}

func heapSort(nums []int) []int {
	var sink = func([]int, int, int) {}
	sink = func(heap []int, root, end int) {
		for {
			child := root*2 + 1
			if child > end {
				return
			}
			if child < end && heap[child] <= heap[child+1] {
				child++
			}
			if heap[root] > heap[child] {
				return
			}
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		}
	}

	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		sink(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		sink(nums, 0, end)
	}
	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var merge = func(left, right []int) []int {
		result := make([]int, len(left)+len(right))
		l, r, i := 0, 0, 0
		for l < len(left) && r < len(right) {
			if left[l] < right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
			i++
		}
		copy(result[i:], left[l:])
		copy(result[i+len(left)-l:], right[r:])
		return result
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func selectSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] { //找最小的数
				minIdx = j //保存最小数索引
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}

func insertSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp { //左边比右边大
			nums[j+1] = nums[j] //右移1位
			j--                 //扫描前一个数
		}
		nums[j+1] = tmp //添加到小于它的数的右边
	}
	return nums
}

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		exchange := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] { //两两比较，前面比后面大
				nums[j], nums[j+1] = nums[j+1], nums[j] //交换
				exchange = true
			}
		}
		if !exchange {
			return nums
		}
	}
	return nums
}
func shellSort(nums []int) []int {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			tmp := nums[i]
			j := i - gap
			for j >= 0 && tmp < nums[j] {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = tmp
		}
	}
	return nums
}

func countingSort(nums []int) []int {
	cnt := [100001]int{}
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]+50000]++
	}
	idx := 0
	for i := 0; i < 100001; i++ {
		for cnt[i] > 0 {
			nums[idx] = i - 50000
			idx++
			cnt[i]--
		}
	}
	return nums
}
