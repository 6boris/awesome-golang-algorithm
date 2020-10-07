package Solution

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	ans := []int{}
	for i := 0; i < len(nums)-k+1; i++ {
		tmp := nums[i]
		for j := i + 1; j < i+k; j++ {
			if tmp < nums[j] {
				tmp = nums[j]
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	} else if k == 1 {
		return nums
	}
	queue, ans := list.New(), make([]int, len(nums)-k+1)

	for left, right := 1-k, 0; right < len(nums); left, right = left+1, right+1 {
		if left > 0 && queue.Front().Value.(int) == nums[left-1] {
			queue.Remove(queue.Front())
		}
		for queue.Len() != 0 && queue.Back().Value.(int) < nums[right] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(nums[right])
		if left >= 0 {
			ans[left] = queue.Front().Value.(int)
		}
	}
	return ans
}

func maxSlidingWindow3(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	queue := make([]int, 0, k)
	ans := make([]int, len(nums)-k+1)

	for left, right := 1-k, 0; right < len(nums); left, right = left+1, right+1 {
		if left > 0 && queue[0] == nums[left-1] {
			queue = queue[1:len(queue)]
		}
		for len(queue) != 0 && queue[len(queue)-1] < nums[right] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[right])
		if left >= 0 {
			ans[left] = queue[0]
		}
	}
	return ans
}
