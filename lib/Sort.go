package Solution

import "math/rand"

/***********************************************************************************************************
	排序算法
		1.冒泡排序(Bubble Sort)			BubbleSort()			https://en.wikipedia.org/wiki/Bubble_sort
		2.插入排序(Insert Sort) 			InsertSort()			https://en.wikipedia.org/wiki/Insertion_sort
		3.选择排序(Select Sort) 			SelectSort()			https://en.wikipedia.org/wiki/Selection_sort
		4.快速排序(Quick Sort) 			QuickSort()				https://en.wikipedia.org/wiki/Quicksort
		5.归并排序(Merge Sort) 			MergeSort()				https://en.wikipedia.org/wiki/Merge_sort
		6.堆排序(Heap Sort)   			HeapSort()				https://en.wikipedia.org/wiki/Heapsort
		7.希尔排序(Shell Sort) 			ShellSort()				https://en.wikipedia.org/wiki/Shellsort
		8.鸡尾酒排序(Cocktail Shaker Sort)CocktailShakerSort()	https://en.wikipedia.org/wiki/Cocktail_sort
		9.梳排序(Comb Sort) 				CombSort()				https://en.wikipedia.org/wiki/Combsort
		10.计数排序(Counting Sort) 		CountingSort()			https://en.wikipedia.org/wiki/Counting_sort
		11.地精排序(Gnome Sort) 			GnomeSort()				https://en.wikipedia.org/wiki/Gnome_sort
		12.奇偶排序(Odd Even Sort) 		OddEvenSort()			https://en.wikipedia.org/wiki/Odd-even_sort
***********************************************************************************************************/

// 冒泡排序(Bubble Sort)
func BubbleSort(arr []int) []int {
	tmp := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

// 插入排序(Insert Sort)
func InsertSort(arr []int) []int {
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

// 选择排序(Select Sort)
func SelectSort(arr []int) []int {
	var min = 0
	var tmp = 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
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

// 归并排序(Merge Sort)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return Merge(left, right)
}
func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

// 堆排序(Heap Sort)
func HeapSort(arr []int) []int {
	i := 0
	tmp := 0
	for i = len(arr)/2 - 1; i >= 0; i-- {
		arr = Sift(arr, i, len(arr))
	}

	for i = len(arr) - 1; i >= 1; i-- {
		tmp = arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		arr = Sift(arr, 0, i)
	}
	return arr
}

func Sift(arr []int, i int, arrLen int) []int {
	done := false

	tmp := 0
	maxChild := 0

	for (i*2+1 < arrLen) && (!done) {
		if i*2+1 == arrLen-1 {
			maxChild = i*2 + 1
		} else if arr[i*2+1] > arr[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if arr[i] < arr[maxChild] {
			tmp = arr[i]
			arr[i] = arr[maxChild]
			arr[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}

// 希尔排序(Shell Sort)
func ShellSort(arr []int) []int {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return arr
}

// 鸡尾酒排序(Cocktail shaker sort)
func CocktailShakerSort(arr []int) []int {
	tmp := 0
	for i := 0; i < len(arr)/2; i++ {

		left := 0
		right := len(arr) - 1

		for left <= right {
			if arr[left] > arr[left+1] {
				tmp = arr[left]
				arr[left] = arr[left+1]
				arr[left+1] = tmp
			}

			left++

			if arr[right-1] > arr[right] {
				tmp = arr[right-1]
				arr[right-1] = arr[right]
				arr[right] = tmp
			}

			right--

		}
	}
	return arr
}

// 梳排序(Comb Sort)
func CombSort(arr []int) []int {
	tmp := 0
	arrLen := len(arr)
	gap := arrLen

	for gap > 1 {
		gap = gap * 10 / 13 //shrink factor is 1.3

		for i := 0; i+gap < arrLen; i++ {
			if arr[i] > arr[i+gap] {
				tmp = arr[i]
				arr[i] = arr[i+gap]
				arr[i+gap] = tmp
			}
		}
	}
	return arr
}

// 计数排序(Counting Sort)
func CountingSort(arr []int) []int {
	k := getK(arr)
	array_of_counts := make([]int, k)

	for i := 0; i < len(arr); i++ {
		array_of_counts[arr[i]] += 1
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if array_of_counts[i] > 0 {
				arr[j] = i
				j += 1
				array_of_counts[i] -= 1
				continue
			}
			break
		}
	}
	return arr
}

func getK(arr []int) int {
	if len(arr) == 0 {
		return 1
	}
	k := arr[0]
	for _, v := range arr {
		if v > k {
			k = v
		}
	}
	return k + 1
}

// 地精排序(Gnome Sort)
func GnomeSort(arr []int) []int {
	i := 1
	tmp := 0

	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp

			if i > 1 {
				i--
			}
		}
	}
	return arr
}

// 奇偶排序(Odd Even Sort)
func OddEvenSort(arr []int) []int {
	tmp, isSorted := 0, false

	for isSorted == false {

		isSorted = true

		for i := 1; i < len(arr)-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp

				isSorted = false
			}
		}

		for i := 0; i < len(arr)-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp

				isSorted = false
			}
		}
	}
	return arr
}
